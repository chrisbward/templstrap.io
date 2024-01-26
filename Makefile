.PHONY: build templ
.DEFAULT: help

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=main
VERSION := $(shell git describe --tags)
 
 
build:
	@echo "+ $@"
	$(GOBUILD) -o $(BINARY_NAME) -ldflags "-X main.version=${VERSION}" -v ./cmd/

templ:
	@echo "+ $@"
	@templ generate

generate:
	@echo "+ $@"
	go generate ./...

test:
	@echo "+ $@" 
	@go test -v -tags "$(BUILDTAGS) cgo" $(shell go list ./testsuite/...) 