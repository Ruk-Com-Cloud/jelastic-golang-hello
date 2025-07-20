.PHONY: help build run seed rollback list-seeders clean test

# Default goal
.DEFAULT_GOAL := help

# Application
APP_NAME := jelastic-golang-hello
BUILD_DIR := ./build

help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build the application
	@echo "Building application..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) .
	@echo "Application built successfully: $(BUILD_DIR)/$(APP_NAME)"

run: ## Run the application
	@echo "Starting application..."
	@go run main.go

test: ## Run tests
	@echo "Running tests..."
	@go test ./...

clean: ## Clean build artifacts
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@go clean

# Database commands
seed: ## Run all database seeders
	@echo "Running database seeders..."
	@go run cmd/seeder/main.go -action=seed

seed-users: ## Run only user seeder
	@echo "Running user seeder..."
	@go run cmd/seeder/main.go -action=seed -seeder=UserSeeder

rollback: ## Rollback all seeders
	@echo "Rolling back database seeders..."
	@go run cmd/seeder/main.go -action=rollback

list-seeders: ## List available seeders
	@echo "Available seeders:"
	@go run cmd/seeder/main.go -action=list

# Docker commands
docker-up: ## Start PostgreSQL with Docker Compose
	@echo "Starting PostgreSQL..."
	@docker-compose up -d

docker-down: ## Stop PostgreSQL
	@echo "Stopping PostgreSQL..."
	@docker-compose down

docker-logs: ## View PostgreSQL logs
	@docker-compose logs -f postgres

docker-reset: ## Reset PostgreSQL (remove volume)
	@echo "Resetting PostgreSQL database..."
	@docker-compose down -v
	@docker-compose up -d

# Development workflow
dev-setup: docker-up ## Setup development environment
	@echo "Waiting for database to be ready..."
	@sleep 5
	@$(MAKE) seed
	@echo "Development environment ready!"

dev-reset: docker-reset ## Reset development environment
	@echo "Waiting for database to be ready..."
	@sleep 5
	@$(MAKE) seed
	@echo "Development environment reset!"

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