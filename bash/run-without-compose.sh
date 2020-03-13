#!/bin/sh

# This script will build and run the containers for the application without docker-compose.

# CD to the UI folder.
cd $GOPATH/src/app/ui

# Build the docker container.
# Reference: https://vuejs.org/v2/cookbook/dockerize-vuejs-app.html
docker build -t govueapp-ui:1.0 .

# Run the docker container.
docker run -it -p 80:80 --rm --name govueapp-ui govueapp-ui:1.0

# CD to the API folder.
cd $GOPATH/src/app/api

# Build the docker container.
docker build -t govueapp-api:1.0 .

# Run the API docker container.
docker run -it -p 8081:8081 --rm --name govueapp-api govueapp-api:1.0

# Launch a MySQL database.
docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d --rm --name mysql57 mysql:5.7
docker exec mysql56 sh -c 'exec mysql -uroot -ppassword -e "CREATE DATABASE IF NOT EXISTS main DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;"'