package endpoint_test

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"app/api/internal/testrequest"
	"app/api/internal/testutil"
	"app/api/model"

	"github.com/stretchr/testify/assert"
)

func TestLoginSuccess(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	p, m := testutil.Services(db)
	tr := testrequest.New()

	p.Token = m.Token
	m.Token.GenerateFunc = func(userID string) (string, error) {
		enc := base64.StdEncoding.EncodeToString([]byte("0123456789ABCDEF0123456789ABCDEF"))
		return enc, nil
	}

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	tr.SendJSON(t, p, "POST", "/v1/register", form)

	// Login with the user.
	form = url.Values{}
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := tr.SendJSON(t, p, "POST", "/v1/login", form)

	// Verify the response.
	r := new(model.LoginResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "MDEyMzQ1Njc4OUFCQ0RFRjAxMjM0NTY3ODlBQkNERUY=", r.Body.Token)
}

func TestLoginFail(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	p, m := testutil.Services(db)
	tr := testrequest.New()

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	tr.SendJSON(t, p, "POST", "/v1/register", form)

	// Wrong password.
	form = url.Values{}
	form.Set("email", "a@a.com")
	form.Set("password", "wrong-password")
	w := tr.SendJSON(t, p, "POST", "/v1/login", form)
	r := new(model.BadRequestResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Missing password.
	form = url.Values{}
	form.Set("email", "a@a.com")
	w = tr.SendJSON(t, p, "POST", "/v1/login", form)
	r = new(model.BadRequestResponse)
	err = json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Wrong user.
	form = url.Values{}
	form.Set("email", "b@b.com")
	form.Set("password", "a")
	w = tr.SendJSON(t, p, "POST", "/v1/login", form)
	r = new(model.BadRequestResponse)
	err = json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Login with the user.
	w = tr.SendJSON(t, p, "POST", "/v1/login", nil)
	r = new(model.BadRequestResponse)
	err = json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Invalid unmarshal.
	e := errors.New("bad error")
	m.Mock.Add("Binder.Unmarshal", e)
	w = tr.SendJSON(t, p, "POST", "/v1/login", nil)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLoginTokenBad(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	p, m := testutil.Services(db)
	tr := testrequest.New()

	p.Token = m.Token
	m.Token.GenerateFunc = func(userID string) (string, error) {
		return "", errors.New("bad token generation")
	}

	// Register the user.
	register(t, tr, p)

	// Login with the user.
	form := url.Values{}
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := tr.SendJSON(t, p, "POST", "/v1/login", form)

	// Verify the response.
	r := new(model.LoginResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "", r.Body.Token)
}

func TestLoginFailDatabase(t *testing.T) {
	p, _ := testutil.Services(nil)
	tr := testrequest.New()

	// Login with the user.
	form := url.Values{}
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := tr.SendJSON(t, p, "POST", "/v1/login", form)

	// Verify the response.
	r := new(model.InternalServerErrorResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
