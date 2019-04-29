#!/bin/sh

# This script will build the containers for the application.
CURDIR=`pwd`
source .env

# Build the UI container.
# Reference: https://vuejs.org/v2/cookbook/dockerize-vuejs-app.html
cd $CURDIR/ui
docker build -t govueapp-ui:$APP_VERSION .

# Build the API container.
cd $CURDIR/api
docker build -t govueapp-api:$APP_VERSION .

cd $CURDIR