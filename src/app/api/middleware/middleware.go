package middleware

import (
	"net/http"

	"app/api"
	"app/api/middleware/cors"
	"app/api/middleware/jwt"
	"app/api/middleware/logrequest"
)

// Wrap will return the http.Handler wrapped in middleware.
func Wrap(h http.Handler, l logrequest.ILog, secret []byte,
	ctx api.IContext) http.Handler {
	// JWT whitelist.
	whitelist := []string{
		"GET /favicon.ico",
		"GET /v1",
		"GET /static/*",
		"POST /v1/login",
		"POST /v1/register",
	}

	// JWT validation.
	token := jwt.New(secret, whitelist, ctx)
	h = token.Handler(h)

	// CORS for the endpoints.
	h = cors.Handler(h)

	// Log every request.
	lr := logrequest.New()
	lr.SetLog(l)
	h = lr.Handler(h)

	return h
}
