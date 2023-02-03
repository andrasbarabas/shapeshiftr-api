include .env

.PHONY: all
all: build

# Build project
.PHONY: build
build:
	@go build -o ${PROJECT_NAME} -v ./cmd/main.go

.PHONY: help
help:
	@echo "make: build project"
	@echo "make build: build project"
	@echo "make lint: run golangci-lint (requires Docker)"
	@echo "make lint-dockerfile: run Hadolint linter on Dockerfile (requires Docker)"
	@echo "make run: run project in development mode"
	@echo "make test: run tests"
	@echo "make update: update dependencies"

# Run golangci-lint
.PHONY: lint
lint:
	@DOCKER_COMPOSE_TARGET=lint docker compose up --build api && DOCKER_COMPOSE_TARGET= docker compose rm -fsv

.PHONY: lint-dockerfile
lint-dockerfile:
	@docker run --rm -i hadolint/hadolint < Dockerfile

# Run project in development mode
.PHONY: run
run:
	@go run ./cmd/main.go

# Run tests
.PHONY: test
test:
	@go test -cover ./...

# Update dependencies
.PHONY: update
update:
	@go clean -modcache
	@go get -u all
	@go mod tidy
	@go mod vendor
