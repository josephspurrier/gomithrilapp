# govueapp

[![Go Report Card](https://goreportcard.com/badge/github.com/josephspurrier/govueapp)](https://goreportcard.com/report/github.com/josephspurrier/govueapp)
[![Build Status](https://travis-ci.org/josephspurrier/govueapp.svg)](https://travis-ci.org/josephspurrier/govueapp)
[![Coverage Status](https://coveralls.io/repos/github/josephspurrier/govueapp/badge.svg?branch=master&timestamp=20190531-01)](https://coveralls.io/github/josephspurrier/govueapp?branch=master)

[![Swagger Validator](http://online.swagger.io/validator?url=https://raw.githubusercontent.com/josephspurrier/govueapp/master/src/app/api/static/swagger/swagger.json)](http://petstore.swagger.io/?url=https://raw.githubusercontent.com/josephspurrier/govueapp/master/src/app/api/static/swagger/swagger.json)

This is a sample notepad application that uses Vue on the front-end (UI) and Go on the back-end (API). This project is designed to show good development and CI/CD practices as well as integrations between modern development tools. This project uses a [Makefile](Makefile) to centralize frequently used commands.

## Quick Start Guide (QSG)

To run the application locally, you can run these commands. You don't need any of the the dev tools (Go/npm) installed, you only need Docker (and Docker Compose).

```bash
# Clone the repo.
git clone git@github.com:josephspurrier/govueapp.git

# CD to the project directory.
cd govueapp

# Build the docker containers.
make docker-build

# Run the docker containers: DB, API, and UI.
docker-compose up

# Open your browser to the UI: http://localhost
# Open your browser to the API: http://localhost:8081
# Open your MySQL tool to the DB: localhost:3306

# Stop and remove the docker containers.
docker-compose down
```

## Environment Preparation

Once you have cloned the repo, you will need the following tools for local development.

### Go

You should use Go 1.11 or newer. We recommend [gvm](https://github.com/moovweb/gvm) for installing and managing your versions of Go.

All of the commands below assume you have your GOPATH set to the root of this project directory. This does prevent you from being able to use this command to download and run the project like a typical Go application: `go get github.com/josephspurrier/govueapp`, but we found (after many projects) it is now much easier for you to clone this repo and make changes without having to rewrite imports. This project also has a separate front-end and back-end so it makes sense `go get` would not work anyway.

### Node and npm

You should install [NodeJS and npm](https://nodejs.org/).

These are the current versions on the front-end components:

- Nuxt.js [v2.6.3](https://github.com/nuxt/nuxt.js/releases/tag/v2.6.3) - you should use the [2.5.X documentation](https://nuxtjs.org/guide/release-notes)
- Vue.js [v2.6.10](https://github.com/vuejs/vue/releases/tag/v2.6.10)
- Vuex [v3.1.0](https://github.com/vuejs/vuex/releases/tag/v3.1.0)
- Bulma [v1.2.3](https://www.npmjs.com/package/@nuxtjs/bulma/v/1.2.3) - you should use the [0.7.4 documentation](https://bulma.io/documentation/)

Yuu can use these commands to interact with nuxt.

```bash
# CD to the UI folder.
cd ${GOPATH}/src/app/ui

# Check the versions of packages.
npm outdated

# Update all the packages to the latest version (specified by the tag config), respecting semver.
# https://docs.npmjs.com/cli-commands/update.html
npm update

# Scan your project for vulnerabilities and automatically install any compatible updates to vulnerable dependencies.
# https://docs.npmjs.com/cli/audit
npm audit fix

# Get the current verson of a package.
npm -v eslint

# Get a list of all packages that rely on another package
# https://docs.npmjs.com/cli/ls
npm ls typescript

# View all top level packages.
npm ls --depth=0

# Install the latest version of a package.
# https://bytearcher.com/articles/using-npm-update-and-npm-outdated-to-update-dependencies/
npm install eslint@latest

# Install the latest verson package of a major version.
npm install eslint@^5.0.0

# Install the exact version of a package.
npm install lodash@4.17.4

# Remove a package.
# https://docs.npmjs.com/cli/uninstall
npm uninstall eslint-config-standard

# Use --save (-S) to add a package to the package.json dependencies (packages required when the app is built).
# Use --save-dev (-D) to add a package to the package.json devDependencies (packages used during dev to build, bundle, lint).

# Check the version of nuxt.
make nuxt-version

# Upgrade nuxt to the new version.
# When doing an update to nuxt, you should upgrade, remove node_modules dir,
# delete the package-lock.json, and then run 'npm install' again.
make nuxt-upgrade
```

### Environment Variables

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

## Getting Started with Development

You can run these commands from different terminals to start the services.

```bash
# Create and run the database container.
make db-init

# Start the UI in local dev mode after installing dependencies.
make ui-dep
make ui-dev

# Start the API in local dev mode after installing the dependencies.
make api-dep
make api-dev
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

Run tests in Go:

```bash
# Create and run the database container.
make db-init

# Run the API tests.
make api-test
```

## Database Migrations

MySQL migrations are performed at boot by [Rove](https://github.com/josephspurrier/rove), a tool very similiar to Liquibase.

## Go Dependency Management

This projects does not use Go modules - it uses [gvt](https://github.com/FiloSottile/gvt/blob/master/README.old.md) to vendor dependencies to Go. This decision was made because Visual Studio Code support is still lacking and that just happens to be our preferred IDE: ["⚠️ These tools do not provide a good support for Go modules yet."](https://github.com/Microsoft/vscode-go/wiki/Go-modules-support-in-Visual-Studio-Code). We've used gvt on large teams for years so even though it's deprecated, it still works extremely well for our purposes.

```bash
# Download gvt.
make gvt-get

# You can now remove the folder: src/github.com/FiloSottile/gvt
# You should now add the {PROJECTROOT}/bin folder to your $PATH to make gvt available from your terminal.

# Make sure you CD to the api folder before using gvt.
cd $GOPATH/src/app/api

# Here is a sample command to add a new dependency to the project.
gvt fetch github.com/user/project
```

## Swagger

This projects uses [Swagger v2](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md) to document the API. The entire Swagger spec is generated from comments (annotations) in and by analyzing structs and variables.

```bash
# Download the Swagger generation tool.
make swagger-get

# Generate the swagger spec.
make swagger-gen

# Your browser will open to: http://petstore.swagger.io/?url=http://localhost:{RANDOMPORT}/swagger.json

# The output file will be here:
# src/app/api/static/swagger/swagger.json
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
