package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    fmt.Println("NimbusX Node Agent starting...")

    // Health check endpoint (scheduler pings this)
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"status":"ok"}`))
    })

    // Job execution endpoint (scheduler sends jobs here)
    http.HandleFunc("/run-job", RunJobHandler)

    // Hardware registration endpoint
    http.HandleFunc("/register", RegisterHandler)

    // Prometheus metrics endpoint
    http.Handle("/metrics", promhttp.Handler())

    log.Println("Agent listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
