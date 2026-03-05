# Makefile for building agent and API server binaries

# Variables
BINARY_AGENT = agent
BINARY_API_SERVER = api-server

# Targets

.PHONY: all build-agent build-api-server test clean

all: build-agent build-api-server test

build-agent:
	@echo "Building agent..."
	go build -o $(BINARY_AGENT) ./cmd/agent

build-api-server:
	@echo "Building API server..."
	go build -o $(BINARY_API_SERVER) ./cmd/api-server

# Run tests

test:
	@echo "Running tests..."
	go test ./...

# Clean build artifacts

clean:
	@echo "Cleaning build artifacts..."
	rm -f $(BINARY_AGENT) $(BINARY_API_SERVER)
