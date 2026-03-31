package main

import (
    "encoding/json"
    "net/http"
    "runtime"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
)

var (
    jobsRun = promauto.NewCounter(prometheus.CounterOpts{
        Name: "nimbusX_jobs_run_total",
        Help: "Total number of jobs executed",
    })
    jobsFailed = promauto.NewCounter(prometheus.CounterOpts{
        Name: "nimbusX_jobs_failed_total",
        Help: "Total number of failed jobs",
    })
    cpuUsage = promauto.NewGauge(prometheus.GaugeOpts{
        Name: "nimbusX_cpu_goroutines",
        Help: "Active goroutines (proxy for CPU load)",
    })
)

type HardwareInfo struct {
    CPUs   int    `json:"cpus"`
    OS     string `json:"os"`
    AgentVersion string `json:"agent_version"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    cpuUsage.Set(float64(runtime.NumGoroutine()))
    info := HardwareInfo{
        CPUs:         runtime.NumCPU(),
        OS:           runtime.GOOS,
        AgentVersion: "0.1.0",
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(info)
}
