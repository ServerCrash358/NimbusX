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
	port := os.Getenv("SCHEDULER_PORT")
	if port == "" {
		port = "9090"
	}

	oracleURL := os.Getenv("ORACLE_URL")
	if oracleURL == "" {
		oracleURL = "http://localhost:7070"
	}

	pm := NewProviderManager()

	bcClient, err := NewBlockchainClient()
	if err != nil {
		log.Printf("[scheduler] WARNING: Blockchain integration disabled: %v", err)
	} else {
		log.Printf("[scheduler] Blockchain integration enabled")
	}

	sc := NewScheduler(pm, oracleURL, bcClient)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /jobs", sc.HandleSubmitJob)
	mux.HandleFunc("GET /providers", pm.HandleListProviders)
	mux.HandleFunc("GET /jobs/{id}", sc.HandleGetJob)
	mux.HandleFunc("POST /providers/register", pm.HandleRegisterProvider)
	mux.HandleFunc("GET /metrics", sc.HandleMetrics)

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Printf("[scheduler] Listening on :%s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[scheduler] Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("[scheduler] Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
