APP_NAME := ocr
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
export GOWORK := off

.PHONY: build run test race lint cover clean

build:
	@mkdir -p bin
	@BUILD_VERSION="$(VERSION)"; \
	go build -trimpath -ldflags "-s -w -X main.version=$${BUILD_VERSION}" -o bin/$(APP_NAME) ./cmd/ocr; \
	ACTUAL_VERSION="$$(./bin/$(APP_NAME) --version | tr -d '\r\n')"; \
	if [ "$${ACTUAL_VERSION}" != "$${BUILD_VERSION}" ]; then \
		echo "$(APP_NAME) --version mismatch: expected '$${BUILD_VERSION}', got '$${ACTUAL_VERSION}'"; \
		exit 1; \
	fi

run:
	go run ./cmd/ocr --help

test:
	go test ./...

race:
	go test -race ./...

lint:
	golangci-lint run

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

clean:
	rm -rf bin coverage.out
