# ==============================================================================
# Define dependencies

GOLANG          := golang:1.24
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
		-f infrastructure/docker/Dockerfile.harvester \
		-t $(HARVESTER_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
		.

# ==============================================================================
# Running from within docker

compose-up:
	docker compose -f "./infrastructure/docker/compose.yaml" up

compose-down:
	docker compose -f "./infrastructure/docker/compose.yaml" down
	docker image rm docker-harvester
	docker image rm docker-admin
	docker volume prune -f

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor