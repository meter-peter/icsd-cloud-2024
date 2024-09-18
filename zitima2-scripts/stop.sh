#!/bin/bash
docker stop restful-service
docker rm restful-service
docker stop postgres
docker rm postgres
docker network rm myapp-network
echo "Backend services stopped and containers removed."