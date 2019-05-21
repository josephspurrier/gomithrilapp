package response_test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"app/api/internal/response"
	"app/api/model"

	"github.com/stretchr/testify/assert"
)

func TestJSONSuccess(t *testing.T) {
	w := httptest.NewRecorder()

	in := new(model.OKResponse)
	in.Body.Message = "hello"
	in.Body.Status = "nice"

	out := response.New()
	status, err := out.JSON(w, in.Body)
	assert.Equal(t, http.StatusOK, status)
	assert.Nil(t, err)

	c := new(model.OKResponse).Body
	err = json.Unmarshal(w.Body.Bytes(), &c)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "nice", c.Status)
	assert.Equal(t, "hello", c.Message)
}

func TestJSONFail(t *testing.T) {
	w := httptest.NewRecorder()

	out := response.New()
	status, err := out.JSON(w, make(chan int))
	assert.Equal(t, http.StatusInternalServerError, status)
	assert.NotNil(t, err)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "", w.Body.String())
}

func TestOKSuccess(t *testing.T) {
	w := httptest.NewRecorder()

	out := response.New()
	status, err := out.OK(w, "test")
	assert.Equal(t, http.StatusOK, status)
	assert.Nil(t, err)

	c := new(model.OKResponse).Body
	err = json.Unmarshal(w.Body.Bytes(), &c)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, http.StatusText(http.StatusOK), c.Status)
	assert.Equal(t, "test", c.Message)
}

func TestCreated(t *testing.T) {
	w := httptest.NewRecorder()

	out := response.New()
	status, err := out.Created(w, "1")
	assert.Equal(t, http.StatusCreated, status)
	assert.Nil(t, err)

	log.Println(w.Code, w.Body.String())
	c := new(model.CreatedResponse).Body
	err = json.Unmarshal(w.Body.Bytes(), &c)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, http.StatusText(http.StatusCreated), c.Status)
	assert.Equal(t, "1", c.RecordID)
}
