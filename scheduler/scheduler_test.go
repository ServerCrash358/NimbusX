package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSubmitJobAssignsBestEligibleProvider(t *testing.T) {
	oracle := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]float64{"price": 0.2})
	}))
	defer oracle.Close()

	pm := NewProviderManager()
	pm.Register(Provider{ID: "p1", CPU: 8, MemoryMB: 8192, PricePerHour: 0.3, Reputation: 80, Address: "0xabc"})
	pm.Register(Provider{ID: "p2", CPU: 16, MemoryMB: 32768, PricePerHour: 0.25, Reputation: 95, Address: "0xdef"})

	s := NewScheduler(pm, oracle.URL, nil)
	job, err := s.SubmitJob(JobSubmitRequest{
		ClientID:        "c1",
		CPURequired:     4,
		MemoryMB:        2048,
		MaxDurationSecs: 3600,
		MaxPricePerHour: 1.0,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if job.Status != JobStatusAssigned {
		t.Fatalf("expected assigned status, got %s", job.Status)
	}
	if job.AssignedProvider == "" {
		t.Fatal("expected assigned provider")
	}
}

func TestSubmitJobWithNoEligibleProviderReturnsPending(t *testing.T) {
	pm := NewProviderManager()
	pm.Register(Provider{ID: "small", CPU: 1, MemoryMB: 512, PricePerHour: 10, Reputation: 90})

	s := NewScheduler(pm, "http://127.0.0.1:0", nil)
	job, err := s.SubmitJob(JobSubmitRequest{
		ClientID:        "c1",
		CPURequired:     8,
		MemoryMB:        16384,
		MaxDurationSecs: 3600,
		MaxPricePerHour: 1.0,
	})
	if err == nil {
		t.Fatal("expected error when no provider is eligible")
	}
	if job.Status != JobStatusPending {
		t.Fatalf("expected pending status, got %s", job.Status)
	}
}

func TestHandleSubmitJobValidationAndAccepted(t *testing.T) {
	pm := NewProviderManager()
	s := NewScheduler(pm, "http://127.0.0.1:0", nil)

	badReq := httptest.NewRequest(http.MethodPost, "/jobs", strings.NewReader(`{"cpuRequired":0,"memoryMB":0}`))
	badRec := httptest.NewRecorder()
	s.HandleSubmitJob(badRec, badReq)
	if badRec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 for invalid payload, got %d", badRec.Code)
	}

	validReq := httptest.NewRequest(http.MethodPost, "/jobs", strings.NewReader(`{"clientId":"c1","cpuRequired":2,"memoryMB":2048,"maxDurationSecs":60,"maxPricePerHour":1}`))
	validRec := httptest.NewRecorder()
	s.HandleSubmitJob(validRec, validReq)
	if validRec.Code != http.StatusAccepted {
		t.Fatalf("expected 202 when no providers can be assigned, got %d", validRec.Code)
	}
}

func TestHandleGetJobFoundAndNotFound(t *testing.T) {
	pm := NewProviderManager()
	pm.Register(Provider{ID: "p1", CPU: 8, MemoryMB: 8192, PricePerHour: 0.2, Reputation: 90})

	s := NewScheduler(pm, "http://127.0.0.1:0", nil)
	job, err := s.SubmitJob(JobSubmitRequest{
		ClientID:        "c1",
		CPURequired:     1,
		MemoryMB:        512,
		MaxDurationSecs: 60,
		MaxPricePerHour: 1,
	})
	if err != nil {
		t.Fatalf("submit job: %v", err)
	}

	foundReq := httptest.NewRequest(http.MethodGet, "/jobs/"+job.ID, nil)
	foundReq.SetPathValue("id", job.ID)
	foundRec := httptest.NewRecorder()
	s.HandleGetJob(foundRec, foundReq)
	if foundRec.Code != http.StatusOK {
		t.Fatalf("expected 200 for existing job, got %d", foundRec.Code)
	}

	notFoundReq := httptest.NewRequest(http.MethodGet, "/jobs/missing", nil)
	notFoundReq.SetPathValue("id", "missing")
	notFoundRec := httptest.NewRecorder()
	s.HandleGetJob(notFoundRec, notFoundReq)
	if notFoundRec.Code != http.StatusNotFound {
		t.Fatalf("expected 404 for missing job, got %d", notFoundRec.Code)
	}
}

func TestHandleMetricsIncludesJobAndProviderCounts(t *testing.T) {
	pm := NewProviderManager()
	pm.Register(Provider{ID: "p1", CPU: 4, MemoryMB: 4096, PricePerHour: 0.2, Reputation: 80})
	pm.Register(Provider{ID: "p2", CPU: 4, MemoryMB: 4096, PricePerHour: 0.3, Reputation: 70})
	_ = pm.Deregister("p2")

	s := NewScheduler(pm, "http://127.0.0.1:0", nil)
	s.jobs["j1"] = &Job{ID: "j1", Status: JobStatusPending}
	s.jobs["j2"] = &Job{ID: "j2", Status: JobStatusAssigned}

	req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	rec := httptest.NewRecorder()
	s.HandleMetrics(rec, req)

	body := rec.Body.String()
	if !strings.Contains(body, "nimbusx_scheduler_total_jobs_total 2") {
		t.Fatalf("expected total job metric in response, got:\n%s", body)
	}
	if !strings.Contains(body, "nimbusx_scheduler_pending_jobs 1") {
		t.Fatalf("expected pending metric in response, got:\n%s", body)
	}
	if !strings.Contains(body, "nimbusx_scheduler_assigned_jobs 1") {
		t.Fatalf("expected assigned metric in response, got:\n%s", body)
	}
	if !strings.Contains(body, "nimbusx_scheduler_active_providers 1") {
		t.Fatalf("expected active provider metric in response, got:\n%s", body)
	}
}
