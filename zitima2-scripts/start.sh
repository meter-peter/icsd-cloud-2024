#!/bin/bash

# Configuration
DB_HOST="postgres"
DB_USER="your_username"
DB_PASSWORD="your_password"
DB_NAME="your_database_name"
DB_PORT="5432"
SERVER_PORT="8080"

# Function to remove container if it exists
remove_container_if_exists() {
    if [ $(docker ps -aq -f name=$1) ]; then
        echo "Removing existing $1 container..."
        docker stop $1 >/dev/null 2>&1
        docker rm $1 >/dev/null 2>&1
    fi
}

# Remove existing network if it exists
if [ $(docker network ls -q -f name=myapp-network) ]; then
    echo "Removing existing myapp-network..."
    docker network rm myapp-network >/dev/null 2>&1
fi

# Create a user-defined network
echo "Creating myapp-network..."
docker network create myapp-network

# Remove existing containers if they exist
remove_container_if_exists postgres
remove_container_if_exists restful-service

# Start the PostgreSQL container
echo "Starting PostgreSQL container..."
docker run -d --name postgres \
    --network myapp-network \
    -e POSTGRES_USER=$DB_USER \
    -e POSTGRES_PASSWORD=$DB_PASSWORD \
    -e POSTGRES_DB=$DB_NAME \
    -v pgdata:/var/lib/postgresql/data \
    postgres:13

# Wait for PostgreSQL to be ready
echo "Waiting for PostgreSQL to be ready..."
until docker exec postgres pg_isready; do
    sleep 2
done

# Start the RESTful service container
echo "Starting RESTful service container..."



docker run -d --name restful-service \
    --network myapp-network \
    -p $SERVER_PORT:8080 \
    -e DB_HOST=$DB_HOST \
    -e DB_USER=$DB_USER \
    -e DB_PASSWORD=$DB_PASSWORD \
    -e DB_NAME=$DB_NAME \
    -e DB_PORT=$DB_PORT \
    -e SERVER_PORT=$SERVER_PORT \
    restful-service:latest

echo "Backend services started. RESTful service is accessible at http://localhost:$SERVER_PORT"