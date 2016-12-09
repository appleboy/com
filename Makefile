.PHONY: test fmt vet errcheck lint install

PACKAGES ?= $(shell go list ./... | grep -v /vendor/)

all: build

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
	@which golint > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/golang/lint/golint; \
	fi
	for PKG in $(PACKAGES); do golint -set_exit_status $$PKG || exit 1; done;

test:
	for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.txt $$PKG || exit 1; done;

html:
	go tool cover -html=coverage.txt

install:
	go get -t -d -v ./...

clean:
	go clean -x -i ./...
	rm -rf coverage.txt $(EXECUTABLE) $(DIST) vendor $(DEPLOY_WEBHOOK_IMAGE)
