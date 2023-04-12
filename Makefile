.PHONY: build

build:
	go build -v ./cmd/apiserver/main.go

.PHONY: test

test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build

.PHONY: run

run:
	go run -v ./cmd/apiserver/main.go
