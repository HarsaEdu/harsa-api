# Makefile for Go application

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

# Binary name
BINARY_NAME=harsa-api.exe

# Build targets
all: build

build:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/api

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd/api
	./$(BINARY_NAME)

deps:
	$(GOCMD) mod tidy

.PHONY: all build test clean run deps