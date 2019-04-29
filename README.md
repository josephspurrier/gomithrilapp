# govueapp

This is an application that uses Vue on the frontend (UI) and Go on the backend (API).

## Current Versions

These are the current versions:

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

# Open your browser for the UI: 127.0.0.1
# Open your browser for the API: 127.0.0.1:8081

# Bring down the docker containers.
docker-compose down
```

If you want to run any of the containers manually, you can build and run them using these commands.

```bash
# CD to the UI folder.
cd ui

# Build the docker container.
# Reference: https://vuejs.org/v2/cookbook/dockerize-vuejs-app.html
docker build -t govueapp-ui:1.0 .

# Run the docker container.
docker run -it -p 80:80 --rm --name govueapp-ui govueapp-ui:1.0

# CD to the API folder.
cd api

# Build the docker container.
docker build -t govueapp-api:1.0 .

# Run the API docker container.
docker run -it -p 8081:8081 --rm --name govueapp-api govueapp-api:1.0

# Launch a test database.
docker run -p 3306:3306 -e MYSQL_ROOT_PASSWORD=rootpass --rm --name mysql56 mysql:5.6
```

## Getting Started with Development

```bash
# Start the UI.
cd ui
npm install
npm run dev

# Start the API.
cd api
go run main.go
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