package testutil

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"app/api/model"

	"github.com/stretchr/testify/assert"
)

// EnsureOK tests for HTTP status 200.
func EnsureOK(t *testing.T, w *httptest.ResponseRecorder) *model.OKResponse {
	r := new(model.OKResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	return r
}

// EnsureCreated tests for http status 201.
func EnsureCreated(t *testing.T, w *httptest.ResponseRecorder) *model.CreatedResponse {
	r := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	return r
}

// EnsureBadRequest tests for HTTP status 400.
func EnsureBadRequest(t *testing.T, w *httptest.ResponseRecorder) *model.BadRequestResponse {
	r := new(model.BadRequestResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	return r
}

// EnsureUnauthorized tests for HTTP status 401.
func EnsureUnauthorized(t *testing.T, w *httptest.ResponseRecorder) *model.UnauthorizedResponse {
	r := new(model.UnauthorizedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	return r
}

// EnsureInternalServerError tests for http status 500.
func EnsureInternalServerError(t *testing.T, w *httptest.ResponseRecorder) *model.InternalServerErrorResponse {
	r := new(model.InternalServerErrorResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	return r
}
