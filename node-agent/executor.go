package main

import (
	"context"
	"log"
	"time"
)

// DockerExecutor manages spinning up and monitoring Docker containers for jobs.
// NOTE: Mocked for this phase of development.
type DockerExecutor struct {
}

// NewDockerExecutor creates a new executor that connects to the local Docker daemon.
func NewDockerExecutor() (*DockerExecutor, error) {
	return &DockerExecutor{}, nil
}

// ExecuteJob pulls the requested image and runs a container for the given job.
func (e *DockerExecutor) ExecuteJob(ctx context.Context, job Job, imageName string) error {
	log.Printf("[executor] MOCK Pulling image %s for job %s...", imageName, job.ID)
	
	// Simulate pull delay
	time.Sleep(2 * time.Second)

	log.Printf("[executor] MOCK Starting container for job %s...", job.ID)
	
	// Simulate execution delay
	time.Sleep(1 * time.Second)

	log.Printf("[executor] MOCK Job %s is now running in container (simulated)", job.ID)

	return nil
}
