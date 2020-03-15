package testutil

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

// Context .
type Context struct {
	mock.Mock
}

// SetUserID .
func (m *Context) SetUserID(r *http.Request, userID string) {
	m.Called(userID)
}

// UserID .
func (m *Context) UserID(r *http.Request) (userID string, found bool) {
	args := m.Called()
	return args.String(0), args.Bool(1)
}
