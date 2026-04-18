package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Scheduler is the core dispatch engine.
type Scheduler struct {
	mu        sync.RWMutex
	jobs      map[string]*Job
	pm        *ProviderManager
	oracleURL string
}

func NewScheduler(pm *ProviderManager, oracleURL string) *Scheduler {
	return &Scheduler{
		jobs:      make(map[string]*Job),
		pm:        pm,
		oracleURL: oracleURL,
	}
}

// SubmitJob accepts a new job, scores available providers, and assigns the best match.
func (s *Scheduler) SubmitJob(req JobSubmitRequest) (*Job, error) {
	job := &Job{
		ID:              uuid.New().String(),
		ClientID:        req.ClientID,
		CPURequired:     req.CPURequired,
		MemoryMB:        req.MemoryMB,
		MaxDurationSecs: req.MaxDurationSecs,
		MaxPricePerHour: req.MaxPricePerHour,
		Status:          JobStatusPending,
		SubmittedAt:     time.Now(),
	}

	s.mu.Lock()
	s.jobs[job.ID] = job
	s.mu.Unlock()

	// Find and assign best provider
	if err := s.assignBestProvider(job); err != nil {
		// Job stays in Pending — will be retried or returned to client
		return job, fmt.Errorf("no suitable provider: %w", err)
	}

	return job, nil
}

// assignBestProvider scores all eligible providers and assigns the highest-scoring one.
func (s *Scheduler) assignBestProvider(job *Job) error {
	candidates := s.pm.GetActive()

	// Filter to eligible providers
	var eligible []*Provider
	for _, p := range candidates {
		if p.CPU >= job.CPURequired &&
			p.MemoryMB >= job.MemoryMB &&
			p.PricePerHour <= job.MaxPricePerHour {
			eligible = append(eligible, p)
		}
	}

	if len(eligible) == 0 {
		return fmt.Errorf("no eligible providers for job %s", job.ID)
	}

	// Score and sort
	scored := ScoreProviders(eligible, job, s.fetchOraclePrice)
	sort.Slice(scored, func(i, j int) bool {
		return scored[i].Score > scored[j].Score
	})

	best := scored[0].Provider

	s.mu.Lock()
	now := time.Now()
	job.AssignedProvider = best.ID
	job.Status = JobStatusAssigned
	job.StartedAt = &now
	s.mu.Unlock()

	return nil
}

// fetchOraclePrice retrieves the current market price from the oracle service.
// Returns 0 on error (scoring will use provider's self-reported price instead).
func (s *Scheduler) fetchOraclePrice(providerID string) float64 {
	resp, err := http.Get(fmt.Sprintf("%s/price/%s", s.oracleURL, providerID))
	if err != nil || resp.StatusCode != http.StatusOK {
		return 0
	}
	defer resp.Body.Close()

	var result struct {
		Price float64 `json:"price"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0
	}
	return result.Price
}

// GetJob retrieves a job by ID.
func (s *Scheduler) GetJob(id string) (*Job, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	j, ok := s.jobs[id]
	if !ok {
		return nil, false
	}
	copy := *j
	return &copy, true
}

// ── HTTP Handlers ──────────────────────────────────────────────────────────

// HandleSubmitJob handles POST /jobs
func (s *Scheduler) HandleSubmitJob(w http.ResponseWriter, r *http.Request) {
	var req JobSubmitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid body: %v", err), http.StatusBadRequest)
		return
	}
	if req.CPURequired <= 0 || req.MemoryMB <= 0 {
		http.Error(w, "cpuRequired and memoryMB must be positive", http.StatusBadRequest)
		return
	}

	job, err := s.SubmitJob(req)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusAccepted) // queued but unassigned
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	_ = json.NewEncoder(w).Encode(job)
}

// HandleGetJob handles GET /jobs/{id}
func (s *Scheduler) HandleGetJob(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	job, ok := s.GetJob(id)
	if !ok {
		http.Error(w, "job not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(job)
}
