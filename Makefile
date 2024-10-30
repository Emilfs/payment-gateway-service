
BINARY_NAME=paymentgateway
BUILD_DIR=./build

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/$(BUILD_DIR)
GOFILES=$(wildcard *.go)

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

DOCKER_COMPOSE_CMD=docker-compose

.PHONY: all build test clean deps docker-build docker-up

all: build

build:
	@echo "Building $(BINARY_NAME)..."
	@$(GOBUILD) -o $(GOBIN)/$(BINARY_NAME) -v $(GOFILES)

test:
	@echo "Running tests..."
	@$(GOTEST) -v ./...

clean:
	@echo "Cleaning up..."
	@$(GOCLEAN)
	@rm -f $(GOBIN)/$(BINARY_NAME)

deps:
	@echo "Checking and downloading dependencies..."
	@$(GOGET) -v -d ./...

docker-build:
	@echo "Building Docker image..."
	@$(DOCKER_COMPOSE_CMD) build

docker-up:
	@echo "Starting Docker containers..."
	@$(DOCKER_COMPOSE_CMD) up
