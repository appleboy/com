GO ?= go
GOFILES := $(shell find . -type f -name "*.go")

test: ## run tests
	@$(GO) test -v -cover -coverprofile coverage.txt ./... && echo "\n==>\033[32m Ok\033[m\n" || exit 1

fmt: ## format go files using golangci-lint
	@command -v golangci-lint >/dev/null 2>&1 || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $$($(GO) env GOPATH)/bin v2.9.0
	golangci-lint fmt

lint: ## run golangci-lint to check for issues
	@command -v golangci-lint >/dev/null 2>&1 || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $$($(GO) env GOPATH)/bin v2.9.0
	golangci-lint run

clean: ## remove build artifacts and test coverage
	rm -rf coverage.txt

.PHONY: help test fmt lint clean

benchmark: ## run performance benchmarks for all packages
  @$(GO) test -v -benchmem -run=^$ -count=5 -bench=^Benchmark -benchtime=3s ./array/
  @$(GO) test -v -benchmem -run=^$ -count=5 -bench=^Benchmark -benchtime=3s ./convert/
  @$(GO) test -v -benchmem -run=^$ -count=5 -bench=^Benchmark -benchtime=3s ./random/

.PHONY: help
help: ## print this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo ""
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
