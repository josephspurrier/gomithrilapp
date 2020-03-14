package requestcontext

import (
	"context"
	"net/http"
)

var (
	keyUserID = contextKey("user_id")
)

type contextKey string

// Context is a request context handler.
type Context struct{}

// New returns a new context handler.
func New() Context {
	return Context{}
}

// SetUserID will set the user ID in the context.
func (ctx Context) SetUserID(r *http.Request, val string) {
	*r = *r.WithContext(context.WithValue(r.Context(), keyUserID, val))
}

// UserID gets the user ID from the context.
func (ctx Context) UserID(r *http.Request) (string, bool) {
	val, ok := r.Context().Value(keyUserID).(string)
	return val, ok
}
