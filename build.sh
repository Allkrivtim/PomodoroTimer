#!/bin/bash
set -e

GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/main ./cmd/bot/main.go