#!/bin/bash
read -p "press any key to start"

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o build/dashboard-executable ./src
echo "Built dashboard-executable"

read -p "press any key to exit"