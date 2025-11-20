BINARY_API := pvz-api

.PHONY: all run test cover tidy

all: test

run:
	go run ./cmd/api

test:
	go test ./... -count=1

cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

tidy:
	go mod tidy
