.PHONY: test build

VERSION = $(shell grep "const Version =" cmd/sub/version.go | grep "const Version =" | sed -e 's-.*= "--' -e 's-".*--')
CURRENT_DIR = $(shell pwd)

build:
	go build -o build/pvault -ldflags "-s -w" cmd/main.go

build-docker:
	docker build -f Dockerfile.release -t ghcr.io/okieoth/pvault:$(VERSION) .
	docker tag ghcr.io/okieoth/pvault:$(VERSION) ghcr.io/okieoth/pvault:latest

build-publish:
	docker publish ghcr.io/okieoth/pvault:$(VERSION)
	docker publish ghcr.io/okieoth/pvault

test:
	go test ./... && echo ":)" || echo ":-/"
