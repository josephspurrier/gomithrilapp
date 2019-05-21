# govueapp

[![Go Report Card](https://goreportcard.com/badge/github.com/josephspurrier/govueapp)](https://goreportcard.com/report/github.com/josephspurrier/govueapp)
[![Build Status](https://travis-ci.org/josephspurrier/govueapp.svg)](https://travis-ci.org/josephspurrier/govueapp)
[![Coverage Status](https://coveralls.io/repos/github/josephspurrier/govueapp/badge.svg?branch=master&timestamp=20190521-01)](https://coveralls.io/github/josephspurrier/govueapp?branch=master)

[![Swagger Validator](http://online.swagger.io/validator?url=https://raw.githubusercontent.com/josephspurrier/govueapp/master/src/app/api/swagger.json)](http://petstore.swagger.io/?url=https://raw.githubusercontent.com/josephspurrier/govueapp/master/src/app/api/swagger.json)

This is an application that uses Vue on the frontend (UI) and Go on the backend (API).

## Current Versions

You should use Go 1.9 or newer.

These are the current versions on the front-end components:

- Nuxt.js [v2.6.3](https://github.com/nuxt/nuxt.js/releases/tag/v2.6.3) - you should use the [2.5.X documentation](https://nuxtjs.org/guide/release-notes)
- Vue.js [v2.6.10](https://github.com/vuejs/vue/releases/tag/v2.6.10)
- Vuex [v3.1.0](https://github.com/vuejs/vuex/releases/tag/v3.1.0)
- Bulma [v1.2.3](https://www.npmjs.com/package/@nuxtjs/bulma/v/1.2.3) - you should use the [0.7.4 documentation](https://bulma.io/documentation/)

```bash
# Upgrade nuxt to the 1.0 version.
npm upgrade nuxt

# Run NPM apps from terminal.
export PATH=$PATH:$(npm bin)

# Check the version of nuxt.
nuxt --version
```

## Run Application Locally

To run the application locally, you can run these commands. You don't need any of the the dev tools installed, you only need Docker.

```bash
# Build the docker containers.
bash build-containers.sh

# Launch the docker containers.
docker-compose up

# Open your browser for the UI: localhost
# Open your browser for the API: localhost:8081

# Bring down the docker containers.
docker-compose down
```

If you want to run any of the containers manually, you can build and run them using these commands.

```bash
# Set the GOPATH to the current directory.
export GOPATH=`pwd`

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
docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -d --rm --name mysql56 mysql:5.6
docker exec mysql56 sh -c 'exec mysql -uroot -ppassword -e "CREATE DATABASE IF NOT EXISTS main DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;"'
```

## Getting Started with Development

```bash
# Set the GOPATH to the current directory.
export GOPATH=`pwd`

# Start the UI.
cd $GOPATH/src/app/ui
npm install
npm run dev

# Start the API.
cd $GOPATH/src/app/api
go run main.go
```

## Database Migrations

Migrations are perform at boot by Rove: https://github.com/josephspurrier/rove.

## Go Dependency Management

I was going to use gvt, but decided to use `go mod` instead.

```bash
# Example of how to vendor a dependency with the experimental module support in Go 1.11.X
GO111MODULE=on go get github.com/josephspurrier/rove
# Example of how to vendor all missing dependencies
GO111MODULE=on go mod vendor
```

This is how I vendored the first dependencies.

```bash
# Reference: https://github.com/FiloSottile/gvt
# Reference: https://github.com/golang/go/wiki/Modules
GO111MODULE=on go mod init
GO111MODULE=on go mod vendor
```

## Swagger

This projects uses [Swagger v2](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md) to document the API. The entire Swagger spec is generated from the code in this repository.

### Install Swagger

This tool will generate the Swagger spec from annotations in the Go code. It will read the comments in the code and will pull types from structs.

```bash
go get github.com/go-swagger/go-swagger/cmd/swagger
```

### Generate Swagger Spec

```bash
# CD to the webapi folder.
cd src/app/api

# Generate the swagger spec.
swagger generate spec -o ./swagger.json

# Replace 'example' with 'x-example' in the swagger spec.
## MacOS
sed -i '' -e 's/example/x\-example/' ./swagger.json
## Linux
sed -i'' -e 's/example/x\-example/' ./swagger.json

# Validate the swagger spec.
swagger validate ./swagger.json

# Serve the spec for the browser.
swagger serve -F=swagger ./swagger.json
```

## References

These are notes on the project.

```bash
# Reference: https://scotch.io/tutorials/implementing-authentication-in-nuxtjs-app
# I didn't use Auth, but I did use the examples of the getters.

# Reference: https://nuxtjs.org/examples/auth-external-jwt/
npm install cookieparser --save
npm install js-cookie --save

# Reference: https://nuxtjs.org/api/context/
# The context provides additional objects that are available to Vue components
# like the middleware.

# Reference: https://vuex.vuejs.org/guide/getters.html
# Getters for the vuex state.
```
