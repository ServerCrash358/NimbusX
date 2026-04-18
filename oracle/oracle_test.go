package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOracleSetGetAndCopySemantics(t *testing.T) {
	o := NewOracle()
	o.SetPrice("node-1", 0.05, "mock")

	entry, ok := o.GetPrice("node-1")
	if !ok {
		t.Fatal("expected price entry")
	}
	entry.Price = 999

	entry2, ok := o.GetPrice("node-1")
	if !ok {
		t.Fatal("expected price entry")
	}
	if entry2.Price == 999 {
		t.Fatal("expected GetPrice to return a copy")
	}
}

func TestHandleGetPriceFoundAndNotFound(t *testing.T) {
	o := NewOracle()
	o.SetPrice("node-1", 0.05, "mock")

	foundReq := httptest.NewRequest(http.MethodGet, "/price/node-1", nil)
	foundReq.SetPathValue("providerID", "node-1")
	foundRec := httptest.NewRecorder()
	o.HandleGetPrice(foundRec, foundReq)
	if foundRec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", foundRec.Code)
	}

	var payload PriceEntry
	if err := json.Unmarshal(foundRec.Body.Bytes(), &payload); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if payload.ProviderID != "node-1" {
		t.Fatalf("expected provider id node-1, got %s", payload.ProviderID)
	}

	missReq := httptest.NewRequest(http.MethodGet, "/price/missing", nil)
	missReq.SetPathValue("providerID", "missing")
	missRec := httptest.NewRecorder()
	o.HandleGetPrice(missRec, missReq)
	if missRec.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", missRec.Code)
	}
}

func TestHandleGetAllPricesReturnsArray(t *testing.T) {
	o := NewOracle()

	req := httptest.NewRequest(http.MethodGet, "/prices", nil)
	rec := httptest.NewRecorder()
	o.HandleGetAllPrices(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var entries []PriceEntry
	if err := json.Unmarshal(rec.Body.Bytes(), &entries); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(entries) != 0 {
		t.Fatalf("expected empty array, got %d entries", len(entries))
	}
}
