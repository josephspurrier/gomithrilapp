# govueapp

[![Go Report Card](https://goreportcard.com/badge/github.com/josephspurrier/govueapp)](https://goreportcard.com/report/github.com/josephspurrier/govueapp)
[![Build Status](https://travis-ci.org/josephspurrier/govueapp.svg)](https://travis-ci.org/josephspurrier/govueapp)
[![Coverage Status](https://coveralls.io/repos/github/josephspurrier/govueapp/badge.svg?branch=master&timestamp=20190531-01)](https://coveralls.io/github/josephspurrier/govueapp?branch=master)

[![Swagger Validator](http://online.swagger.io/validator?url=https://raw.githubusercontent.com/josephspurrier/govueapp/master/src/app/api/static/swagger/swagger.json)](http://petstore.swagger.io/?url=https://raw.githubusercontent.com/josephspurrier/govueapp/master/src/app/api/static/swagger/swagger.json)

This is an application that uses Vue on the frontend (UI) and Go on the backend (API). This project uses a [Makefile](Makefile) to help centralize frequently used commands.

## Quickstart

To run the application locally, you can run these commands. You don't need any of the the dev tools (Go/npm) installed, you only need Docker (and Docker Compose).

```bash
# Build the docker containers.
make docker-build

# Launch the docker containers.
docker-compose up

# Open your browser for the UI: localhost
# Open your browser for the API: localhost:8081

# Bring down the docker containers.
docker-compose down
```

## Environment Preparation

You should use Go 1.10 or newer. All of the commands below assume you have your GOPATH set to the root of this project directory. This does prevent you from being able to use this command to download and run the project like a typical Go application: `go get github.com/josephspurrier/govueapp`, but we found (after many projects) it is now much easier for you to clone this repo and make changes without having to rewrite imports. This project also has a separate front-end and back-end so it makes sense `go get` would not work anyway.

You can also use [direnv](https://direnv.net/) which sets your environment variables based on your current directory. For instance, you can install direnv, create a file in the root of this project called `.envrc`, and paste in the following:

```bash
# Set $GOPATH for Go.
export GOPATH=`pwd`
# Add the bin directory to $PATH.
export PATH=$PATH:`pwd`/bin
# Add the npm bin directory to $PATH to allow running NPM apps.
export PATH=$PATH:$(npm bin)
```

Save the file and type `direnv allow`. That will automatically set environment variables when you `CD` into the project root and child folders.

## Current Versions

You should use Go 1.10 or newer.

These are the current versions on the front-end components:

- Nuxt.js [v2.6.3](https://github.com/nuxt/nuxt.js/releases/tag/v2.6.3) - you should use the [2.5.X documentation](https://nuxtjs.org/guide/release-notes)
- Vue.js [v2.6.10](https://github.com/vuejs/vue/releases/tag/v2.6.10)
- Vuex [v3.1.0](https://github.com/vuejs/vuex/releases/tag/v3.1.0)
- Bulma [v1.2.3](https://www.npmjs.com/package/@nuxtjs/bulma/v/1.2.3) - you should use the [0.7.4 documentation](https://bulma.io/documentation/)

### Update nuxt

```bash
# Upgrade nuxt to the 1.0 version.
make nuxt-upgrade

# Check the version of nuxt.
make nuxt-upgrade
```

## Getting Started with Development

You can run these commands from different terminals to start the services.

```bash
# Create and run the database container.
make db-init

# Start the UI in local dev mode.
make dev-ui

# Start the API in local dev mode.
make dev-api
```

These are other database commands you can use:

```bash
# Start the DB container.
make db-start

# Stop the DB container.
make db-stop

# Drop the database, create a new database, and apply new migrations.
make db-reset

# Delete the DB container.
make db-rm
```

## Database Migrations

MySQL migrations are performed at boot by [Rove](https://github.com/josephspurrier/rove.), a tool very similiar to Liquibase.

## Go Dependency Management

This projects does not use Go modules - it uses [gvt](https://github.com/FiloSottile/gvt) to vendor dependencies to Go. This decision was made because Visual Studio Code support is still lacking and that just happens to be our preferred IDE: ["⚠️ These tools do not provide a good support for Go modules yet."](https://github.com/Microsoft/vscode-go/wiki/Go-modules-support-in-Visual-Studio-Code). We've used gvt on large teams for years so even though it's deprecated, it still works extremely well for our purposes.

You only need to download gvt if you want to add or update packages. All the vendored files are included to make this project easy to compile.

```bash
# Download gvt.
go get github.com/FiloSottile/gvt

# You should now add the {PROJECTROOT}/bin folder to your $PATH to make gvt available from your terminal.

# Make sure you CD to the api folder before using gvt:
cd $GOPATH/src/app/api
gvt fetch github.com/user/project
```

## Swagger

This projects uses [Swagger v2](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md) to document the API. The entire Swagger spec is generated from comments (annotations) in and by analyzing structs and variables.

```bash
# Download the Swagger generation tool.
make swagger-get

# Generate the swagger spec.
make swagger-gen
```

## Debug UI Tests

Jest is used to run tests on Vue. You can use the steps below to debug the tests.

```bash
# Open chrome and go to this URL.
chrome://inspect

# Click: Open dedicated DevTools for Node

# Add this text to any test:
debugger

# Run this command:
npm run test:debug
```

You can also reference this article to set up a debugger in your IDE:
https://jestjs.io/docs/en/troubleshooting

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
