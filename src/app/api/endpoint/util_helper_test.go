package endpoint_test

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"app/api/endpoint"
	"app/api/internal/testrequest"
	"app/api/model"

	"github.com/stretchr/testify/assert"
)

// Register a user and return the user ID.
func Register(t *testing.T, tr *testrequest.TR, p endpoint.Core) (userID string) {
	// Register the user.
	form := url.Values{}
	form.Set("first_name", "Foo")
	form.Set("last_name", "Barr")
	form.Set("email", "fbar@example.com")
	form.Set("password", "guess123")
	w := tr.SendJSON(t, p, "POST", "/api/v1/register", form)

	// Verify the response.
	r := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.NotEmpty(t, r.Body.RecordID)

	return r.Body.RecordID
}

// Auth will register a user, login, and return the user token and user ID.
func Auth(t *testing.T, tr *testrequest.TR, p endpoint.Core) (userToken string, userID string) {
	userID = Register(t, tr, p)

	// Login with the user.
	form := url.Values{}
	form.Set("email", "fbar@example.com")
	form.Set("password", "guess123")
	w := tr.SendJSON(t, p, "POST", "/api/v1/login", form)

	r := new(model.LoginResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	return r.Body.Token, userID
}
