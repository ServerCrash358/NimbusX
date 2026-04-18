package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// ProviderManager maintains the in-memory registry of compute providers.
type ProviderManager struct {
	mu        sync.RWMutex
	providers map[string]*Provider
}

func NewProviderManager() *ProviderManager {
	return &ProviderManager{
		providers: make(map[string]*Provider),
	}
}

// Register adds or updates a provider in the registry.
func (pm *ProviderManager) Register(p Provider) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	p.Active = true
	p.RegisteredAt = time.Now()
	pm.providers[p.ID] = &p
}

// Deregister marks a provider as inactive.
func (pm *ProviderManager) Deregister(id string) bool {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if p, ok := pm.providers[id]; ok {
		p.Active = false
		return true
	}
	return false
}

// GetActive returns a snapshot of all active providers.
func (pm *ProviderManager) GetActive() []*Provider {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	var active []*Provider
	for _, p := range pm.providers {
		if p.Active {
			// Return a copy to avoid race conditions
			copy := *p
			active = append(active, &copy)
		}
	}
	return active
}

// Get returns a single provider by ID.
func (pm *ProviderManager) Get(id string) (*Provider, bool) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	p, ok := pm.providers[id]
	if !ok {
		return nil, false
	}
	copy := *p
	return &copy, true
}

// GetCount returns the number of active providers.
func (pm *ProviderManager) GetCount() int {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	count := 0
	for _, p := range pm.providers {
		if p.Active {
			count++
		}
	}
	return count
}

// UpdateReputation adjusts a provider's reputation score (clamped 0-100).
func (pm *ProviderManager) UpdateReputation(id string, delta int) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if p, ok := pm.providers[id]; ok {
		p.Reputation += delta
		if p.Reputation < 0 {
			p.Reputation = 0
		}
		if p.Reputation > 100 {
			p.Reputation = 100
		}
	}
}

// ── HTTP Handlers ──────────────────────────────────────────────────────────

// HandleRegisterProvider handles POST /providers/register
func (pm *ProviderManager) HandleRegisterProvider(w http.ResponseWriter, r *http.Request) {
	var req ProviderRegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("invalid body: %v", err), http.StatusBadRequest)
		return
	}
	if req.ID == "" || req.CPU <= 0 || req.MemoryMB <= 0 {
		http.Error(w, "id, cpu, and memoryMB are required", http.StatusBadRequest)
		return
	}

	pm.Register(Provider{
		ID:           req.ID,
		Address:      req.Address,
		CPU:          req.CPU,
		MemoryMB:     req.MemoryMB,
		PricePerHour: req.PricePerHour,
		Reputation:   100,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "registered", "id": req.ID})
}

// HandleListProviders handles GET /providers
func (pm *ProviderManager) HandleListProviders(w http.ResponseWriter, r *http.Request) {
	active := pm.GetActive()
	if active == nil {
		active = []*Provider{}
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(active)
}
