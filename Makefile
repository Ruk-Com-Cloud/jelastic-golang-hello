.PHONY: help build run clean test tidy fmt vet docker-build docker-run docker-stop docker-clean docker-logs

# Default goal
.DEFAULT_GOAL := help

# Application
APP_NAME := jelastic-golang-hello
BUILD_DIR := ./build
DOCKER_IMAGE := $(APP_NAME):latest

help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build the application
	@echo "Building application..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) .
	@echo "Application built successfully: $(BUILD_DIR)/$(APP_NAME)"

run: ## Run the application
	@echo "Starting HTTP-only application..."
	@go run main.go

test: ## Run tests
	@echo "Running tests..."
	@go test ./...

clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@go clean

# Docker commands
docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE) .
	@echo "Docker image built: $(DOCKER_IMAGE)"

docker-run: ## Run application with Docker Compose
	@echo "Starting application with Docker Compose..."
	@docker-compose up -d
	@echo "Application is running at http://localhost:3000"

docker-stop: ## Stop Docker containers
	@echo "Stopping Docker containers..."
	@docker-compose down

docker-logs: ## View Docker logs
	@docker-compose logs -f

docker-clean: ## Clean Docker resources
	@echo "Cleaning Docker resources..."
	@docker-compose down --rmi all --volumes --remove-orphans
	@docker system prune -f

docker-shell: ## Get shell access to running container
	@docker exec -it jelastic-golang-hello /bin/sh

# Utility commands
tidy: ## Tidy go modules
	@echo "Tidying go modules..."
	@go mod tidy

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...

vet: ## Vet code
	@echo "Vetting code..."
	@go vet ./...