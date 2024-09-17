#!/bin/bash

# Stop and remove the RESTful service container
docker stop restful-service
docker rm restful-service

# Stop and remove the PostgreSQL container
docker stop postgres
docker rm postgres

# Remove the user-defined network
docker network rm myapp-network

# Note: We're not removing the volume to persist data
# If you want to remove the volume, uncomment the following line:
# docker volume rm pgdata

echo "Backend services stopped and containers removed."