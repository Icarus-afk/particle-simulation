#!/bin/bash

# Set the target OS and architecture
export GOOS=windows
export GOARCH=amd64

# Build the executable
go build -o particle_simulator.exe 2> build_errors.txt

echo "Build complete: particle_simulator.exe"