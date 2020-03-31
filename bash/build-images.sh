#!/bin/sh

# This script will build the docker images for the application.

CURDIR=`pwd`
source .env

# Build the API image.
cd $CURDIR/src/app
docker build -t gomithrilapp:$APP_VERSION .

cd $CURDIR