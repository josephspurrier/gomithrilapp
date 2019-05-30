#!/bin/bash

# If TRAVIS is not set, load from the environment variables.
if [ -z "$TRAVIS" ]; then
    source $GOPATH/.env
fi  

docker exec mysql56 sh -c "exec mysql -h 127.0.0.1 -uroot -p${MYSQL_ROOT_PASSWORD} -e 'DROP DATABASE IF EXISTS main;'"
docker exec mysql56 sh -c "exec mysql -h 127.0.0.1 -uroot -p${MYSQL_ROOT_PASSWORD} -e 'CREATE DATABASE IF NOT EXISTS main DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;'"

go install -v ../...

$GOPATH/bin/dbmigrate