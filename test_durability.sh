#!/bin/bash

# Start Docker Compose and wait for services to be healthy
docker-compose up -d --build
echo "Waiting for services to be healthy..."
sleep 10  # Wait for containers to initialize

# Run the initial client test to insert data and verify
docker-compose exec client ./kv739_test
echo "Initial data insert and verification complete."

# Simulate server crash
echo "Simulating server crash..."
docker-compose stop server
echo "Server crashed. Waiting for 5 seconds..."
sleep 5  # Wait for a few seconds to simulate downtime

# Restart the server to simulate recovery
echo "Restarting the server..."
docker-compose start server

# Wait for the server to come back online
echo "Waiting for the server to recover..."
sleep 10

# Restart the client container to ensure it is connected to the server
echo "Restarting the client container..."
docker-compose start client

# Run the recovery test to verify data after server restart
echo "Running recovery test..."
docker-compose exec client ./kv739_test --test-recovery
echo "Recovery test complete."

# Clean up
echo "Stopping all services..."
docker-compose down