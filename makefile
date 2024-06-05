# ==============================================================================
# Define dependencies

GOLANG          := golang:1.22
ALPINE          := alpine:3.19
KIND            := kindest/node:v1.30.0
POSTGRES        := postgres:16.2

KIND_CLUSTER    := harvester-cluster
NAMESPACE       := harvester-system
HARVESTER_APP   := harvester
AUTH_APP        := auth
BASE_IMAGE_NAME := localhost/harvester
VERSION         := 0.0.1
HARVESTER_IMAGE := $(BASE_IMAGE_NAME)/$(HARVESTER_APP):$(VERSION)
METRICS_IMAGE   := $(BASE_IMAGE_NAME)/metrics:$(VERSION)
AUTH_IMAGE      := $(BASE_IMAGE_NAME)/$(AUTH_APP):$(VERSION)

# VERSION       := "0.0.1-$(shell git rev-parse --short HEAD)"

# ==============================================================================
# Install dependencies

dev-gotooling:
	go install github.com/divan/expvarmon@latest
	go install github.com/rakyll/hey@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install golang.org/x/tools/cmd/goimports@latest

dev-brew:
	brew update
	brew list kind || brew install kind
	brew list kubectl || brew install kubectl
	brew list kustomize || brew install kustomize
	brew list pgcli || brew install pgcli
	brew list watch || brew instal watch

dev-docker:
	docker pull $(GOLANG) & \
	docker pull $(ALPINE) & \
	docker pull $(KIND) & \
	docker pull $(POSTGRES) & \
	wait;

# ==============================================================================
# Building containers

build: harvester

harvester:
	docker build \
		-f infrastructure/docker/dockerfile.harvester \
		-t $(HARVESTER_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
		.

# ==============================================================================
# Running from within k8s/kind

dev-up:
	kind create cluster \
		--image $(KIND) \
		--name $(KIND_CLUSTER) \
		--config infrastructure/k8s/dev/kind-config.yaml

	kubectl wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner

	kind load docker-image $(POSTGRES) --name $(KIND_CLUSTER)

dev-load-db:
	kind load docker-image $(POSTGRES) --name $(KIND_CLUSTER)

dev-down:
	kind delete cluster --name $(KIND_CLUSTER)

dev-status-all:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

dev-status:
	watch -n 2 kubectl get pods -o wide --all-namespaces


# ------------------------------------------------------------------------------

dev-load:
	kind load docker-image $(HARVESTER_IMAGE) --name $(KIND_CLUSTER)

dev-apply:
	kustomize build zarf/k8s/dev/database | kubectl apply -f -
	kubectl rollout status --namespace=$(NAMESPACE) --watch --timeout=120s sts/database

	kustomize build infrastructure/k8s/dev/harvester | kubectl apply -f -
	kubectl wait pods --namespace=$(NAMESPACE) --selector app=$(HARVESTER_APP) --timeout=120s --for=condition=Ready

dev-restart:
	kubectl rollout restart deployment $(HARVESTER_APP) --namespace=$(NAMESPACE)

dev-update: build dev-load dev-restart

dev-update-apply: build dev-load dev-apply

dev-logs:
	kubectl logs --namespace=$(NAMESPACE) -l app=$(HARVESTER_APP) --all-containers=true -f --tail=100 --max-log-requests=6 | go run api/cmd/tooling/logfmt/main.go -service=$(HARVESTER_APP)

# ------------------------------------------------------------------------------

dev-logs-init:
	kubectl logs --namespace=$(NAMESPACE) -l app=$(HARVESTER_APP) -f --tail=100 -c init-migrate-seed

dev-describe-deployment:
	kubectl describe deployment --namespace=$(NAMESPACE) $HARVESTER_APP)

dev-describe-HARVESTER:
	kubectl describe pod --namespace=$(NAMESPACE) -l app=$(HARVESTER_APP)

pgcli:
	pgcli postgresql://postgres:postgres@localhost

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor