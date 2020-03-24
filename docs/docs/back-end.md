---
id: back-end
title: Back-End
---

## Go

You should use Go 1.11 or newer. We recommend [gvm](https://github.com/moovweb/gvm) for installing and managing your versions of Go.

All of the commands below assume you have your `GOPATH` set to the root of the project directory. This is by design because we found (after many projects) it is much easier for you to clone the repo and make changes without having to rewrite imports if they are all contained within the project.

We also recommend [A Tour of Go](https://tour.golang.org/) if you're new to the language. You can also read the [Go Spec](https://golang.org/ref/spec) in an afternoon to see all the constructs in the language. The last stop would be to read through the [standard library](https://golang.org/pkg/) to see what's available out of the box.

## Dependency Management

This projects does not use Go modules - it uses [gvt](https://github.com/FiloSottile/gvt/blob/master/README.old.md) to vendor dependencies to Go. This decision was made because Visual Studio Code support is still lacking and that just happens to be our preferred IDE: ["⚠️ These tools do not provide a good support for Go modules yet."](https://github.com/Microsoft/vscode-go/wiki/Go-modules-support-in-Visual-Studio-Code). We've used gvt on large teams for years so even though it's deprecated, it still works extremely well for our purposes.

### Download gvt

Download gvt using `go get`. You can then remove the folder: `$GOPATH/src/github.com/FiloSottile/gvt`. You can add the **{PROJECTROOT}/bin** folder to your `$PATH` to make gvt available from your terminal.

```bash
# Makefile
make gvt-get

# Manual
go get github.com/FiloSottile/gvt
```

### Fetch Dependencies

The `gvt restore` command will then download all the dependencies to the `vendor` directory.

```bash
# Makefile
make api-dep

# Manual
cd ${GOPATH}/src/app/api
gvt restore
```

### Add a Dependency

Make sure you CD to the `api` folder before using gvt.

```bash
# Manual
cd ${GOPATH}/src/app/api
gvt fetch github.com/user/project
```

## Folder Structure

### cmd
### config
### endpoint
### internal
### middlware
### model
### pkg
### static
### store

## Routing

## Middleware

## Endpoints

### Request Validation

## Models

WIP.