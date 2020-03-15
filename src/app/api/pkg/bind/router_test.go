package bind_test

import (
	"net/http"

	"github.com/matryer/way"
)

// Handler is used to wrapper all endpoint functions so they work with generic
// routers.
type Handler func(http.ResponseWriter, *http.Request) (int, error)

// ServeHTTP handles all the errors from the HTTP handlers.
func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := fn(w, r)
	DefaultServeHTTP(w, r, status, err)
}

// DefaultServeHTTP is the default ServeHTTP function that receives the status and error from
// the function call.
var DefaultServeHTTP = func(w http.ResponseWriter, r *http.Request, status int,
	err error) {
	if status >= 400 {
		if err != nil {
			http.Error(w, err.Error(), status)
		} else {
			http.Error(w, "", status)
		}
	}
}

// Mux contains the router.
type Mux struct {
	*way.Router
}

// NewRouter returns a new router.
func NewRouter() *Mux {
	return &Mux{
		way.NewRouter(),
	}
}

// Param returns a URL parameter.
func (m *Mux) Param(r *http.Request, param string) string {
	return way.Param(r.Context(), param)
}
