package router

import "net/http"

// Handler is a main handler.
type Handler struct {
	CustomHandler
	CustomServeHTTP func(w http.ResponseWriter, r *http.Request, status int, err error)
}

// ServeHTTP handles all the errors from the HTTP handlers.
func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := fn.CustomHandler(w, r)
	fn.CustomServeHTTP(w, r, status, err)
}

// CustomHandler is used to wrapper all endpoint functions so they work with generic
// routers.
type CustomHandler func(http.ResponseWriter, *http.Request) (int, error)

// ServeHTTP handles all the errors from the HTTP handlers.
func (fn CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
