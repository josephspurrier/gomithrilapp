---
id: env-prep
title: Environment Prep
---

Once you have cloned the repo, you will need the following tools for local development.

### Go

You should use Go 1.11 or newer. We recommend [gvm](https://github.com/moovweb/gvm) for installing and managing your versions of Go.

All of the commands below assume you have your GOPATH set to the root of this project directory. This is by design because we found (after many projects) it is much easier for you to clone this repo and make changes without having to rewrite imports if they are all contained within the project.

### Node and npm

You should install [NodeJS and npm](https://nodejs.org/).

These are the current versions on the front-end components:

- Nuxt.js [v2.11.0](https://github.com/nuxt/nuxt.js/releases/tag/v2.11.0) - [docs](https://nuxtjs.org/guide/release-notes)
- Vue.js [v2.6.11](https://github.com/vuejs/vue/releases/tag/v2.6.11) - [docs](https://vuejs.org/v2/guide/)
- Vuex [v3.1.2](https://github.com/vuejs/vuex/releases/tag/v3.1.2) - [docs](https://vuex.vuejs.org/)
- Bulma [v1.2.7](https://www.npmjs.com/package/@nuxtjs/bulma/v/1.2.7) - [docs](https://bulma.io/documentation/)
- Swagger UI [v3.25.0](https://github.com/swagger-api/swagger-ui/releases/tag/v3.25.0) - [docs](https://swagger.io/tools/swagger-ui/)

### Visual Studio Code (VSCode) Setup

You can use any IDE, but here is what you need for VSCode. It was quite a challenge getting ESLint to work properly when the .eslintrc.js file is not in the root of the project - the trick was the "eslint.workingDirectories" setting. All the settings are included in the .vscode/settings.json file. I use VSCode open only to the root of the project with no other projects. I recommend the following VSCode extensions:

- [ESLint](https://marketplace.visualstudio.com/items?itemName=dbaeumer.vscode-eslint)
- [Go](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)
- [Vetur](https://marketplace.visualstudio.com/items?itemName=octref.vetur)

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

# Start the documentation service in local dev mode after install the dependencies.
make doc-dep
make doc-dev
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

Run tests for UI:

```bash
# Run the UI tests.
make ui-test
```

Run tests for the API:

```bash
# Create and run the database container.
make db-init

# Run the API tests.
make api-test
```