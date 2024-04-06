#!/bin/bash

# Capture the current directory
current_dir=$(pwd)

# Check if ping-ack-api directory exists
if [ -d "ping-ack-api" ]; then
    echo "ping-ack-api directory found"
else
    echo "Error: ping-ack-api directory not found"
    exit 1
fi

# Change directory to ping-ack-api
cd ping-ack-api || { echo "Error: Failed to change directory to ping-ack-api"; exit 1; }

# Run swag init
swag init || { echo "Error: Failed to run 'swag init'"; exit 1; }

# Change back to the previous directory
cd "$current_dir" || { echo "Error: Failed to change directory back to $current_dir"; exit 1; }

# Stop and remove existing Docker containers (if any), ignoring errors
docker-compose down || true

# Start Docker containers, building if necessary
docker-compose up --build

