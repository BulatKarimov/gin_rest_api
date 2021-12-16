# .PHONY: build
# 
run:
	go run ./cmd/api_server/main.go
build:
	go build -v ./cmd/api_server

.DEFAULT_GOAL := run