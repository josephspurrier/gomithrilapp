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
