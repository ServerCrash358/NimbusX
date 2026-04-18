package main

import (
	"math"
	"testing"
)

func TestScoreProvidersUsesOracleOverrideAndRanks(t *testing.T) {
	job := &Job{CPURequired: 4, MemoryMB: 4096, MaxPricePerHour: 1.0}
	p1 := &Provider{ID: "p1", CPU: 8, MemoryMB: 8192, PricePerHour: 0.9, Reputation: 90}
	p2 := &Provider{ID: "p2", CPU: 8, MemoryMB: 8192, PricePerHour: 0.5, Reputation: 80}

	oracle := func(id string) float64 {
		if id == "p1" {
			return 0.1
		}
		return 0
	}

	results := ScoreProviders([]*Provider{p1, p2}, job, oracle)
	if len(results) != 2 {
		t.Fatalf("expected 2 results, got %d", len(results))
	}
	if results[0].Provider.ID != "p1" {
		t.Fatalf("expected p1 to score higher with oracle price override, got %s", results[0].Provider.ID)
	}
}

func TestComputeScoreClampsAndExpectedValue(t *testing.T) {
	job := &Job{CPURequired: 4, MemoryMB: 4096, MaxPricePerHour: 1.0}
	p := &Provider{ID: "p1", CPU: 8, MemoryMB: 8192, PricePerHour: 2.0, Reputation: 100}

	// costScore clamps to 0 when provider price > max.
	// reputationScore = 1.0
	// headroom = ((8-4)/8 + (8192-4096)/8192) / 2 = 0.5
	// final = 0.40*0 + 0.35*1 + 0.25*0.5 = 0.475
	got := computeScore(p, job, nil)
	if math.Abs(got-0.475) > 1e-9 {
		t.Fatalf("expected score 0.475, got %f", got)
	}
}
