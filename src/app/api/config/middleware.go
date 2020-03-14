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
			"GET /api/favicon.ico",
			"GET /api/v1",
			"GET /api/static/*",
			"POST /api/v1/login",
			"POST /api/v1/register",
		},
		core.Token,
		core.Context)
}
