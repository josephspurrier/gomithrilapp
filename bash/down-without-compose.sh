#!/bin/sh

# This script will shutdown the containers for the application without docker-compose.

docker rm gomithrilapp -f
docker rm mysql57 -f