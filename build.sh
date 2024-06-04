#!/bin/bash

echo "Building stdRouterApi example..."
if [ ! -d "bin" ]; then
    echo "bin directory does not exist. Creating bin directory..."
    mkdir bin
else
    echo "bin directory already exists."
fi

go build -o bin/stdRouterApi ./src/cmd/main.go

echo "Build completed."
