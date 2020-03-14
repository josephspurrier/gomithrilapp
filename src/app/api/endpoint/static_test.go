package endpoint_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"app/api/internal/testrequest"
	"app/api/internal/testutil"
	"app/api/model"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	p, _ := testutil.Services(nil)
	tr := testrequest.New()

	// Home route.
	w := tr.SendForm(t, p, "GET", "/api/v1", nil)
	r := new(model.OKResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", r.Body.Status)
}

func TestStatic(t *testing.T) {
	p, _ := testutil.Services(nil)
	tr := testrequest.New()

	// Success.
	w := tr.SendForm(t, p, "GET", "/api/static/healthcheck.html", nil)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ok", w.Body.String())

	// Not exist serve content.
	w = tr.SendForm(t, p, "GET", "/api/static/healthcheck-bad.html", nil)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "404 page not found\n", w.Body.String())

	// Not exist route.
	w = tr.SendForm(t, p, "GET", "/api/static/", nil)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "{\"status\":\"Not Found\"}\n", w.Body.String())

	// Not exist folder.
	w = tr.SendForm(t, p, "GET", "/api/static/folder/", nil)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "404 page not found\n", w.Body.String())
}
