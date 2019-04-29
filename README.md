# govueapp

This is an application that uses Vue on the frontend (UI) and Go on the backend (API).

## Current Versions

These are the current versions:

- nuxt.js [v2.6.3](https://github.com/nuxt/nuxt.js/releases/tag/v2.6.3) - you should use the [2.5.X documentation](https://nuxtjs.org/guide/release-notes).
- Vue.js [v2.6.10](https://github.com/vuejs/vue/releases/tag/v2.6.10)
- vuex [v3.1.0](https://github.com/vuejs/vuex/releases/tag/v3.1.0)

```bash
# Upgrade nuxt to the latest version.
npm upgrade nuxt

# Run NPM apps from terminal.
export PATH=$PATH:$(npm bin)

# Check the version of nuxt.
nuxt --version
```

## Getting Started

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
```