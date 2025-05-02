.ONESHELL:
SHELL = /bin/bash

build:
	@echo "building🛠️"
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin ./cmd/main.go

doc:
	@echo "generating docs 🚀"
	@rm -rf docs/
	@go install github.com/swaggo/swag/cmd/swag@v1.8.12 && swag init -d ./cmd,./internal/server/http/router/v1,./internal/server/http/handler/v1,./internal/service/models

run: build
	@echo "running ✅ "
	trap 'rm -f ./bin' EXIT
	./bin

lint:
	@golangci-lint run

test:
	@go test -v ./... --race -cover