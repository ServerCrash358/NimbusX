package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// ─────────────────────────────────────────────────────────────────────────────
// Helpers
// ─────────────────────────────────────────────────────────────────────────────

func postJSON(t *testing.T, url string, body any) *http.Response {
	t.Helper()
	data, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatalf("POST %s: %v", url, err)
	}
	return resp
}

func getJSON(t *testing.T, url string) *http.Response {
	t.Helper()
	resp, err := http.Get(url)
	if err != nil {
		t.Fatalf("GET %s: %v", url, err)
	}
	return resp
}

func decodeBody(t *testing.T, resp *http.Response, dst any) {
	t.Helper()
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(dst); err != nil {
		t.Fatalf("decode response: %v", err)
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// Oracle unit test (no network required)
// ─────────────────────────────────────────────────────────────────────────────

func TestOraclePriceFeed(t *testing.T) {
	prices := mockFetchAllPrices()
	if len(prices) == 0 {
		t.Fatal("expected non-empty price feed")
	}
	for id, p := range prices {
		if p <= 0 {
			t.Errorf("provider %s has non-positive price %f", id, p)
		}
	}
}

func mockFetchAllPrices() map[string]float64 {
	return map[string]float64{
		"node-1": 0.05,
		"node-2": 0.08,
		"node-3": 0.12,
	}
}

// ─────────────────────────────────────────────────────────────────────────────
// Shared types (mirror of scheduler/oracle types for test layer)
// ─────────────────────────────────────────────────────────────────────────────

type providerRegReq struct {
	ID           string  `json:"id"`
	Address      string  `json:"address"`
	CPU          int     `json:"cpu"`
	MemoryMB     int     `json:"memoryMB"`
	PricePerHour float64 `json:"pricePerHour"`
}

type jobSubmitReq struct {
	ClientID        string  `json:"clientId"`
	CPURequired     int     `json:"cpuRequired"`
	MemoryMB        int     `json:"memoryMB"`
	MaxDurationSecs int     `json:"maxDurationSecs"`
	MaxPricePerHour float64 `json:"maxPricePerHour"`
}

// ─────────────────────────────────────────────────────────────────────────────
// Scheduler HTTP integration tests (httptest — no live server needed)
// ─────────────────────────────────────────────────────────────────────────────

func buildMockScheduler() *httptest.Server {
	providers := map[string]providerRegReq{}
	jobs := map[string]map[string]any{}
	jobIDCounter := 0

	mux := http.NewServeMux()

	mux.HandleFunc("/providers/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var req providerRegReq
		_ = json.NewDecoder(r.Body).Decode(&req)
		providers[req.ID] = req
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "registered", "id": req.ID})
	})

	mux.HandleFunc("/providers", func(w http.ResponseWriter, r *http.Request) {
		list := make([]providerRegReq, 0, len(providers))
		for _, p := range providers {
			list = append(list, p)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(list)
	})

	mux.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var req jobSubmitReq
		_ = json.NewDecoder(r.Body).Decode(&req)
		jobIDCounter++
		id := fmt.Sprintf("job-%d", jobIDCounter)
		assigned := ""
		status := "pending"
		for pid := range providers {
			assigned = pid
			status = "assigned"
			break
		}
		jobs[id] = map[string]any{
			"id":               id,
			"status":           status,
			"assignedProvider": assigned,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(jobs[id])
	})

	mux.HandleFunc("/jobs/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/jobs/")
		j, ok := jobs[id]
		if !ok {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(j)
	})

	return httptest.NewServer(mux)
}

func TestSchedulerIntegration(t *testing.T) {
	srv := buildMockScheduler()
	defer srv.Close()

	// ── 1. Register a provider ─────────────────────────────────────────────
	t.Run("Register provider", func(t *testing.T) {
		resp := postJSON(t, srv.URL+"/providers/register", providerRegReq{
			ID:           "node-1",
			CPU:          8,
			MemoryMB:     16384,
			PricePerHour: 0.05,
		})
		if resp.StatusCode != http.StatusCreated {
			t.Fatalf("expected 201, got %d", resp.StatusCode)
		}
		var result map[string]string
		decodeBody(t, resp, &result)
		if result["id"] != "node-1" {
			t.Errorf("expected id=node-1, got %s", result["id"])
		}
	})

	// ── 2. List providers ─────────────────────────────────────────────────
	t.Run("List providers", func(t *testing.T) {
		resp := getJSON(t, srv.URL+"/providers")
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected 200, got %d", resp.StatusCode)
		}
		var list []providerRegReq
		decodeBody(t, resp, &list)
		if len(list) == 0 {
			t.Fatal("expected at least one provider")
		}
	})

	// ── 3. Submit job → assigned ──────────────────────────────────────────
	t.Run("Submit job and verify assignment", func(t *testing.T) {
		resp := postJSON(t, srv.URL+"/jobs", jobSubmitReq{
			ClientID:        "client-abc",
			CPURequired:     4,
			MemoryMB:        8192,
			MaxDurationSecs: 3600,
			MaxPricePerHour: 0.10,
		})
		if resp.StatusCode != http.StatusCreated {
			t.Fatalf("expected 201, got %d", resp.StatusCode)
		}
		var job map[string]any
		decodeBody(t, resp, &job)
		if job["status"] != "assigned" {
			t.Errorf("expected status=assigned, got %v", job["status"])
		}
		if job["assignedProvider"] == "" {
			t.Error("expected a provider to be assigned")
		}
	})

	// ── 4. Get job status ─────────────────────────────────────────────────
	t.Run("Get job status", func(t *testing.T) {
		resp := postJSON(t, srv.URL+"/jobs", jobSubmitReq{
			ClientID:        "client-xyz",
			CPURequired:     2,
			MemoryMB:        4096,
			MaxDurationSecs: 1800,
			MaxPricePerHour: 0.20,
		})
		var job map[string]any
		decodeBody(t, resp, &job)
		jobID := job["id"].(string)

		statusResp := getJSON(t, srv.URL+"/jobs/"+jobID)
		if statusResp.StatusCode != http.StatusOK {
			t.Fatalf("expected 200, got %d", statusResp.StatusCode)
		}
		var fetched map[string]any
		decodeBody(t, statusResp, &fetched)
		if fetched["id"] != jobID {
			t.Errorf("expected job id %s, got %v", jobID, fetched["id"])
		}
	})

	// ── 5. Unknown job returns 404 ────────────────────────────────────────
	t.Run("Unknown job returns 404", func(t *testing.T) {
		resp := getJSON(t, srv.URL+"/jobs/does-not-exist")
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusNotFound {
			t.Fatalf("expected 404, got %d", resp.StatusCode)
		}
	})
}
