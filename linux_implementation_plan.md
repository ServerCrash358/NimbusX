# NimbusX – Linux Developer Implementation Plan

> You are the **Linux Developer**. Work is divided into 4 steps per the spec.
> Each step below is a self-contained push-ready batch.
> I will complete one step at a time and tell you when to push.

---

## Progress Tracker

| Step | What | Status |
|------|------|--------|
| 1 | Node Agent (`node-agent/`) & CI/CD pipeline | 🔄 In Progress |
| 2 | Monitoring Stack (`monitoring/`) | ⬜ Not Started |
| 3 | Kubernetes Cluster manifests (`k8s/`) | ⬜ Not Started |
| 4 | Infrastructure simulation & integration | ⬜ Not Started |

---

## Step 1 – Node Agent & CI/CD

**Goal:** Build the compute node software that registers with the scheduler, executes Docker containers, and reports telemetry. Setup basic CI/CD.

### Files to create/modify:

| File | Purpose |
|------|---------|
| `node-agent/main.go` | Entry point, registration, and polling logic |
| `node-agent/executor.go` | Spawns and manages Docker containers for jobs |
| `node-agent/metrics.go` | Collects CPU/memory telemetry (for Prometheus) |
| `node-agent/Dockerfile` | Dockerfile for the agent |
| `.github/workflows/ci.yml` | GitHub Actions workflow for build/test |
| `go.mod` | Add Docker SDK dependencies |

---

## Step 2 – Monitoring Stack

**Goal:** Implement observability using Prometheus and Grafana.

### Files to create:

| File | Purpose |
|------|---------|
| `monitoring/prometheus.yml` | Scrape configurations for scheduler and nodes |
| `monitoring/grafana/provisioning/` | Auto-provisioned data sources and dashboards |
| `monitoring/grafana/dashboards/` | Dashboard JSONs for Node metrics and Job stats |

---

## Step 3 – Kubernetes Infrastructure

**Goal:** Create K8s manifests for the entire application stack.

### Files to create:

| File | Purpose |
|------|---------|
| `k8s/api.yaml` | API Gateway deployment and service |
| `k8s/scheduler.yaml` | Scheduler deployment and service |
| `k8s/oracle.yaml` | Price Oracle deployment and service |
| `k8s/hardhat.yaml` | Local simulated blockchain network |
| `k8s/node-agent.yaml` | DaemonSet or Deployments for the compute nodes |

---

## Step 4 – Cluster Simulation

**Goal:** Scripting to deploy everything to a local cluster (e.g., Minikube/Docker Desktop) and run end-to-end load tests.

### Files to create:

| File | Purpose |
|------|---------|
| `scripts/deploy-k8s.sh` | Shell script to apply all K8s manifests in order |
| `scripts/simulate-jobs.sh` | Shell script to spam the API Gateway with jobs |

---

## Repository Structure (Linux Dev Additions)

```
NimbusX/
├── node-agent/
│   ├── main.go
│   ├── executor.go
│   ├── metrics.go
│   └── Dockerfile
├── monitoring/
│   ├── prometheus.yml
│   └── grafana/
├── k8s/
│   ├── api.yaml
│   ├── scheduler.yaml
│   ├── oracle.yaml
│   ├── hardhat.yaml
│   └── node-agent.yaml
└── .github/
    └── workflows/
        └── ci.yml
```
