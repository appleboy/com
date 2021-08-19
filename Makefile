GO ?= go
GOFILES := $(shell find . -name "*.go" -type f)
GOFMT ?= gofumpt -l -s

all: test

.PHONY: test
test:
	@$(GO) test -v -cover -coverprofile coverage.out ./... && echo "\n==>\033[32m Ok\033[m\n" || exit 1

.PHONY: benchmark
benchmark:
	@$(GO) test -v -benchmem -run=^$ -count=5 -bench=^Benchmark -benchtime=3s ./array/
	@$(GO) test -v -benchmem -run=^$ -count=5 -bench=^Benchmark -benchtime=3s ./convert/
	@$(GO) test -v -benchmem -run=^$ -count=5 -bench=^Benchmark -benchtime=3s ./random/
