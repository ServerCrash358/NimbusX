#!/bin/bash

# NimbusX Job Simulation Script
# Submits a batch of jobs to the API Gateway to test the scheduler and node agents.

API_URL="http://localhost:8080"

echo "🧪 Starting Job Simulation..."

for i in {1..5}
do
  echo "Submission $i: Posting job..."
  curl -X POST "$API_URL/jobs" \
    -H "Content-Type: application/json" \
    -d '{
      "clientId": "test-client",
      "cpuRequired": 1,
      "memoryMB": 1024,
      "maxDurationSecs": 300,
      "maxPricePerHour": 0.1
    }'
  echo -e "\n"
  sleep 1
done

echo "✅ Simulation batch complete. Check Grafana (http://localhost:3000) for metrics!"
