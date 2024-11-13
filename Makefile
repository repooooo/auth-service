.PHONY: build run test clean

build:
	go build -o ./bin/app ./cmd/app/main.go

run: build
	./bin/app --config=./config/config.yaml

test:
	go test -v ./tests