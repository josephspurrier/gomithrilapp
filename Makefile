# Load the shared environment variables (shared with docker-compose.yml).
include ${GOPATH}/.env

# Set local environment variables.
MYSQL_NAME=mysql56

.PHONY: nuxt-upgrade
nuxt-upgrade:
	# Upgrade nuxt to the 1.0 version.
	$(shell npm bin)/npm upgrade nuxt

.PHONY: nuxt-version
nuxt-version:
	# Output the version of nuxt.
	$(shell npm bin)/nuxt --version

.PHONY: docker-build
docker-build:
	# Build the docker containers.
	bash ${GOPATH}/bash/build-containers.sh

.PHONY: dev-ui
dev-ui:
	# Start the UI.
	cd ${GOPATH}/src/app/ui
	npm install
	npm run dev

.PHONY: dev-api
dev-api:
	# Start the API.
	cd ${GOPATH}/src/app/api
	go run main.go

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
	# CD to the api folder.
	cd ${GOPATH}/src/app/api/cmd/api

	# Generate the swagger spec.
	swagger generate spec -o ${GOPATH}/src/app/api/static/swagger/swagger.json

	# Replace 'example' with 'x-example' in the swagger spec.
	## MacOS
	sed -i '' -e 's/example/x\-example/' ${GOPATH}/src/app/api/static/swagger/swagger.json
	## Linux
	sed -i'' -e 's/example/x\-example/' ${GOPATH}/src/app/api/static/swagger/swagger.json

	# Validate the swagger spec.
	swagger validate ${GOPATH}/src/app/api/static/swagger/swagger.json

	# Serve the spec for the browser.
	swagger serve -F=swagger ${GOPATH}/src/app/api/static/swagger/swagger.json

.PHONY: clean
clean:
	rm -rf ${GOPATH}/src/app/api/cmd/api/api
	rm -rf ${GOPATH}/src/app/api/cmd/dbmigrate/dbmigrate

.PHONY: db-init
db-init:
	docker run -d --name=${MYSQL_NAME} -p 3306:3306 -e MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} ${MYSQL_CONTAINER}

.PHONY: db-start
db-start: 
	docker start ${MYSQL_NAME}

.PHONY: db-stop
db-stop:
	docker stop ${MYSQL_NAME}

.PHONY: db-reset
db-reset:
	docker exec ${MYSQL_NAME} sh -c "exec mysql -h 127.0.0.1 -uroot -p${MYSQL_ROOT_PASSWORD} -e 'DROP DATABASE IF EXISTS main;'"
	docker exec ${MYSQL_NAME} sh -c "exec mysql -h 127.0.0.1 -uroot -p${MYSQL_ROOT_PASSWORD} -e 'CREATE DATABASE IF NOT EXISTS main DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;'"
	go run ${GOPATH}/src/app/api/cmd/dbmigrate/main.go

.PHONY: db-rm
db-rm:
	docker rm -f ${MYSQL_NAME}