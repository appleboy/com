.PHONY: test fmt vet errcheck lint install

PACKAGES ?= $(shell go list ./...)

all: build

install:
	@hash govendor > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go get -u github.com/kardianos/govendor; \
	fi
	govendor sync

fmt:
	go fmt $(PACKAGES)

vet:
	go vet $(PACKAGES)

errcheck:
	@which errcheck > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/kisielk/errcheck; \
	fi
	errcheck $(PACKAGES)

lint:
	@hash revive > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/mgechev/revive; \
	fi
	revive -config config.toml ./... || exit 1

test:
	for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.txt $$PKG || exit 1; done;

html:
	go tool cover -html=coverage.txt

coverage:
	curl -s https://codecov.io/bash > .codecov && \
	chmod +x .codecov && \
	./.codecov -f .cover/coverage.txt

clean:
	go clean -x -i ./...
	find . -name "coverage.txt" -delete
