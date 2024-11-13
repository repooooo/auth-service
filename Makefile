.PHONY: build run test clean docker-build docker-run
include .env

build:
	go build -o ./bin/app ./cmd/app/main.go

run: build
	./bin/app --config=./config/config.yaml

test:
	go test -v ./tests

docker-build:
	docker build --build-arg GITHUB_TOKEN=${GITHUB_TOKEN} -f Dockerfile -t auth-service:latest .
docker-run:
	docker run -v $(shell pwd)/config/config.yaml:/root/config.yaml auth-service:latest ./app --config=/root/config.yaml