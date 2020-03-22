# This Makefile is an easy way to run common operations.
# Execute commands this:
# * make db-init
# * make api-dep
# * make api-dev
# * make nuxt-version
#
# Tip: Each command is run on its own line so you can't CD unless you
# connect commands together using operators. See examples:
# A; B    # Run A and then B, regardless of success of A
# A && B  # Run B if and only if A succeeded
# A || B  # Run B if and only if A failed
# A &     # Run A in background.
# Source: https://askubuntu.com/a/539293
#
# Tip: Use $(shell app param) syntax when expanding a shell return value.

# Load the shared environment variables (shared with docker-compose.yml).
include ${GOPATH}/.env

# Set local environment variables.
MYSQL_NAME=govueapp_db_1
GITHUB_USERNAME=josephspurrier

.PHONY: docker-build
docker-build:
	# Build the docker images.
	bash ${GOPATH}/bash/build-images.sh

.PHONY: ui-dep
ui-dep:
	# Install the dependencies.
	cd ${GOPATH}/src/app/ui && npm install

.PHONY: ui-dev
ui-dev:
	# Start the UI.
	cd ${GOPATH}/src/app/ui && npm run dev

.PHONY: ui-test
ui-test:
	# Run the Jest UI tests.
	cd ${GOPATH}/src/app/ui && npm test

.PHONY: api-dep
api-dep:
	# Restore the dependencies. Get gvt if it's not found in $PATH.
	which gvt || go get github.com/FiloSottile/gvt
	cd ${GOPATH}/src/app/api && gvt restore

.PHONY: api-dev
api-dev:
	# Start the API.	
	go run ${GOPATH}/src/app/api/cmd/api/main.go

.PHONY: api-test
api-test:
	# Run the Go tests.
	cd ${GOPATH}/src/app/api && go test ./...

.PHONY: clean
clean:
	# Remove binaries.
	rm -rf ${GOPATH}/src/app/api/cmd/api/api
	rm -rf ${GOPATH}/src/app/api/cmd/dbmigrate/dbmigrate

.PHONY: gvt-get
gvt-get:
	# Download gvt.
	go get github.com/FiloSottile/gvt

.PHONY: swagger-get
swagger-get:
	# Download the Swagger generation tool.
	go get github.com/go-swagger/go-swagger/cmd/swagger

.PHONY: swagger-gen
swagger-gen:
	# Generate the swagger spec.
	cd ${GOPATH}/src/app/api/cmd/api && \
	swagger generate spec -o ${GOPATH}/src/app/api/static/swagger/swagger.json

	# Replace 'example' with 'x-example' in the swagger spec.
	## MacOS
	sed -i '' -e 's/example/x\-example/' ${GOPATH}/src/app/api/static/swagger/swagger.json
	## Linux
	#sed -i'' -e 's/example/x\-example/' ${GOPATH}/src/app/api/static/swagger/swagger.json

	# Validate the swagger spec.
	swagger validate ${GOPATH}/src/app/api/static/swagger/swagger.json

	# Serve the spec for the browser.
	swagger serve -F=swagger ${GOPATH}/src/app/api/static/swagger/swagger.json

.PHONY: db-init
db-init:
	# Launch database container.
	docker run -d --name=${MYSQL_NAME} -p 3306:3306 -e MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} ${MYSQL_CONTAINER}

.PHONY: db-start
db-start:
	# Start the stopped database container.
	docker start ${MYSQL_NAME}

.PHONY: db-stop
db-stop:
	# Stop the running database container.
	docker stop ${MYSQL_NAME}

.PHONY: db-reset
db-reset:
	# Drop the database, create the database, and perform the migrations.
	docker exec ${MYSQL_NAME} sh -c "exec mysql -h 127.0.0.1 -uroot -p${MYSQL_ROOT_PASSWORD} -e 'DROP DATABASE IF EXISTS main;'"
	docker exec ${MYSQL_NAME} sh -c "exec mysql -h 127.0.0.1 -uroot -p${MYSQL_ROOT_PASSWORD} -e 'CREATE DATABASE IF NOT EXISTS main DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;'"
	go run ${GOPATH}/src/app/api/cmd/dbmigrate/main.go

.PHONY: db-rm
db-rm:
	# Stop and remove the database container.
	docker rm -f ${MYSQL_NAME}

.PHONY: nuxt-upgrade
nuxt-upgrade:
	# Upgrade nuxt to the latest version.
	cd ${GOPATH}/src/app/ui && npm upgrade nuxt

.PHONY: nuxt-version
nuxt-version:
	# Output the version of nuxt.
	${GOPATH}/src/app/ui/node_modules/.bin/nuxt --version

.PHONY: doc-publish
nuxt-version:
	# Push the docs to GitHub pages.
	cd ${GOPATH}/docs/website && \
	GIT_USER=${GITHUB_USERNAME} CURRENT_BRANCH=master USE_SSH=true npm run publish-gh-pages