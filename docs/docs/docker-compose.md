---
id: docker-compose
title: Docker Compose
---

## Docker Compose

[Compose](https://docs.docker.com/compose/) is a tool that helps you easily run multi-container Docker applications. The configuration is defined using a YAML file and you have control over different aspects of the containers like mountpoints, networks, and environment variables.

You should ensure you have [Docker Compose installed](https://docs.docker.com/compose/install/). You probably already do if you have [Docker installed](https://www.docker.com/get-started).

The commands below should be run in the root of the project directory. The command will use the **docker-compose.yml** file in the current directory if it's found.

### Build the Containers

Build the docker containers.

```bash
# Makefile
make docker-build

# Manual
bash ${GOPATH}/bash/build-images.sh
```

### Start the Containers

Create or start stopped docker containers in detached mode.

```bash
docker-compose up -d
```

### View Container Logs

View the logs of all the docker containers. You can also add `-f` to the end to follow the logs.

```bash
docker-compose logs
```

### Stop the Containers

Stop the docker containers but retain the data.

```bash
docker-compose down
```

### Remove the Containers

Stop and delete the docker containers and all data.

```bash
docker-compose rm -f
```

## docker-compose.yml

The configuration is defined in the [**docker-compose.yml**](https://github.com/josephspurrier/gomithrilapp/blob/master/docker-compose.yml) file.

### Version

The version of the configuration is the at the top of the document. We chose [3.4](https://docs.docker.com/compose/compose-file/compose-versioning/#version-34) because it's the first version where extension fields are support - they allow us to reuse configuration fragments which is covered in the next section. Version 3.4 only supports Docker Engine version 17.09.0 and higher.

### Templates

This configuration fragment is essentially a template that we'll use for all of our services. This helps reduce typing and allows us to centralize all of our configurations. We'll define a few fields:

- `restart: always` - this ensures the containers will restarted if they stop
- `env_file` - this tells docker that we want to load the `.env` for all the containers and set them as environment variables
- `networks` - this defines the network `dnet` that all the services will connect to so they can communicate

We are using the [.env](https://github.com/josephspurrier/gomithrilapp/blob/master/.env) file for two purposes:
- [declaring default environment variables](https://docs.docker.com/compose/env-file/) for the docker-compose.yml file
- [declaring environment variables](https://docs.docker.com/compose/environment-variables/) for each of the services (containers)

By default, Compose will pull the variables from the **.env** file if it's in the current directory. We also pass it to the services so the containers can reference the environment variables as well.

#### Makefile

As an added benefit, at the top of the [Makefile](https://github.com/josephspurrier/gomithrilapp/blob/master/Makefile), we also load in the same environment variables from the **.env** file so we can share them across our local environment and our container environments. Now we only have to define our variables in a single place and we can use the variables across all of our files. Here is what is at the top of the **Makefile**:

```bash
# Load the shared environment variables (shared with docker-compose.yml).
include ${GOPATH}/.env
```

### Networks

By default, a network is created for all the containers so they can communicate. To support more complex topologies and because it's [superior to the default network](https://docs.docker.com/network/bridge/), we'll define the [network](https://docs.docker.com/compose/networking/) and use the `bridge` driver. We'll call the network, `dnet`, which we chose to stand for "docker network".

### Services

The services define which images are used to spin up containers for them. The images are all built using the [build-images.sh](https://github.com/josephspurrier/gomithrilapp/blob/master/bash/build-images.sh) bash file. Again - notice the **.env** file is sourced before building the docker images.

We have four services defined:

- `ingress` - this is a Go proxy that routes traffic that requests any URL to `/api` to the `api` container and any other requests to the `ui` container. The code for the ingress is [here](https://github.com/josephspurrier/gomithrilapp/blob/master/src/app/ingress/main.go) with the respective [Dockerfile](https://github.com/josephspurrier/gomithrilapp/blob/master/src/app/ingress/Dockerfile).
- `db` - this is the MySQL database. It's using the standard [mysql:5.7](https://hub.docker.com/_/mysql) Dockerfile.
- `ui` - this is the UI written in Mithril. The code for the UI is [here](https://github.com/josephspurrier/gomithrilapp/tree/master/src/app/ui) with the respective [Dockerfile](https://github.com/josephspurrier/gomithrilapp/tree/master/src/app/ui/Dockerfile).
- `api` - this is the Go API. The code for the API is [here](https://github.com/josephspurrier/gomithrilapp/tree/master/src/app/api) with the respective [Dockerfile](https://github.com/josephspurrier/gomithrilapp/tree/master/src/app/api/Dockerfile).

Notice only the `ports` for the ingress and db containers are exposed - this is because the requests for the api and ui are proxied through the ingress so there is no reason to expose them. The benefit of the ingress (reverse-proxy) is to allow you to use a single URL, a single port, and a single SSL certification for the application in the future even though the ui and api are listening on different ports internally.

Also notice the `service-template` is set at the top of each service - this uses the template defined at the top of the file to set defaults. Any of the values in the template can be overwritten if needed below it in the service.