---
id: run-locally
title: Run Locally
---

## Quick Start with Docker Compose

To test the application locally, you can run these commands. You don't need any of the the dev tools (Go/npm) installed, you only need Docker (and Docker Compose). The Go application serves both the UI and the API depending on the request URL.

```bash
# Clone the repo.
git clone git@github.com:josephspurrier/gomithrilapp.git

# CD to the project directory.
cd gomithrilapp

# Build the docker containers.
make docker-build

# Run the docker containers: DB and app.
docker-compose up -d

# Open your browser to the UI: http://localhost
# Open your browser to the API: http://localhost/api
# Open your MySQL tool to the DB: localhost:3306 (root:password)

# Stop and remove the docker containers.
docker-compose down
```