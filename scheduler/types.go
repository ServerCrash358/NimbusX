package main

import "time"

// JobStatus represents the lifecycle state of a compute job.
type JobStatus string

const (
	JobStatusPending   JobStatus = "pending"
	JobStatusAssigned  JobStatus = "assigned"
	JobStatusRunning   JobStatus = "running"
	JobStatusCompleted JobStatus = "completed"
	JobStatusFailed    JobStatus = "failed"
)

// Job is a compute workload submitted by a client.
type Job struct {
	ID              string    `json:"id"`
	ClientID        string    `json:"clientId"`
	CPURequired     int       `json:"cpuRequired"`     // cores
	MemoryMB        int       `json:"memoryMB"`        // megabytes
	MaxDurationSecs int       `json:"maxDurationSecs"` // seconds
	MaxPricePerHour float64   `json:"maxPricePerHour"` // USD/hour
	Status          JobStatus `json:"status"`
	AssignedProvider string   `json:"assignedProvider,omitempty"`
	SubmittedAt     time.Time `json:"submittedAt"`
	StartedAt       *time.Time `json:"startedAt,omitempty"`
	CompletedAt     *time.Time `json:"completedAt,omitempty"`
}

// Provider represents a compute node registered with the scheduler.
type Provider struct {
	ID           string  `json:"id"`
	Address      string  `json:"address"`  // Ethereum address
	CPU          int     `json:"cpu"`      // total cores
	MemoryMB     int     `json:"memoryMB"` // total MB
	PricePerHour float64 `json:"pricePerHour"` // USD/hour
	Reputation   int     `json:"reputation"`   // 0-100
	Active       bool    `json:"active"`
	RegisteredAt time.Time `json:"registeredAt"`
}

// ProviderRegisterRequest is the payload for provider self-registration.
type ProviderRegisterRequest struct {
	ID           string  `json:"id"`
	Address      string  `json:"address"`
	CPU          int     `json:"cpu"`
	MemoryMB     int     `json:"memoryMB"`
	PricePerHour float64 `json:"pricePerHour"`
}

// JobSubmitRequest is the payload sent by clients to submit a job.
type JobSubmitRequest struct {
	ClientID        string  `json:"clientId"`
	CPURequired     int     `json:"cpuRequired"`
	MemoryMB        int     `json:"memoryMB"`
	MaxDurationSecs int     `json:"maxDurationSecs"`
	MaxPricePerHour float64 `json:"maxPricePerHour"`
}

// ScoreResult pairs a provider with its computed priority score.
type ScoreResult struct {
	Provider *Provider
	Score    float64
}
