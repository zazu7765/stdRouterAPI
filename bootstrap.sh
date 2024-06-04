#!/bin/bash

# Function to check if a command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

if ! command_exists sqlc; then
    echo "sqlc is not installed. Installing sqlc..."
    go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
    else
    echo "sqlc is already installed, proceeding to next step..."
fi

echo "Generating schema code..."
sqlc generate -f src/configs/sqlc.yaml

# 2. Go get all Go dependencies
echo "Fetching Go dependencies..."
go mod tidy
echo 

echo "Bootstrap completed."