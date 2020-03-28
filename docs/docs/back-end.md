---
id: back-end
title: Back-End
---

## Go

You should use Go 1.11 or newer. We recommend [gvm](https://github.com/moovweb/gvm) for installing and managing your versions of Go.

All of the commands below assume you have your `GOPATH` set to the root of the project directory. This is by design because we found (after many projects) it is much easier for you to clone the repo and make changes without having to rewrite imports if they are all contained within the project.

We also recommend [A Tour of Go](https://tour.golang.org/) if you're new to the language. Follow this guide on [How to Write Go Code](https://golang.org/doc/code.html). You can also read the [Go Language Specification](https://golang.org/ref/spec) in an afternoon to see all the constructs in the language. The last stop would be to read through the [standard library](https://golang.org/pkg/) to see what's available out of the box.

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

The `gvt restore` command will then download all the dependencies to the **vendor** directory.

```bash
# Makefile
make api-dep

# Manual
cd ${GOPATH}/src/app/api
gvt restore
```

### Add a Dependency

Make sure you CD to the **api** folder before using gvt.

```bash
# Manual
cd ${GOPATH}/src/app/api
gvt fetch github.com/user/project
```

## Folder Structure

If you have a simple project, you don't need to worry about the folder structure upfront. All your Go code can fit in a single **main.go** file in the root of your project. As your project grows, it's a good idea to standardize on the folder structure so that the Go community can easily `go get` your library and collaborate.

There are many design patterns available for web applications like [Model-View-Controller (MVC)](https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller) and [Action-Domain-Responder (ADR)](https://en.wikipedia.org/wiki/Action%E2%80%93domain%E2%80%93responder). We are actually using more of the [Package Oriented Design](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html) pattern.

### cmd

The [cmd](https://github.com/josephspurrier/govueapp/tree/master/src/app/api/cmd) folder is where the code containing the `func main()` should go. Since a project may contain multiple compilable applications, you should create a folder with the name of the application and then place the file containing the `main` function in that directory. These application folders should typically only contain a single file. It's also a good practice to name the file `main.go` so others know it contains the `main` function.

In this example, assume there are two applications:
- Hello World
- Foobar

Here is an example folder structure for the applications:

```plaintext
cmd/hello-world/main.go
cmd/foobar/main.go
```

Syntax suggestions:
- folders and filenames should be in all lowercase
- any spaces should be replaced by hypens

### config

The [config](https://github.com/josephspurrier/govueapp/tree/master/src/app/api/config) folder is where the code that set up the application should go. This includes the code that connects to the database, loads the environment variables, and generally sets up the services needed in the application. You should try to remove as much of the logic from the **main.go** file so your application is more testable. This folder doesn't contain any tests because all of the functions are tested by other parts of the project.

Example folder structure:

```plaintext
config/database.go    - set up the connection pool to the database
config/env.go         - load the environment variables
config/middleware.go  - wrap the router with middleware
config/route.go       - load the routes
config/router.go      - initialize the router
config/service.go     - set up the remainder of the services
```

### endpoint

The [endpoint](https://github.com/josephspurrier/govueapp/tree/master/src/app/api/endpoint) folder is where code that sets up the each of the HTTP handlers and route definition should go. Each of the handlers should also have a respective test file.

Example folder structure:

```plaintext
endpoint/core.go          - core services used by all of the handlers
endpoint/login.go         - user login handlers
endpoint/login_test.go    - tests for the user login handlers
endpoint/note.go          - note handlers
endpoint/note_test.go     - tests for the note handlers
endpoint/register.go      - user registration handlers
endpoint/register_test.go - tests for the user registration handlers
endpoint/static.go        - serves static code
endpoint/static_test.go   - tests for serving static code
```

### internal

The [internal](https://github.com/josephspurrier/govueapp/tree/master/src/app/api/internal) folder is where packages that are imported by the project should go. Any folder named **internal** is only allowed to be imported by packages at the same level or deeper. This prevents other libraries from referencing the packages. You'll notice that all of these packages are very specific to this project and can't be easily moved to another project. A good rule of thumb is if the package imports other packages in your project, it should go in **internal** if there is no other place for it.

Example folder structure:

```plaintext
internal/query/          - provides basic queries for MySQL
internal/requestcontext/ - provides ability to set and retrieve variables set on
                           the request context
internal/response/       - provides the helpers to convert structs to JSON and
                           then write them to a http.ResponseWriter
internal/testutil/       - provides all the test utilities for the project
                           including mocks, test DB setup/teardown, and other
                           commonly used helpers
```

### middleware

The [middleware](https://github.com/josephspurrier/govueapp/tree/master/src/app/api/middleware) folder is where packages that "wrap" around the router or routes should go. In Go, examples of middleware are packages that log incoming requests, deny/grant access using access control lists (ACLs), and append headers to permit/restrict web resources through cross-origin resource sharing (CORS). These packages are typically called before or after the HTTP handlers.

Example folder structure:

```plaintext
middleware/cors/       - provides CORS support
middleware/jwt/        - provides JSON web token (JWT) validation
middleware/logrequest/ - logs incoming requests
middleware/factory.go  - returns a http.Handler wrapped in all of the
                         middleware
```

### model

The [model](https://github.com/josephspurrier/govueapp/tree/master/src/app/api/model) folder is where all the response structs should go.

Example folder structure:

```plaintext
model/generic.go - contains the generic responses
model/login.go   - contains the login responses
model/note.go    - contains the note responses
```

### pkg

The [pkg](https://github.com/josephspurrier/govueapp/tree/master/src/app/api/pkg) folder is where the packages that can be easily be moved to a new project should go. The imports in these packages should only reference the Go standard library and other third-party packages. If any of the packages reference other packages in the project, they should be moved to the **internal** folder instead.

Example folder structure:

```plaintext
pkg/bind/      - unmarshals JSON and validates requests
pkg/database/  - connects to the database and performs migrations
pkg/env/       - loads environment variables
pkg/logger/    - logging capabilities
pkg/mock/      - provides mocking capabilities
pkg/passhash/  - password hashing and validation
pkg/router/    - sets up the router
pkg/securegen/ - unique ID generation
pkg/webtoken/  - JWT generation and validation
```

### static

The [static](https://github.com/josephspurrier/govueapp/tree/master/src/app/api/static) folder contains static assets like HTML, CSS, JavaScript, and images.

### store

The [store](https://github.com/josephspurrier/govueapp/tree/master/src/app/api/store) folder contains the files that interact with the MySQL datbase.

Example folder structure:

```plaintext
store/core.go      - core services used by all of the handlers
store/factory.go   - returns an object that contains all of the stores
store/note.go      - stores and retrieves notes from the database
store/note_test.go - tests the notes code
store/user.go      - stores and retrieves users from the database
store/user_test.go - tests the user code
```

## Abstraction

You'll notice many of the packages in the **pkg** folder are wrappers around third-party packages. There are a few advantages to creating these layers of abstraction:

- you are not dependent on the public interfaces of the third-party package
- you can easily swap out the third-party package
- you can extended functionality of the third-party package

Even with these advantages, you may be [overengineering](https://solidstudio.io/blog/origin-of-overengineering.html) your solution. You may not need abstractions in all cases - especially if the public interfaces are simple or unchanging throughout the life of the project.

## Interfaces

We use interfaces in this project primarily to increase testability. Most of the interfaces are in the root of the **api** folder in the [interface.go](https://github.com/josephspurrier/govueapp/blob/master/src/app/api/interface.go) file. By placing the interfaces at the top, you can use them by all packages below without the worry of circular dependencies (when one package imports another package that imports the first package). Circular dependencies are not allowed in Go and will throw an error at build time.

Here are a few of the interfaces:

```go
// IRecord provides table information.
type IRecord interface {
	Table() string
	PrimaryKey() string
}

// IResponse provides outputs for data.
type IResponse interface {
	JSON(w http.ResponseWriter, body interface{}) (int, error)
	Created(w http.ResponseWriter, recordID string) (int, error)
	OK(w http.ResponseWriter, message string) (int, error)
}

// IToken provides outputs for the JWT.
type IToken interface {
	Generate(userID string) (string, error)
	Verify(s string) (string, error)
}

// IContext provides handlers for type request context.
type IContext interface {
	SetUserID(r *http.Request, val string)
	UserID(r *http.Request) (string, bool)
}
```

To help distinguish between interfaces, we've add a capital `I` to the beginning of each one. This is not a [standard Go convention](https://golang.org/doc/effective_go.html#interface-names), but we like it because it's easy to see if we've used an interface or not.

## Middleware



## Routing

Go has a [built-in router](https://golang.org/pkg/net/http/#ServeMux), but it doesn't support path parameters. For this project, we selected [Way](https://github.com/matryer/way) because it's "deliberately simple" and "extremely fast".

In addtion to using a third-party router, we are also using a custom `http.Handler` inspired by [Caddy](https://github.com/caddyserver/caddy/wiki/Writing-a-Plugin:-HTTP-Middleware#writing-a-handler). The [h](https://github.com/josephspurrier/h) project is an example of how to extend the HTTP handler. This convention forces you to return both the HTTP status and an optional error so it's easily to see what the response will be for each request.

The router logic is configured in the [router.go](https://github.com/josephspurrier/govueapp/blob/master/src/app/api/config/router.go) file. You can see the type of response is based on the status code. Any status code below 400 leaves it up to the calling function to output a response. Any status code 400 and above will return an error and then any status code of 500 and above will log an error since it's probably a bug or system error. This greatly simplifies logic that is otherwise scattered throughout codebases.

```go
// Set the handling of all responses.
mux.CustomServeHTTP = func(w http.ResponseWriter, r *http.Request, status int, err error) {
    // Handle only errors.
    if status >= 400 {
        resp := new(model.GenericResponse)
        resp.Body.Status = http.StatusText(status)
        if err != nil {
            resp.Body.Message = err.Error()
        }

        // Write the content.
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(status)
        err := json.NewEncoder(w).Encode(resp.Body)
        if err != nil {
            w.Write([]byte(`{"status":"Internal Server Error",` +
                `"message":"problem encoding JSON"}`))
            return
        }
    }

    // Display server errors.
    if status >= 500 {
        if err != nil {
            l.Printf("%v", err)
        }
    }
}
```


### Defining the Route

You'll find the routes at the top of every file in the [endpoints](https://github.com/josephspurrier/govueapp/tree/master/src/app/api/endpoint) folder. Below is an example of the routes in the [note.go](https://github.com/josephspurrier/govueapp/blob/master/src/app/api/endpoint/note.go) endpoint file.

```go
// NoteEndpoint .
type NoteEndpoint struct {
	Core
}

// SetupNotepad .
func SetupNotepad(c Core) {
	p := new(NoteEndpoint)
	p.Core = c

	p.Router.Post("/api/v1/note", p.Create)
	p.Router.Get("/api/v1/note", p.Index)
	p.Router.Get("/api/v1/note/:note_id", p.Show)
	p.Router.Put("/api/v1/note/:note_id", p.Update)
	p.Router.Delete("/api/v1/note/:note_id", p.Destroy)
}
```

### Path Parameters

The router supports prefixing path parameter with a colon and then you can retrieve them like using the `Bind` package. You need to make sure the `json` tag for the field matches the variable name. It must also have the `in: path` annotation for the `Bind` package to extract it.

```go
// swagger:parameters NoteShow
type request struct {
    // in: path
    NoteID string `json:"note_id" validate:"required"`
}

// Request validation.
req := new(request)
if err := p.Bind.UnmarshalAndValidate(req, r); err != nil {
    return http.StatusBadRequest, err
}
```

You can alternatively return the path parameter like this:

```go
noteID := p.Router.Param(r, "note_id")
```

## Endpoints

Each of the endpoint files contain the functions that process requests and return responses. These are the API endpoints. For these examples, we'll reference the [note.go](https://github.com/josephspurrier/govueapp/blob/master/src/app/api/endpoint/note.go) file.

### Imports

Notice there are only a few imports - the majority of the functionality we are using in the endpoint code is from the `endpoint.Core`. This makes it easier for us to write tests for our endpoints because we can control the dependencies that are injected into the functions.

```go
import (
	"errors"
	"net/http"

	"app/api/model"
	"app/api/pkg/structcopy"
)
```

### Core

Each endpoint has an anonymous `Core` struct inside of it. This provides all the methods with the same core functionality and prevents naming collisions for methods that are in the same package.

```go
// NoteEndpoint .
type NoteEndpoint struct {
	Core
}
```

### Setup

You should also have a setup function at the top of the endpoint file as well. Each of these setup functions should be called by the [route.go](https://github.com/josephspurrier/govueapp/blob/master/src/app/api/config/route.go) file that is in the **config** directory. The setup function should always take in `endpoint.Core`, assign it, and then register each fo the routes on the `Router`. The router is a pointer so there is nothing the needs to be returned.

```go
// SetupNotepad .
func SetupNotepad(c Core) {
	p := new(NoteEndpoint)
	p.Core = c

	p.Router.Post("/api/v1/note", p.Create)
	p.Router.Get("/api/v1/note", p.Index)
	p.Router.Get("/api/v1/note/:note_id", p.Show)
	p.Router.Put("/api/v1/note/:note_id", p.Update)
	p.Router.Delete("/api/v1/note/:note_id", p.Destroy)
}
```

### Handler

Each handler method should take in the standard `http.ResponseWriter` and `*http.Request` and then return the status code and an optional error.

```go
func (p *NoteEndpoint) Create(w http.ResponseWriter, r *http.Request) (int, error) {
	// swagger:parameters NoteCreate
	type request struct {
		// in: body
		Body struct {
			Message string `json:"message"`
		}
	}

	// Request validation.
	req := new(request)
	if err := p.Bind.UnmarshalAndValidate(req, r); err != nil {
		return http.StatusBadRequest, err
	}

	// Get the user ID.
	userID, ok := p.Context.UserID(r)
	if !ok {
		return http.StatusInternalServerError, errors.New("invalid user")
	}

	// Create the note.
	ID, err := p.Store.Note.Create(userID, req.Body.Message)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return p.Response.Created(w, ID)
}
```

## Requests

## Models

## Responses

## Testing Methodologies

Don't want to much testing. Don't want to go for code coverage alone.

WIP.