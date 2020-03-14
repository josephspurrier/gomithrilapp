package endpoint_test

import (
	"encoding/json"
	"net/url"
	"testing"

	"app/api/endpoint"
	"app/api/internal/testrequest"
	"app/api/model"

	"github.com/stretchr/testify/assert"
)

func register(t *testing.T, tr *testrequest.TR, p endpoint.Core) {
	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a")
	form.Set("last_name", "a")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	tr.SendJSON(t, p, "POST", "/api/v1/register", form)
}

func auth(t *testing.T, tr *testrequest.TR, p endpoint.Core) string {
	register(t, tr, p)

	// Login with the user.
	form := url.Values{}
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := tr.SendJSON(t, p, "POST", "/api/v1/login", form)

	r := new(model.LoginResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	return r.Body.Token
}
