package requestcontext

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserID(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)

	c := New()
	c.SetUserID(r, "foo")
	s, b := c.UserID(r)

	assert.Equal(t, "foo", s)
	assert.Equal(t, true, b)
}
