#!/bin/sh

# This script will build the docker images for the application.

CURDIR=`pwd`
source .env

# Build the UI image.
# Reference: https://vuejs.org/v2/cookbook/dockerize-vuejs-app.html
cd $CURDIR/src/app/ui
docker build -t govueapp-ui:$APP_VERSION .

# Build the API image.
cd $CURDIR/src/app/api
docker build -t govueapp-api:$APP_VERSION .

# Build the ingress image.
cd $CURDIR/src/app/ingress
docker build -t govueapp-ingress:$APP_VERSION .

cd $CURDIR