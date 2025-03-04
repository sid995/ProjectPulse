.PHONY: help build run test clean docker-build docker-up docker-down

BINARY_NAME=api
BUILD_DIR=bin

help: ## Display this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build the API server
	cd api && go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/api

run: ## Run the API server locally
	cd api && go run ./cmd/api

dev: ## Run the API server with hot reloading
	cd api && air -c .air.toml

docker-dev: ## Run the entire stack in development mode with live reloading
	docker compose up

test: ## Run tests
	cd api && go test ./... -v

clean: ## Clean build artifacts
	rm -rf api/$(BUILD_DIR)
	rm -rf api/tmp

docker-build: ## Build Docker images
	docker-compose build

docker-up: ## Start Docker containers in development mode
	docker-compose up -d

docker-prod: ## Start Docker containers in production mode
	docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d

docker-down: ## Stop Docker containers
	docker-compose down

docker-logs: ## View Docker logs
	docker-compose logs -f

lint: ## Run linters
	cd api && go vet ./...
	cd api && go fmt ./...

swagger: ## Generate Swagger documentation
	cd api && swag init -g cmd/api/main.go -o ./docs/swagger

tidy: ## Tidy Go modules
	cd api && go mod tidy

# Default target
.DEFAULT_GOAL := help 