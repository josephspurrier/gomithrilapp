---
id: back-end
title: Back-End
---

## Go Dependency Management

This projects does not use Go modules - it uses [gvt](https://github.com/FiloSottile/gvt/blob/master/README.old.md) to vendor dependencies to Go. This decision was made because Visual Studio Code support is still lacking and that just happens to be our preferred IDE: ["⚠️ These tools do not provide a good support for Go modules yet."](https://github.com/Microsoft/vscode-go/wiki/Go-modules-support-in-Visual-Studio-Code). We've used gvt on large teams for years so even though it's deprecated, it still works extremely well for our purposes.

```bash
# Install gvt and download all dependencies to the vendor directory.
make api-dep

# You can now remove the folder: src/github.com/FiloSottile/gvt
# You should now add the {PROJECTROOT}/bin folder to your $PATH to make gvt
# available from your terminal.
```

These are other commands you can use:

```bash
# Download gvt.
make gvt-get

# You can now remove the folder: src/github.com/FiloSottile/gvt
# You should now add the {PROJECTROOT}/bin folder to your $PATH to make gvt
# available from your terminal.

# Make sure you CD to the api folder before using gvt.
cd $GOPATH/src/app/api

# Here is a sample command to add a new dependency to the project.
gvt fetch github.com/user/project

# Here is the command to download all dependencies to the vendor directory.
gvt restore
```

## Swagger for the API

This projects uses [Swagger v2](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md) to document the API. The entire Swagger spec is generated from comments (annotations) in and by analyzing structs and variables.

```bash
# Download the Swagger generation tool.
make swagger-get

# You can now remove all the folders from the 'src' directory except the
# 'app' folder.
# You should now add the {PROJECTROOT}/bin folder to your $PATH to make swagger
# available from your terminal.

# Generate the swagger spec.
make swagger-gen

# Your browser will open to:
# http://petstore.swagger.io/?url=http://localhost:{RANDOMPORT}/swagger.json

# The output file will be here:
# src/app/api/static/swagger/swagger.json
```