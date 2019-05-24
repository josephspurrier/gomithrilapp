package endpoint_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"app/api/boot"
	"app/api/internal/testrequest"
	"app/api/model"

	"github.com/stretchr/testify/assert"
)

func TestIndexGET(t *testing.T) {
	p, _ := boot.TestServices(nil)

	w := testrequest.SendForm(t, p, "GET", "/v1", nil)

	r := new(model.OKResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", r.Body.Status)
}

func TestStaticGET(t *testing.T) {
	p, _ := boot.TestServices(nil)

	w := testrequest.SendForm(t, p, "GET", "/static/healthcheck.html", nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ok", w.Body.String())
}

func TestStaticGETNotFound(t *testing.T) {
	p, _ := boot.TestServices(nil)

	w := testrequest.SendForm(t, p, "GET", "/static/healthcheck-bad.html", nil)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "404 page not found\n", w.Body.String())
}

func TestStaticGETDirNotFound(t *testing.T) {
	p, _ := boot.TestServices(nil)

	w := testrequest.SendForm(t, p, "GET", "/static/folder/", nil)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "404 page not found\n", w.Body.String())
}
