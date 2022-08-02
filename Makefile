.DEFAULT_GOAL:=help
SHELL:=/bin/bash

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\033[36m***This is the Makefile to build and start myapp locally***\033[0m \n \nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

kill:
	docker-compose down

rebuild:
	docker-compose build --no-cache

local:
	docker-compose up --build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build cmd/main.go
.PHONY:build

