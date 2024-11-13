FROM golang:1.22-alpine AS builder
RUN apk update && apk add --no-cache git

ARG GITHUB_TOKEN
ENV GITHUB_TOKEN=$GITHUB_TOKEN
RUN git config --global url."https://${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"

WORKDIR /usr/app
COPY go.mod go.sum ./
RUN go env -w GOPRIVATE=github.com/repooooo
RUN go mod download

COPY . .

RUN go build -o ./app ./cmd/app/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /usr/app/app .