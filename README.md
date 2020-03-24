# govueapp

[![Go Report Card](https://goreportcard.com/badge/github.com/josephspurrier/govueapp)](https://goreportcard.com/report/github.com/josephspurrier/govueapp)
[![Build Status](https://travis-ci.org/josephspurrier/govueapp.svg)](https://travis-ci.org/josephspurrier/govueapp)
[![Coverage Status](https://coveralls.io/repos/github/josephspurrier/govueapp/badge.svg?branch=master&timestamp=20200313-01)](https://coveralls.io/github/josephspurrier/govueapp?branch=master)

[![Swagger Validator](https://online.swagger.io/validator?url=https://raw.githubusercontent.com/josephspurrier/govueapp/master/src/app/api/static/swagger/swagger.json)](https://petstore.swagger.io/?url=https://raw.githubusercontent.com/josephspurrier/govueapp/master/src/app/api/static/swagger/swagger.json)

This is a sample notepad application that uses Vue on the front-end (UI) and Go on the back-end (API). This project is designed to show good development and CI/CD practices as well as integrations between modern development tools.

[Documentation](https://josephspurrier.github.io/govueapp/) is generated using [Docusaurus](https://docusaurus.io/) and hosted using [GitHub Pages](https://pages.github.com/). This project uses a [Makefile](Makefile) to centralize frequently used commands. The code coverage badge above is only for the back-end API - not the front-end.

![Demo gif](https://user-images.githubusercontent.com/2394539/76177148-ac753e00-6189-11ea-963b-bff38b29e8ed.gif)

## Quick Start Guide (QSG)

To run the application locally, you can run these commands. You don't need any of the the dev tools (Go/npm) installed, you only need Docker (and Docker Compose). The ingress container is a reverse proxy that allows you to use a single URL to access multiple different docker containers depending on the URL.

```bash
# Clone the repo.
git clone git@github.com:josephspurrier/govueapp.git

# CD to the project directory.
cd govueapp

# Build the docker containers.
make docker-build

# Run the docker containers: ingress, DB, API, and UI.
docker-compose up -d

# Open your browser to the UI (via ingress): http://localhost
# Open your browser to the API (via ingress): http://localhost/api
# Open your MySQL tool to the DB: localhost:3306 (root:password)

# Stop and remove the docker containers.
docker-compose down
```

## Additional Documentation

Below are links to various section of the documentation.

- [Environment Preparation](https://josephspurrier.github.io/govueapp/docs/tutorial/env-prep)
- [Database](https://josephspurrier.github.io/govueapp/docs/database)
- [Front-End](https://josephspurrier.github.io/govueapp/docs/front-end)
- [Back-End](https://josephspurrier.github.io/govueapp/docs/back-end)
- [Swagger](https://josephspurrier.github.io/govueapp/docs/swagger)
- [Docker Compose](https://josephspurrier.github.io/govueapp/docs/docker-compose)
- [Documentation](https://josephspurrier.github.io/govueapp/docs/documentation)