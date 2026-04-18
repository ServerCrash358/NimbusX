package main

import (
	"fmt"
	"net/http"
)

// HandleMetrics returns simple Prometheus-formatted metrics for the scheduler.
func (s *Scheduler) HandleMetrics(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	totalJobs := len(s.jobs)
	pending := 0
	assigned := 0
	for _, j := range s.jobs {
		if j.Status == JobStatusPending {
			pending++
		} else if j.Status == JobStatusAssigned {
			assigned++
		}
	}
	s.mu.RUnlock()

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "# HELP nimbusx_scheduler_total_jobs_total Total jobs submitted\n")
	fmt.Fprintf(w, "# TYPE nimbusx_scheduler_total_jobs_total counter\n")
	fmt.Fprintf(w, "nimbusx_scheduler_total_jobs_total %d\n", totalJobs)

	fmt.Fprintf(w, "# HELP nimbusx_scheduler_pending_jobs Total jobs in pending state\n")
	fmt.Fprintf(w, "# TYPE nimbusx_scheduler_pending_jobs gauge\n")
	fmt.Fprintf(w, "nimbusx_scheduler_pending_jobs %d\n", pending)

	fmt.Fprintf(w, "# HELP nimbusx_scheduler_assigned_jobs Total jobs in assigned state\n")
	fmt.Fprintf(w, "# TYPE nimbusx_scheduler_assigned_jobs gauge\n")
	fmt.Fprintf(w, "nimbusx_scheduler_assigned_jobs %d\n", assigned)

	fmt.Fprintf(w, "# HELP nimbusx_scheduler_active_providers Total registered providers\n")
	fmt.Fprintf(w, "# TYPE nimbusx_scheduler_active_providers gauge\n")
	fmt.Fprintf(w, "nimbusx_scheduler_active_providers %d\n", s.pm.GetCount())
}
