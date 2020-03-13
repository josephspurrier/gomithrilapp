# govueapp

> Sample notepad application in Vue and Go.

## Build Setup

``` bash
# Install dependencies.
go get github.com/FiloSottile/gvt
gvt restore

# Serve at localhost:8081.
go run cmd/api/main.go

# Build for production and launch server.
go build cmd/api/main.go
./main

# Run tests (must set up database first).
docker run -d --name=mysql57 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password mysql:5.7
go test ./...
docker rm -f mysql57

# Run coverage tests.
go test ./... -coverprofile cover.out && go tool cover -html=cover.out -o cover.html && open cover.html && sleep 5 && rm cover.html && rm cover.out
```
