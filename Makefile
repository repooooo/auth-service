.PHONY: build run test clean docker-build docker-run
include .env

build:
	go build -o ./bin/app ./cmd/app/main.go

run: build
	./bin/app --config=./config/config.yaml

test:
	go test -v ./tests