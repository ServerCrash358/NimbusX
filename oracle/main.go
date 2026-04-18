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
	port := os.Getenv("ORACLE_PORT")
	if port == "" {
		port = "7070"
	}

	oracle := NewOracle()

	// Seed with some initial price data so the scheduler can start immediately
	oracle.SeedMockPrices()

	// Start background price refresh loop (every 60 seconds)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go oracle.StartRefreshLoop(ctx, 60*time.Second)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /price/{providerID}", oracle.HandleGetPrice)
	mux.HandleFunc("GET /prices", oracle.HandleGetAllPrices)
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Printf("[oracle] Listening on :%s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[oracle] Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("[oracle] Shutting down...")
	cancel()
	shutCtx, shutCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutCancel()
	_ = srv.Shutdown(shutCtx)
}
