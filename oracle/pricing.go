package main

import (
	"log"
	"math/rand"
)

// CloudPriceFeed is a normalised pricing record from a cloud provider or mock feed.
type CloudPriceFeed struct {
	PricePerHour float64
	Source       string
}

// FetchAllPrices returns pricing data for all known providers.
// In production this would call AWS Pricing API, GCP Cloud Billing API, etc.
// For now it uses realistic mock values with slight random jitter to simulate drift.
func FetchAllPrices() map[string]CloudPriceFeed {
	return map[string]CloudPriceFeed{
		// AWS-style providers
		"aws-t3-medium":  {PricePerHour: jitter(0.0416), Source: "aws"},
		"aws-c5-xlarge":  {PricePerHour: jitter(0.1700), Source: "aws"},
		"aws-m5-2xlarge": {PricePerHour: jitter(0.3840), Source: "aws"},

		// GCP-style providers
		"gcp-n1-standard-2": {PricePerHour: jitter(0.0950), Source: "gcp"},
		"gcp-n2-standard-4": {PricePerHour: jitter(0.1900), Source: "gcp"},
		"gcp-c2-standard-8": {PricePerHour: jitter(0.4100), Source: "gcp"},

		// Generic compute nodes (used in tests/simulation)
		"node-1": {PricePerHour: jitter(0.05), Source: "mock"},
		"node-2": {PricePerHour: jitter(0.08), Source: "mock"},
		"node-3": {PricePerHour: jitter(0.12), Source: "mock"},
	}
}

// SeedMockPrices populates the oracle with initial mock data on startup.
func (o *Oracle) SeedMockPrices() {
	prices := FetchAllPrices()
	for id, p := range prices {
		o.SetPrice(id, p.PricePerHour, p.Source)
	}
	log.Printf("[oracle] Seeded %d mock price entries", len(prices))
}

// jitter adds ±5% random variance to simulate real price fluctuation.
func jitter(base float64) float64 {
	variance := base * 0.05
	return base + (rand.Float64()*2-1)*variance
}
