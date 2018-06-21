# Runs on Windows as well using mingw32-make.exe installed with msys64

export GOPATH:=$(abspath $(dir $(lastword $(MAKEFILE_LIST)))/../..)

APP?=app
RELEASE?=1.0.0

COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

BINARY_NAME?=hoppity


all: build test
test:
	go test -v ./...


run:
	./$(BINARY_NAME)



.PHONY: build
build: clean
	go build -o $(BINARY_NAME) \
		-ldflags "-X main.Version=${RELEASE} -X main.GitCommit=${COMMIT} -X main.BuildDate=${BUILD_TIME}"

.PHONY: clean
clean:
	go clean ./...


.PHONY: vendor
vendor: prepare_govendor
	govendor sync

HAS_GOVENDOR := $(shell command -v govendor;)
#HAS_METALINTER := $(shell command -v gometalinter;)

.PHONY: prepare_govendor
prepare_govendor:
ifndef HAS_GOVENDOR
	go get -u -v -d github.com/kardianos/govendor && \
	go install -v github.com/kardianos/govendor
endif

