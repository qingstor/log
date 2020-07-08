SHELL := /bin/bash

.PHONY: check format test

help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  check               to format, vet and lint"
	@echo "  build               to create bin directory and build"
	@echo "  test                to execute unit tests"

check: format vet lint

format:
	@go fmt ./...

vet:
	@go vet ./...

lint:
	# staticcheck: go get honnef.co/go/tools/...
	@staticcheck ./...

build: tidy check
	@go build ./...

test:
	@go test -race -v ./...

bench:
	@go list ./... | xargs -n1 go test -bench=. -run="^$$" -benchmem

tidy:
	@go mod tidy && go mod verify
