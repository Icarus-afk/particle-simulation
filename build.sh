#!/bin/bash

# Set the target OS and architecture
export GOOS=windows
export GOARCH=amd64

# Check if the Go compiler is installed
if ! command -v go &> /dev/null
then
    echo "Go compiler not found. Please install it and try again."
    exit 1
fi

# Build the executable
go build -o particle_simulator.exe 2> build_errors.txt

# Check if the executable was created
if [ ! -f "particle_simulator.exe" ]; then
    echo "Failed to build the executable. Check build_errors.txt for details."
    exit 1
fi

echo "Build complete: particle_simulator.exe"