#!/bin/bash

# NimbusX Kubernetes Deployment Script
# Deploys the full stack to the current kubectl context.

echo "🚀 Deploying NimbusX to Kubernetes..."

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

echo "⏳ Waiting for deployments to stabilize..."
kubectl rollout status deployment/hardhat
kubectl rollout status deployment/oracle
kubectl rollout status deployment/scheduler
kubectl rollout status deployment/api
kubectl rollout status deployment/node-agent

echo "✅ NimbusX is now running in Kubernetes!"
echo "API Gateway is available at: http://localhost:8080 (if port-forwarded)"
