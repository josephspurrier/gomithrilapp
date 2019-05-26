package requestcontext

import (
	"context"
	"net/http"

	"app/api/pkg/mock"
)

var (
	keyUserID = contextKey("user_id")
)

type contextKey string

func (c contextKey) String() string {
	return "app" + string(c)
}

// CTX is a request context handler.
type CTX struct {
	Mock *mock.Mocker
}

// New returns a new context handler.
func New(m *mock.Mocker) CTX {
	return CTX{
		Mock: m,
	}
}

// SetUserID will set the user ID in the context.
func (ctx CTX) SetUserID(r *http.Request, val string) {
	*r = *r.WithContext(context.WithValue(r.Context(), keyUserID, val))
}

// UserID gets the user ID from the context.
func (ctx CTX) UserID(r *http.Request) (string, bool) {
	if ctx.Mock != nil && ctx.Mock.Enabled() {
		return ctx.Mock.String(), ctx.Mock.Bool()
	}

	val, ok := r.Context().Value(keyUserID).(string)
	return val, ok
}
