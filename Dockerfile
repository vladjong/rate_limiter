FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./.env ./.env

RUN go build -o rate_limiter ./cmd/rate_limiter/main.go
CMD ["/app/rate_limiter"]