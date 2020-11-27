# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
BUILD_DIR=./bin
BINARY_NAME=go_merge
BINARY_LINUX=$(BINARY_NAME)
BINARY_DARWIN=$(BINARY_NAME)_darwin
BINARY_WINDOWS=$(BINARY_NAME)_windows
MODULE=github.com/alexpfx/golang/go_merge


all: test build
install:
	$(GOINSTALL) -v -ldflags $(LDFLAGS) $(MODULE)/cmd/$(BINARY_NAME)
build:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -v $(MODULE)/cmd/$(BINARY_NAME)

clean:
		$(GOCLEAN)
		rm -f $(BUILD_DIR)/$(BINARY_NAME)*

# Cross compilation
build-linux-amd64:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_LINUX)_amd64_$(GIT_TAG) -v $(MODULE)/cmd/$(BINARY_NAME)

build-all: build-linux-amd64 build-linux-arm build-darwin-amd64 build-windows-amd64