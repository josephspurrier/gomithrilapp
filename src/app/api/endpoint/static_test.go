package endpoint_test

import (
	"net/http"
	"testing"

	"app/api/internal/testutil"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	p, _ := testutil.Services(nil)
	tr := testutil.NewRequest()

	// Home route.
	w := tr.SendForm(t, p, "GET", "/api/v1", nil)
	r := testutil.EnsureOK(t, w)
	assert.Equal(t, "OK", r.Body.Status)
}

func TestStatic(t *testing.T) {
	p, _ := testutil.Services(nil)
	tr := testutil.NewRequest()

	// Success.
	w := tr.SendForm(t, p, "GET", "/static/healthcheck.html", nil)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ok", w.Body.String())

	// Not exist serve content.
	w = tr.SendForm(t, p, "GET", "/static/not-exist.html", nil)
	assert.Equal(t, http.StatusNotFound, w.Code)

	// Not exist route.
	w = tr.SendForm(t, p, "GET", "/api/static/", nil)
	assert.Equal(t, http.StatusNotFound, w.Code)

	// Not exist folder.
	w = tr.SendForm(t, p, "GET", "/api/static/folder/", nil)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
