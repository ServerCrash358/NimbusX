# Variables for easy updates
IMAGE_NAME = nimbusx-node-agent:latest
DEPLOY_FILE = infra/k8s/node-agent-deployement.yaml
AGENT_DIR = cloud-cost-optimizer/node-agent

.PHONY: all start load deploy watch clean

# The default command when you just type 'make'
all: start load deploy watch

setup : start load deploy

# 1. Start minikube
start:
	@echo "Starting Minikube..."
	minikube start
	minikube status

# 2. Reload the image
load:
	@echo "Loading image $(IMAGE_NAME) into Minikube..."
	minikube image load $(IMAGE_NAME)
	minikube status

# 3. Redeploy the pods
deploy:
	@echo "Applying Kubernetes manifests..."
	kubectl apply -f $(DEPLOY_FILE)

# 4. Watch pods
watch:
	@echo "Watching pods (Ctrl+C to stop)..."
	kubectl get pods -w

# Bonus: A way to quickly delete and reset
clean:
	kubectl delete -f $(DEPLOY_FILE)
	
#stop it completely
stop:
	kubectl delete -f $(DEPLOY_FILE)
	minikube stop
	minikube status

all: setup
