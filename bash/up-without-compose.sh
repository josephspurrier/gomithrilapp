#!/bin/sh

# This script will build and run the containers for the application without docker-compose.

# Launch a MySQL database.
docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password --rm --name mysql57 mysql:5.7

# CD to the API folder.
cd $GOPATH/src/app

# Build the docker container.
docker build -t gomithrilapp:1.0 .

# Run the app docker container.
docker run -d -p 8081:8081 -e MYSQL_HOST=host.docker.internal -e MYSQL_ROOT_PASSWORD=password --rm --name gomithrilapp-api gomithrilapp-api:1.0

# Create the database.
docker exec mysql57 sh -c 'exec mysql -u root -ppassword -e "CREATE DATABASE IF NOT EXISTS main DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;"'