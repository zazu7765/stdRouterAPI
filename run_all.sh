#!/bin/bash

echo "Running bootstrap script..."
./bootstrap.sh
if [ $? -ne 0 ]; then
    echo "Bootstrap script failed. Exiting..."
    exit 1
fi

echo "Running build script..."
./build.sh
if [ $? -ne 0 ]; then
    echo "Build script failed. Exiting..."
    exit 1
fi

./bin/stdRouterApi