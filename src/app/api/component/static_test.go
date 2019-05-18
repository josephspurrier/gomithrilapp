package component_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"app/api/component"
	"app/api/internal/testrequest"
	"app/api/model"

	"github.com/stretchr/testify/assert"
)

func TestIndexGET(t *testing.T) {
	core, _ := component.NewCoreMock(nil)

	w := testrequest.SendForm(t, core, "GET", "/v1", nil)

	r := new(model.OKResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", r.Body.Status)
}

func TestStaticGET(t *testing.T) {
	core, _ := component.NewCoreMock(nil)

	w := testrequest.SendForm(t, core, "GET", "/static/healthcheck.html", nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ok", w.Body.String())
}

func TestStaticGETNotFound(t *testing.T) {
	core, _ := component.NewCoreMock(nil)

	w := testrequest.SendForm(t, core, "GET", "/static/healthcheck-bad.html", nil)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "404 page not found\n", w.Body.String())
}

func TestStaticGETDirNotFound(t *testing.T) {
	core, _ := component.NewCoreMock(nil)

	w := testrequest.SendForm(t, core, "GET", "/static/folder/", nil)

	r := new(model.GenericResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, http.StatusText(http.StatusNotFound), r.Body.Status)
}
