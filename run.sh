#!/bin/bash
export $(cat .env | xargs)
go run cmd/bot/main.go