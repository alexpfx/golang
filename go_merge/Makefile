# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
BUILD_DIR=bin
BINARY_NAME=go_merge
BINARY_LINUX=$(BINARY_NAME)


build:
	$(GOBUILD) -o $(BUILD_DIR) ./...

install:
	$(GOINSTALL) ./...
