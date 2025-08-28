#!/bin/bash
read -p "press any key to start"

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o build/mc-dashboard-linux ./src
echo "Built mc-dashboard-linux"

read -p "press any key to exit"