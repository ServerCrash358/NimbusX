package main

import (
	"log"
	"net/http"
)

// StartMetricsServer starts a simple HTTP server to expose Prometheus-compatible metrics.
// In a full implementation, we would use the github.com/prometheus/client_golang library.
func StartMetricsServer(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		// Mock Prometheus metrics response
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("# HELP nimbusx_node_cpu_usage CPU usage percentage\n"))
		w.Write([]byte("# TYPE nimbusx_node_cpu_usage gauge\n"))
		w.Write([]byte("nimbusx_node_cpu_usage 15.2\n"))
		
		w.Write([]byte("# HELP nimbusx_node_memory_mb Memory usage in MB\n"))
		w.Write([]byte("# TYPE nimbusx_node_memory_mb gauge\n"))
		w.Write([]byte("nimbusx_node_memory_mb 2048\n"))
	})

	go func() {
		log.Printf("[metrics] Listening on :%s/metrics", port)
		if err := http.ListenAndServe(":"+port, mux); err != nil {
			log.Fatalf("[metrics] Server failed: %v", err)
		}
	}()
}
