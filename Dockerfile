FROM golang:1.22-alpine AS builder
RUN apk update && apk add --no-cache git

RUN --mount=type=secret,id=github_token \
    git config --global url."https://$(cat /run/secrets/github_token)@github.com/".insteadOf "https://github.com/"

WORKDIR /usr/app
COPY go.mod go.sum ./
RUN go env -w GOPRIVATE=github.com/repooooo
RUN go mod download

COPY . .

RUN go build -o ./app ./cmd/app/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /usr/app/app .