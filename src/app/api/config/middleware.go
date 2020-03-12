package config

import (
	"net/http"

	"app/api/endpoint"
	"app/api/middleware"
)

// Middleware wraps the router with middleware and returns a http.Handler.
func Middleware(core endpoint.Core) http.Handler {
	return middleware.Factory(core.Router,
		core.Log,
		[]string{ // JWT whitelist.
			"GET /favicon.ico",
			"GET /v1",
			"GET /static/*",
			"POST /v1/login",
			"POST /v1/register",
		},
		core.Token,
		core.Context)
}
