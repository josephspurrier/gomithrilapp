#!/bin/sh

# This file is called by Cypress in: src/app/ui/test/support/commands.js

# If TRAVIS is not set, then reset the DB using local tools because TRAVIS
# is always clean.
if [ -z "$TRAVIS" ]; then
    docker exec gomithrilapp_db_1 sh -c "exec mysql -h 127.0.0.1 -u root -p${MYSQL_ROOT_PASSWORD} -e 'DROP DATABASE IF EXISTS main;'"
    docker exec gomithrilapp_db_1 sh -c "exec mysql -h 127.0.0.1 -u root -p${MYSQL_ROOT_PASSWORD} -e 'CREATE DATABASE IF NOT EXISTS main DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;'"
    go run $GOPATH/src/app/api/cmd/dbmigrate/main.go
fi