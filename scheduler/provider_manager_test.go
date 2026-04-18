package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProviderManagerRegisterGetAndCopySemantics(t *testing.T) {
	pm := NewProviderManager()
	pm.Register(Provider{ID: "p1", CPU: 4, MemoryMB: 8192, PricePerHour: 0.1, Reputation: 80})

	if pm.GetCount() != 1 {
		t.Fatalf("expected 1 active provider, got %d", pm.GetCount())
	}

	got, ok := pm.Get("p1")
	if !ok {
		t.Fatal("expected provider to exist")
	}
	got.CPU = 999

	gotAgain, ok := pm.Get("p1")
	if !ok {
		t.Fatal("expected provider to exist")
	}
	if gotAgain.CPU == 999 {
		t.Fatal("expected Get to return a copy, not original pointer")
	}

	active := pm.GetActive()
	if len(active) != 1 {
		t.Fatalf("expected one active provider, got %d", len(active))
	}
	active[0].MemoryMB = 1

	active2 := pm.GetActive()
	if active2[0].MemoryMB == 1 {
		t.Fatal("expected GetActive to return copies, not original pointers")
	}
}

func TestProviderManagerDeregisterAndUpdateReputationClamp(t *testing.T) {
	pm := NewProviderManager()
	pm.Register(Provider{ID: "p1", CPU: 2, MemoryMB: 2048, PricePerHour: 0.05, Reputation: 50})

	if !pm.Deregister("p1") {
		t.Fatal("expected deregister to return true for existing provider")
	}
	if pm.GetCount() != 0 {
		t.Fatalf("expected no active providers, got %d", pm.GetCount())
	}
	if pm.Deregister("missing") {
		t.Fatal("expected deregister to return false for missing provider")
	}

	pm.UpdateReputation("p1", 1000)
	p, _ := pm.Get("p1")
	if p.Reputation != 100 {
		t.Fatalf("expected reputation clamp to 100, got %d", p.Reputation)
	}

	pm.UpdateReputation("p1", -1000)
	p, _ = pm.Get("p1")
	if p.Reputation != 0 {
		t.Fatalf("expected reputation clamp to 0, got %d", p.Reputation)
	}
}

func TestHandleRegisterProviderValidationAndSuccess(t *testing.T) {
	pm := NewProviderManager()

	invalidReq := httptest.NewRequest(http.MethodPost, "/providers/register", strings.NewReader(`{"id":"","cpu":0,"memoryMB":0}`))
	invalidRec := httptest.NewRecorder()
	pm.HandleRegisterProvider(invalidRec, invalidReq)
	if invalidRec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400 for invalid payload, got %d", invalidRec.Code)
	}

	validReq := httptest.NewRequest(http.MethodPost, "/providers/register", strings.NewReader(`{"id":"p1","cpu":4,"memoryMB":8192,"pricePerHour":0.1}`))
	validRec := httptest.NewRecorder()
	pm.HandleRegisterProvider(validRec, validReq)
	if validRec.Code != http.StatusCreated {
		t.Fatalf("expected 201 for valid payload, got %d", validRec.Code)
	}

	var body map[string]string
	if err := json.Unmarshal(validRec.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if body["id"] != "p1" {
		t.Fatalf("expected id p1, got %q", body["id"])
	}
}

func TestHandleListProvidersReturnsEmptyArrayNotNull(t *testing.T) {
	pm := NewProviderManager()
	req := httptest.NewRequest(http.MethodGet, "/providers", nil)
	rec := httptest.NewRecorder()

	pm.HandleListProviders(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var providers []Provider
	if err := json.Unmarshal(rec.Body.Bytes(), &providers); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(providers) != 0 {
		t.Fatalf("expected empty provider list, got %d", len(providers))
	}
}
