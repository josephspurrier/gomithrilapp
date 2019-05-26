package boot

import (
	"net/http"

	"app/api/endpoint"
	"app/api/middleware"
)

// Middleware wraps the routes.
func Middleware(core endpoint.Core) http.Handler {
	return middleware.Wrap(core.Router, core.Log, core.Token.Secret(), core.Context)
}
