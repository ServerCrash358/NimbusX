package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Gateway proxies user-facing API calls to the internal scheduler service.
type Gateway struct {
	schedulerURL string
	client       *http.Client
}

func NewGateway(schedulerURL string) *Gateway {
	return &Gateway{
		schedulerURL: schedulerURL,
		client:       &http.Client{},
	}
}

// ── Request / Response types ───────────────────────────────────────────────

// SubmitJobRequest is the public-facing payload from an API client.
type SubmitJobRequest struct {
	ClientID        string  `json:"clientId"`
	CPURequired     int     `json:"cpuRequired"`
	MemoryMB        int     `json:"memoryMB"`
	MaxDurationSecs int     `json:"maxDurationSecs"`
	MaxPricePerHour float64 `json:"maxPricePerHour"`
}

// ── Handlers ──────────────────────────────────────────────────────────────

// HandleSubmitJob handles POST /submit-job
// Validates the request then forwards it to the scheduler's POST /jobs endpoint.
func (g *Gateway) HandleSubmitJob(w http.ResponseWriter, r *http.Request) {
	var req SubmitJobRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if req.CPURequired <= 0 {
		jsonError(w, "cpuRequired must be > 0", http.StatusBadRequest)
		return
	}
	if req.MemoryMB <= 0 {
		jsonError(w, "memoryMB must be > 0", http.StatusBadRequest)
		return
	}
	if req.MaxPricePerHour <= 0 {
		jsonError(w, "maxPricePerHour must be > 0", http.StatusBadRequest)
		return
	}

	body, _ := json.Marshal(req)
	resp, err := g.client.Post(g.schedulerURL+"/jobs", "application/json", bytes.NewReader(body))
	if err != nil {
		jsonError(w, fmt.Sprintf("scheduler unavailable: %v", err), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	proxyResponse(w, resp)
}

// HandleGetProviders handles GET /providers
// Forwards to scheduler's GET /providers endpoint.
func (g *Gateway) HandleGetProviders(w http.ResponseWriter, r *http.Request) {
	resp, err := g.client.Get(g.schedulerURL + "/providers")
	if err != nil {
		jsonError(w, fmt.Sprintf("scheduler unavailable: %v", err), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	proxyResponse(w, resp)
}

// HandleGetJobStatus handles GET /job-status?id=<jobID>
// Forwards to scheduler's GET /jobs/{id} endpoint.
func (g *Gateway) HandleGetJobStatus(w http.ResponseWriter, r *http.Request) {
	jobID := r.URL.Query().Get("id")
	if jobID == "" {
		jsonError(w, "query param 'id' is required", http.StatusBadRequest)
		return
	}

	resp, err := g.client.Get(fmt.Sprintf("%s/jobs/%s", g.schedulerURL, jobID))
	if err != nil {
		jsonError(w, fmt.Sprintf("scheduler unavailable: %v", err), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	proxyResponse(w, resp)
}

// ── Helpers ────────────────────────────────────────────────────────────────

func proxyResponse(w http.ResponseWriter, resp *http.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)
}

func jsonError(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": msg})
}
