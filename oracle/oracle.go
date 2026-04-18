package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

// PriceEntry holds a normalised price for a given provider.
type PriceEntry struct {
	ProviderID  string    `json:"providerId"`
	Price       float64   `json:"price"`       // USD per hour
	Source      string    `json:"source"`      // "aws", "gcp", "mock", etc.
	LastUpdated time.Time `json:"lastUpdated"`
}

// Oracle aggregates and normalises pricing data from multiple sources.
type Oracle struct {
	mu     sync.RWMutex
	prices map[string]*PriceEntry
}

func NewOracle() *Oracle {
	return &Oracle{
		prices: make(map[string]*PriceEntry),
	}
}

// SetPrice updates the stored price for a provider.
func (o *Oracle) SetPrice(providerID string, price float64, source string) {
	o.mu.Lock()
	defer o.mu.Unlock()

	o.prices[providerID] = &PriceEntry{
		ProviderID:  providerID,
		Price:       price,
		Source:      source,
		LastUpdated: time.Now(),
	}
}

// GetPrice retrieves the current price for a provider.
func (o *Oracle) GetPrice(providerID string) (*PriceEntry, bool) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	entry, ok := o.prices[providerID]
	if !ok {
		return nil, false
	}
	copy := *entry
	return &copy, true
}

// GetAllPrices returns a snapshot of all current prices.
func (o *Oracle) GetAllPrices() []*PriceEntry {
	o.mu.RLock()
	defer o.mu.RUnlock()

	result := make([]*PriceEntry, 0, len(o.prices))
	for _, e := range o.prices {
		copy := *e
		result = append(result, &copy)
	}
	return result
}

// StartRefreshLoop periodically re-fetches prices from external sources.
func (o *Oracle) StartRefreshLoop(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			o.refresh()
		}
	}
}

// refresh pulls updated prices from all configured pricing feeds.
func (o *Oracle) refresh() {
	log.Println("[oracle] Refreshing prices...")
	prices := FetchAllPrices()
	for id, price := range prices {
		o.SetPrice(id, price.PricePerHour, price.Source)
	}
	log.Printf("[oracle] Updated %d provider prices", len(prices))
}

// ── HTTP Handlers ──────────────────────────────────────────────────────────

// HandleGetPrice handles GET /price/{providerID}
func (o *Oracle) HandleGetPrice(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("providerID")
	entry, ok := o.GetPrice(id)
	if !ok {
		http.Error(w, "price not found for provider", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(entry)
}

// HandleGetAllPrices handles GET /prices
func (o *Oracle) HandleGetAllPrices(w http.ResponseWriter, r *http.Request) {
	prices := o.GetAllPrices()
	if prices == nil {
		prices = []*PriceEntry{}
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(prices)
}
