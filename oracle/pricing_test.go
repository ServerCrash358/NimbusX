package main

import "testing"

func TestFetchAllPricesContainsExpectedProvidersWithPositiveValues(t *testing.T) {
	prices := FetchAllPrices()
	if len(prices) == 0 {
		t.Fatal("expected non-empty price map")
	}

	required := []string{"aws-t3-medium", "gcp-n1-standard-2", "node-1"}
	for _, id := range required {
		p, ok := prices[id]
		if !ok {
			t.Fatalf("expected provider %s in feed", id)
		}
		if p.PricePerHour <= 0 {
			t.Fatalf("expected positive price for %s, got %f", id, p.PricePerHour)
		}
		if p.Source == "" {
			t.Fatalf("expected non-empty source for %s", id)
		}
	}
}

func TestJitterIsWithinFivePercent(t *testing.T) {
	base := 100.0
	min := base * 0.95
	max := base * 1.05

	for i := 0; i < 1000; i++ {
		v := jitter(base)
		if v < min || v > max {
			t.Fatalf("jitter value %f out of bounds [%f, %f]", v, min, max)
		}
	}
}
