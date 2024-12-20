.PHONY: all clean build-linux build-darwin

BINARY_NAME=tappd-cli
VERSION=0.1.0
BUILD_TIME=$(shell date +%FT%T%z)
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME}"

# Default target
all: build-linux build-darwin

# Clean build artifacts
clean:
	rm -rf build/
	rm -f ${BINARY_NAME}-linux-amd64
	rm -f ${BINARY_NAME}-darwin-amd64
	rm -f ${BINARY_NAME}-darwin-arm64

# Build for Linux (amd64)
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o build/${BINARY_NAME}-linux-amd64

# Build for macOS (both Intel and Apple Silicon)
build-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o build/${BINARY_NAME}-darwin-amd64
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o build/${BINARY_NAME}-darwin-arm64

# Build all platforms
build: build-linux build-darwin

# Install locally (useful for development)
install:
	go install ${LDFLAGS}

# Run tests
test:
	go test -v ./... 