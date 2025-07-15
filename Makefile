.PHONY: test build

VERSION = $(shell grep "const Version =" cmd/sub/version.go | grep "const Version =" | sed -e 's-.*= "--' -e 's-".*--')
CURRENT_DIR = $(shell pwd)

build:
	go build -o build/pvault -ldflags "-s -w" cmd/main.go

build-docker:
	docker build -f Dockerfile.release -t ghcr.io/okieoth/pvault:$(VERSION) .

test:
	go test ./... && echo ":)" || echo ":-/"
