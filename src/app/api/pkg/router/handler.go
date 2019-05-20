package router

import "net/http"

// handler is a internal handler.
type handler struct {
	Handler
	CustomServeHTTP func(w http.ResponseWriter, r *http.Request, status int, err error)
}

// ServeHTTP handles all the errors from the HTTP handlers.
func (fn handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := fn.Handler(w, r)
	fn.CustomServeHTTP(w, r, status, err)
}

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
