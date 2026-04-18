package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
)

// Config holds agent configuration
type Config struct {
	NodeID       string
	GatewayURL   string
	CPU          int
	MemoryMB     int
	PricePerHour float64
	PollInterval time.Duration
}

// Job represents a workload assigned to this node
type Job struct {
	ID              string  `json:"id"`
	CPURequired     int     `json:"cpuRequired"`
	MemoryMB        int     `json:"memoryMB"`
	MaxDurationSecs int     `json:"maxDurationSecs"`
	Status          string  `json:"status"`
}

func main() {
	cfg := Config{
		NodeID:       getEnv("NODE_ID", "node-"+uuid.New().String()[:8]),
		GatewayURL:   getEnv("API_GATEWAY_URL", "http://localhost:8080"),
		CPU:          4,    // Default 4 cores
		MemoryMB:     8192, // Default 8GB
		PricePerHour: 0.05, // Default $0.05/hr
		PollInterval: 5 * time.Second,
	}

	log.Printf("[node-agent] Starting agent %s", cfg.NodeID)
	log.Printf("[node-agent] Gateway: %s", cfg.GatewayURL)

	// Register with Gateway
	if err := registerNode(cfg); err != nil {
		log.Fatalf("[node-agent] Failed to register: %v", err)
	}
	log.Printf("[node-agent] Successfully registered provider")

	// Initialize Docker executor
	executor, err := NewDockerExecutor()
	if err != nil {
		log.Fatalf("[node-agent] Docker executor init failed: %v", err)
	}

	// Initialize metrics collector
	StartMetricsServer("9091") // Expose /metrics for Prometheus

	// Start polling loop
	stop := make(chan struct{})
	go func() {
		ticker := time.NewTicker(cfg.PollInterval)
		defer ticker.Stop()
		for {
			select {
			case <-stop:
				return
			case <-ticker.C:
				pollJobs(cfg, executor)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("[node-agent] Shutting down...")
	close(stop)
	time.Sleep(1 * time.Second)
}

func registerNode(cfg Config) error {
	payload := map[string]interface{}{
		"id":           cfg.NodeID,
		"address":      "0x" + cfg.NodeID, // Mock Ethereum address for agent
		"cpu":          cfg.CPU,
		"memoryMB":     cfg.MemoryMB,
		"pricePerHour": cfg.PricePerHour,
	}
	data, _ := json.Marshal(payload)

	resp, err := http.Post(cfg.GatewayURL+"/providers/register", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}
	return nil
}

// pollJobs asks the gateway if any jobs are assigned to this node
// For this simple mock implementation, the node asks the scheduler if its ID is in any assigned jobs
// Realistically, the scheduler would push to the node, or the node would poll a specific endpoint.
func pollJobs(cfg Config, executor *DockerExecutor) {
	// For simulation, we pretend to poll here.
	// In reality, we would call GET /jobs?assignedTo=cfg.NodeID
	// Currently the Gateway doesn't have a specific "list jobs by provider" endpoint,
	// so the node-agent would normally just wait for the Scheduler to contact it,
	// or the Gateway would need a new endpoint.
	// For now, we will just log the polling heartbeat.
	log.Printf("[node-agent] Heartbeat: Node %s is active and waiting for jobs...", cfg.NodeID)
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
