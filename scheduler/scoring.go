package main

// ScoreProviders ranks providers for a given job using a weighted multi-factor formula.
//
// Scoring factors and weights:
//   - Cost efficiency  (40%) – how much cheaper the provider is vs the job's max price
//   - Reputation       (35%) – normalised 0-100 provider reputation score
//   - Resource headroom(25%) – extra CPU/memory the provider has beyond job requirements
//
// An optional oracle price function can override the provider's self-reported price
// for a fairer market-rate comparison.
func ScoreProviders(providers []*Provider, job *Job, oraclePrice func(string) float64) []ScoreResult {
	results := make([]ScoreResult, 0, len(providers))

	for _, p := range providers {
		score := computeScore(p, job, oraclePrice)
		results = append(results, ScoreResult{Provider: p, Score: score})
	}

	return results
}

func computeScore(p *Provider, job *Job, oraclePrice func(string) float64) float64 {
	const (
		wCost       = 0.40
		wReputation = 0.35
		wHeadroom   = 0.25
	)

	// ── Cost Efficiency ────────────────────────────────────────────────────
	// Use oracle market price if available, fall back to self-reported price.
	effectivePrice := p.PricePerHour
	if oraclePrice != nil {
		if op := oraclePrice(p.ID); op > 0 {
			effectivePrice = op
		}
	}

	// Score is higher when provider price is well below the job's max price.
	// Clamp to [0,1]: 0 means at max price, 1 means free.
	costScore := 0.0
	if job.MaxPricePerHour > 0 {
		saving := (job.MaxPricePerHour - effectivePrice) / job.MaxPricePerHour
		if saving > 1 {
			saving = 1
		}
		if saving < 0 {
			saving = 0
		}
		costScore = saving
	}

	// ── Reputation ─────────────────────────────────────────────────────────
	reputationScore := float64(p.Reputation) / 100.0 // normalise to [0,1]

	// ── Resource Headroom ──────────────────────────────────────────────────
	// Extra CPU and memory beyond what the job needs, normalised to the
	// total available so providers aren't penalised for being large.
	cpuHeadroom := 0.0
	if p.CPU > 0 && job.CPURequired > 0 {
		excess := float64(p.CPU-job.CPURequired) / float64(p.CPU)
		if excess < 0 {
			excess = 0
		}
		cpuHeadroom = excess
	}

	memHeadroom := 0.0
	if p.MemoryMB > 0 && job.MemoryMB > 0 {
		excess := float64(p.MemoryMB-job.MemoryMB) / float64(p.MemoryMB)
		if excess < 0 {
			excess = 0
		}
		memHeadroom = excess
	}

	headroomScore := (cpuHeadroom + memHeadroom) / 2.0

	// ── Final Weighted Score ───────────────────────────────────────────────
	return (wCost * costScore) +
		(wReputation * reputationScore) +
		(wHeadroom * headroomScore)
}
