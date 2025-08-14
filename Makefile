# Makefile for MCPJungle
# Go project with comprehensive linting and development commands

# Variables
BINARY_NAME=mcpjungle
BUILD_DIR=build
MAIN_PATH=./main.go

# Go related variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOVET=$(GOCMD) vet
GOFMT=$(GOCMD) fmt

# Linting tools
GOLANGCI_LINT=golangci-lint
GOLANGCI_LINT_VERSION=v2.4.0

# Docker related
DOCKER_IMAGE=mcpjungle
DOCKER_TAG=latest

# Default target
.DEFAULT_GOAL := help

# Help target
.PHONY: help
help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Build targets
.PHONY: build
build: ## Build the application
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)

.PHONY: build-linux
build-linux: ## Build for Linux
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-linux $(MAIN_PATH)

.PHONY: build-darwin
build-darwin: ## Build for macOS
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin $(MAIN_PATH)

.PHONY: build-windows
build-windows: ## Build for Windows
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-windows.exe $(MAIN_PATH)

.PHONY: build-all
build-all: build-linux build-darwin build-windows ## Build for all platforms

# Clean target
.PHONY: clean
clean: ## Clean build artifacts
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

# Test targets
.PHONY: test
test: ## Run tests
	$(GOTEST) -v ./...

.PHONY: test-coverage
test-coverage: ## Run tests with coverage
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

.PHONY: test-race
test-race: ## Run tests with race detection
	$(GOTEST) -race -v ./...

# Linting targets
.PHONY: lint
# Run linter
lint:
	@echo "Running linter..."
	@$(GOLANGCI_LINT) run ./...

.PHONY: fmt
fmt: ## Format code with go fmt
	@echo "Formatting code with go fmt..."
	$(GOFMT) ./...

.PHONY: fmt-check
fmt-check: ## Check if code is formatted correctly
	@echo "Checking code formatting..."
	@if [ -n "$$($(GOFMT) -l .)" ]; then \
		echo "Code is not formatted correctly. Run 'make fmt' to fix."; \
		exit 1; \
	fi
	@echo "Code is properly formatted."

.PHONY: vet
vet: ## Run go vet
	@echo "Running go vet..."
	$(GOVET) ./...

.PHONY: golangci-lint
golangci-lint: ## Run golangci-lint
	@echo "Running golangci-lint..."
	@if ! command -v $(GOLANGCI_LINT) > /dev/null; then \
		echo "golangci-lint not found. Installing..."; \
		$(MAKE) install-golangci-lint; \
	fi
	$(GOLANGCI_LINT) run

.PHONY: golangci-lint-fix
golangci-lint-fix: ## Run golangci-lint with auto-fix
	@echo "Running golangci-lint with auto-fix..."
	@if ! command -v $(GOLANGCI_LINT) > /dev/null; then \
		echo "golangci-lint not found. Installing..."; \
		$(MAKE) install-golangci-lint; \
	fi
	$(GOLANGCI_LINT) run --fix

# Install tools
.PHONY: install-golangci-lint
install-golangci-lint: ## Install golangci-lint
	@echo "Installing golangci-lint..."
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin $(GOLANGCI_LINT_VERSION)

.PHONY: install-tools
install-tools: ## Install all development tools
	@echo "Installing development tools..."
	$(MAKE) install-golangci-lint

# Dependency management
.PHONY: deps
deps: ## Download dependencies
	$(GOGET) -v -t -d ./...

.PHONY: deps-update
deps-update: ## Update dependencies
	$(GOMOD) tidy
	$(GOMOD) download

.PHONY: deps-vendor
deps-vendor: ## Vendor dependencies
	$(GOMOD) vendor

# Docker targets
.PHONY: docker-build
docker-build: ## Build Docker image
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

.PHONY: docker-run
docker-run: ## Run Docker container
	docker run -p 8080:8080 $(DOCKER_IMAGE):$(DOCKER_TAG)

.PHONY: docker-compose-up
docker-compose-up: ## Start services with docker-compose
	docker-compose up -d

.PHONY: docker-compose-down
docker-compose-down: ## Stop services with docker-compose
	docker-compose down

# Utility targets
.PHONY: run
run: ## Run the application
	$(GOCMD) run $(MAIN_PATH)

.PHONY: install
install: ## Install the application
	$(GOCMD) install

.PHONY: version
version: ## Show version information
	@echo "Go version:"
	$(GOCMD) version
	@echo "Module info:"
	$(GOCMD) list -m

# Cleanup targets
.PHONY: clean-all
clean-all: clean ## Clean everything including vendor and coverage files
	rm -rf vendor/
	rm -f coverage.out coverage.html
	rm -f *.prof
	rm -f *.out