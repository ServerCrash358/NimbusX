package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os/exec"
)

type Job struct {
    Image   string `json:"image"`
    Command string `json:"command"`
    JobID   string `json:"job_id"`
}

func RunJobHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "POST only", http.StatusMethodNotAllowed)
        return
    }

    var job Job
    if err := json.NewDecoder(r.Body).Decode(&job); err != nil {
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }

    fmt.Printf("Running job %s: docker run %s %s\n", job.JobID, job.Image, job.Command)

    // Run the docker container
    cmd := exec.Command("docker", "run", "--rm", job.Image, job.Command)
    output, err := cmd.CombinedOutput()
    if err != nil {
        jobsFailed.Inc()
        http.Error(w, fmt.Sprintf("job failed: %v\n%s", err, output), 500)
        return
    }

    jobsRun.Inc()
    w.WriteHeader(http.StatusOK)
    w.Write(output)
}
