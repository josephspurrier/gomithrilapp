---
id: run-locally
title: Run Locally
---

## Quick Start with Docker Compose

To run the application locally, you can run these commands. You don't need any of the the dev tools (Go/npm) installed, you only need Docker (and Docker Compose). The ingress container is a reverse proxy that allows you to use 1 URL to access 2 different docker containers by URL path.

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