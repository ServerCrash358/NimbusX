package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	schedulerURL := os.Getenv("SCHEDULER_URL")
	if schedulerURL == "" {
		schedulerURL = "http://localhost:9090"
	}

	gw := NewGateway(schedulerURL)

	mux := http.NewServeMux()

	// Core NimbusX endpoints
	mux.HandleFunc("POST /submit-job", gw.HandleSubmitJob)
	mux.HandleFunc("GET /providers", gw.HandleGetProviders)
	mux.HandleFunc("GET /job-status", gw.HandleGetJobStatus)

	// Health check
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok","service":"api-gateway"}`))
	})

	// Wrap with middleware
	handler := LoggingMiddleware(CORSMiddleware(mux))

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	go func() {
		log.Printf("[api-gateway] Listening on :%s", port)
		log.Printf("[api-gateway] Proxying to scheduler at %s", schedulerURL)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[api-gateway] Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("[api-gateway] Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
