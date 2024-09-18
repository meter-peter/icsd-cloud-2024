#!/bin/bash
cd ..
docker build -t restful-service:latest .
docker volume create pgdata

echo "Setup complete. RESTful service image built and database volume created."