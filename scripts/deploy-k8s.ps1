# NimbusX Kubernetes Deployment Script (PowerShell)
# Deploys the full stack to the current kubectl context.

Write-Host "🚀 Deploying NimbusX to Kubernetes..." -ForegroundColor Cyan

# 1. Config
kubectl apply -f k8s/config.yaml

# 2. Blockchain (Hardhat)
kubectl apply -f k8s/hardhat.yaml

# 3. Oracle
kubectl apply -f k8s/oracle.yaml

# 4. Scheduler
kubectl apply -f k8s/scheduler.yaml

# 5. API Gateway
kubectl apply -f k8s/api.yaml

# 6. Node Agent
kubectl apply -f k8s/node-agent.yaml

Write-Host "⏳ Waiting for deployments to stabilize..." -ForegroundColor Yellow
kubectl rollout status deployment/hardhat
kubectl rollout status deployment/oracle
kubectl rollout status deployment/scheduler
kubectl rollout status deployment/api
kubectl rollout status deployment/node-agent

Write-Host "✅ NimbusX is now running in Kubernetes!" -ForegroundColor Green
Write-Host "API Gateway is available at: http://localhost:8080 (if port-forwarded)"
