.DEFAULT_GOAL := build

.PHONY: test fmt vet build 

test: 
	go test ./...

fmt:  test
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o pokedex
