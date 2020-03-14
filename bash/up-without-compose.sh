#!/bin/sh

# This script will build and run the containers for the application without docker-compose.

# Launch a MySQL database.
docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password --rm --name mysql57 mysql:5.7

# CD to the API folder.
cd $GOPATH/src/app/api

# Build the docker container.
docker build -t govueapp-api:1.0 .

# Run the API docker container.
docker run -d -p 8081:8081 -e MYSQL_HOST=host.docker.internal -e MYSQL_ROOT_PASSWORD=password --rm --name govueapp-api govueapp-api:1.0

# CD to the UI folder.
cd $GOPATH/src/app/ui

# Build the docker container.
# Reference: https://vuejs.org/v2/cookbook/dockerize-vuejs-app.html
docker build -t govueapp-ui:1.0 .

# Run the docker container.
docker run -d -p 8080:8080 --rm --name govueapp-ui govueapp-ui:1.0

# Create the database.
docker exec mysql57 sh -c 'exec mysql -u root -ppassword -e "CREATE DATABASE IF NOT EXISTS main DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;"'

# CD to the ingress folder.
cd $GOPATH/src/app/ingress

# Build the docker container.
docker build -t govueapp-ingress:1.0 .

# Run the ingress docker container.
docker run -d -p 80:80 -e UI_URL=http://host.docker.internal:8080 -e API_URL=http://host.docker.internal:8081 --rm --name govueapp-ingress govueapp-ingress:1.0