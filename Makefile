.PHONY: help build run test clean docker-dev docker-prod docker-down docker-down-volumes

BINARY_NAME=api
BUILD_DIR=bin
WEB_DIR=web

help: ## Display this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build the API server
	cd api && go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/api

run: ## Run the API server locally
	cd api && go run ./cmd/api

dev: ## Run the API server with hot reloading
	cd api && air -c .air.toml

test: ## Run tests
	cd api && go test ./... -v

clean: ## Clean build artifacts
	rm -rf api/$(BUILD_DIR)
	rm -rf api/tmp

docker-dev: ## Run the entire stack in development mode with live reloading
	docker compose -f docker-compose.yml -f docker-compose.dev.yml up

docker-prod: ## Run the entire stack in production mode
	docker compose -f docker-compose.yml -f docker-compose.prod.yml up -d

docker-dev-web: ## Run the web app in development mode with live reloading
	docker compose -f docker-compose.yml -f docker-compose.dev.yml up web

docker-prod-web: ## Run the web app in production mode
	docker compose -f docker-compose.yml -f docker-compose.prod.yml up web

docker-dev-api: ## Run the backend in development mode with live reloading
	docker compose -f docker-compose.yml -f docker-compose.dev.yml up api

docker-prod-api: ## Run the backend in production mode
	docker compose -f docker-compose.yml -f docker-compose.prod.yml up api

docker-dev-db: ## Run the database in development mode
	docker compose -f docker-compose.yml -f docker-compose.dev.yml up db

docker-prod-db: ## Run the database in production mode
	docker compose -f docker-compose.yml -f docker-compose.prod.yml up db

docker-dev-redis: ## Run the redis in development mode
	docker compose -f docker-compose.yml -f docker-compose.dev.yml up redis

docker-prod-redis: ## Run the redis in production mode
	docker compose -f docker-compose.yml -f docker-compose.prod.yml up redis

docker-build-dev: ## Build the Docker images for development
	docker compose -f docker-compose.yml -f docker-compose.dev.yml build

docker-build-prod: ## Build the Docker images for production
	docker compose -f docker-compose.yml -f docker-compose.prod.yml build

docker-down: ## Stop Docker containers
	docker compose down

docker-down-volumes: ## Stop Docker containers and remove volumes
	docker compose down -v

docker-logs: ## View Docker logs
	docker compose logs -f

lint: ## Run linters
	cd api && go vet ./...
	cd api && go fmt ./...

swagger: ## Generate Swagger documentation
	cd api && swag init -g cmd/api/main.go -o ./docs/swagger

tidy: ## Tidy Go modules
	cd api && go mod tidy

web-dev: ## Run the web app in development mode
	cd $(WEB_DIR) && pnpm dev

web-build: ## Build the web app
	cd $(WEB_DIR) && pnpm build

web-start: ## Start the web app
	cd $(WEB_DIR) && pnpm start

# Default target
.DEFAULT_GOAL := help 