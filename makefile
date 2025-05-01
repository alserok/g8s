.ONESHELL:
SHELL = /bin/bash

build:
	@echo "building🛠️"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin ./cmd/main.go

run: build
	@echo "generating docs 🚀"
	@go install github.com/swaggo/swag/cmd/swag@latest && swag init -d ./cmd,./internal/server/http/router/v1,./internal/server/http/handler/v1,./internal/service/models
	@echo "running ✅ "
	trap 'rm -f ./bin' EXIT
	./bin

lint:
	@golangci-lint run

test:
	@go test -v ./... --race -cover