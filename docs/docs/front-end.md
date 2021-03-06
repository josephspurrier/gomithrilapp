---
id: front-end
title: Front-End
---

## npm Dependency Management

Yuu can use these commands to interact with npm.

```bash
# CD to the UI folder.
cd ${GOPATH}/src/app/ui

# Check the versions of packages.
npm outdated

# Update all the packages to the latest version (specified by the tag config),
# respecting semver.
# https://docs.npmjs.com/cli-commands/update.html
npm update

# Scan your project for vulnerabilities and automatically install any compatible
# updates to vulnerable dependencies.
# https://docs.npmjs.com/cli/audit
npm audit fix

# Get a verion number and a list of all packages that rely on another package.
# https://docs.npmjs.com/cli/ls
npm ls babel-loader

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

# Use --save (-S) to add a package to the package.json dependencies (packages
# required when the app is built).
# Use --save-dev (-D) to add a package to the package.json devDependencies
# (packages used during dev to build, bundle, lint).
```

## Cypress for UI Testing

[Cypress](https://docs.cypress.io/guides/overview/why-cypress.html) is used to run tests on the UI. You can use the steps below to run the tests.

```bash
# CD to the UI folder.
cd ${GOPATH}/src/app/ui

# Run the Cypress tests headlessly.
npm run test

# Run the Cypress tests manually with a GUI.
npx cypress open
```

