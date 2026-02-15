# Set GO command, defaults to 'go' but can be overridden
GO ?= go
# Find all Go source files in the project
GOFILES := $(shell find . -type f -name "*.go")

## test: run tests
test:
	# Run tests with verbose output and coverage, generate coverage.txt report
	# Print green "Ok" message on success, exit with error code 1 on failure
	@$(GO) test -v -cover -coverprofile coverage.txt ./... && echo "\n==>\033[32m Ok\033[m\n" || exit 1

## fmt: format go files using golangci-lint
fmt:
	# Check if golangci-lint is installed, if not, install version 2.9.0
	@command -v golangci-lint >/dev/null 2>&1 || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $$($(GO) env GOPATH)/bin v2.9.0
	# Format all Go files using golangci-lint
	golangci-lint fmt

## lint: run golangci-lint to check for issues
lint:
	# Check if golangci-lint is installed, if not, install version 2.9.0
	@command -v golangci-lint >/dev/null 2>&1 || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $$($(GO) env GOPATH)/bin v2.9.0
	# Run linter to check for code quality issues and potential bugs
	golangci-lint run

## clean: remove build artifacts and test coverage
clean:
	# Remove generated coverage report file
	rm -rf coverage.txt

## benchmark: run performance benchmarks for all packages
benchmark:
  # Run benchmarks for array package: 5 iterations, 3s each, with memory statistics
  @$(GO) test -v -benchmem -run=^$ -count=5 -bench=^Benchmark -benchtime=3s ./array/
  # Run benchmarks for convert package: 5 iterations, 3s each, with memory statistics
  @$(GO) test -v -benchmem -run=^$ -count=5 -bench=^Benchmark -benchtime=3s ./convert/
  # Run benchmarks for random package: 5 iterations, 3s each, with memory statistics
  @$(GO) test -v -benchmem -run=^$ -count=5 -bench=^Benchmark -benchtime=3s ./random/

# Declare targets as phony (not actual files) to avoid conflicts
.PHONY: help test fmt lint clean

## help: print this help message
help:
	@echo 'Usage:'
	# Extract all ## comments and format them as a help menu
	# Parse comments starting with ##, format as columns, and indent output
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
