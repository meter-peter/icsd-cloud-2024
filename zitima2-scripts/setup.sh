#!/bin/bash
cd ..
# Build the Docker image for the RESTful service
docker build -t restful-service:latest .

# Create a Docker volume for the database
docker volume create pgdata

echo "Setup complete. RESTful service image built and database volume created."