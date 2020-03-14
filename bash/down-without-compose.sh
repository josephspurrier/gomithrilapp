#!/bin/sh

# This script will shutdown the containers for the application without docker-compose.

docker rm govueapp-ingress -f
docker rm govueapp-ui -f
docker rm govueapp-api -f
docker rm mysql57 -f