package component_test

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"testing"
	"time"

	"app/api/component"
	"app/api/internal/testrequest"
	"app/api/internal/testutil"
	"app/api/model"

	"github.com/stretchr/testify/assert"
)

func TestLoginSuccess(t *testing.T) {
	db := testutil.LoadDatabase()
	core, m := component.NewCoreMock(db)

	m.Token.GenerateFunc = func(userID string, duration time.Duration) (string, error) {
		b := []byte("0123456789ABCDEF0123456789ABCDEF")
		enc := base64.StdEncoding.EncodeToString(b)
		return enc, nil
	}

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	testrequest.SendJSON(t, core, "POST", "/v1/register", form)

	// Login with the user.
	form = url.Values{}
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := testrequest.SendJSON(t, core, "POST", "/v1/login", form)

	// Verify the response.
	r := new(model.LoginResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "MDEyMzQ1Njc4OUFCQ0RFRjAxMjM0NTY3ODlBQkNERUY=", r.Body.Token)

	testutil.TeardownDatabase(db)
}

func TestLoginFail(t *testing.T) {
	db := testutil.LoadDatabase()
	core, m := component.NewCoreMock(db)

	m.Token.GenerateFunc = func(userID string, duration time.Duration) (string, error) {
		b := []byte("0123456789ABCDEF0123456789ABCDEF")
		enc := base64.StdEncoding.EncodeToString(b)
		return enc, nil
	}

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	testrequest.SendJSON(t, core, "POST", "/v1/register", form)

	// Login with the user.
	form = url.Values{}
	form.Set("email", "a@a.com")
	form.Set("password", "wrong-password")
	w := testrequest.SendJSON(t, core, "POST", "/v1/login", form)

	// Verify the response.
	r := new(model.LoginResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", r.Body.Token)

	testutil.TeardownDatabase(db)
}

func TestLoginFailMissingField(t *testing.T) {
	db := testutil.LoadDatabase()
	core, m := component.NewCoreMock(db)

	m.Token.GenerateFunc = func(userID string, duration time.Duration) (string, error) {
		b := []byte("0123456789ABCDEF0123456789ABCDEF")
		enc := base64.StdEncoding.EncodeToString(b)
		return enc, nil
	}

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	testrequest.SendJSON(t, core, "POST", "/v1/register", form)

	// Login with the user.
	form = url.Values{}
	form.Set("email", "a@a.com")
	//form.Set("password", "wrong-password")
	w := testrequest.SendJSON(t, core, "POST", "/v1/login", form)

	// Verify the response.
	r := new(model.LoginResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", r.Body.Token)

	testutil.TeardownDatabase(db)
}

func TestLoginFailMissingUser(t *testing.T) {
	db := testutil.LoadDatabase()
	core, m := component.NewCoreMock(db)

	m.Token.GenerateFunc = func(userID string, duration time.Duration) (string, error) {
		b := []byte("0123456789ABCDEF0123456789ABCDEF")
		enc := base64.StdEncoding.EncodeToString(b)
		return enc, nil
	}

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	testrequest.SendJSON(t, core, "POST", "/v1/register", form)

	// Login with the user.
	form = url.Values{}
	form.Set("email", "b@b.com")
	form.Set("password", "a")
	w := testrequest.SendJSON(t, core, "POST", "/v1/login", form)

	// Verify the response.
	r := new(model.LoginResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", r.Body.Token)

	testutil.TeardownDatabase(db)
}

func TestLoginFailMissingBody(t *testing.T) {
	db := testutil.LoadDatabase()
	core, m := component.NewCoreMock(db)

	m.Token.GenerateFunc = func(userID string, duration time.Duration) (string, error) {
		b := []byte("0123456789ABCDEF0123456789ABCDEF")
		enc := base64.StdEncoding.EncodeToString(b)
		return enc, nil
	}

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	testrequest.SendJSON(t, core, "POST", "/v1/register", form)

	// Login with the user.
	w := testrequest.SendJSON(t, core, "POST", "/v1/login", nil)

	// Verify the response.
	r := new(model.LoginResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", r.Body.Token)

	testutil.TeardownDatabase(db)
}

func TestLoginToken(t *testing.T) {
	db := testutil.LoadDatabase()
	core, m := component.NewCoreMock(db)

	m.Token.GenerateFunc = func(userID string, duration time.Duration) (string, error) {
		return "", errors.New("bad token generation")
	}

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	testrequest.SendJSON(t, core, "POST", "/v1/register", form)

	// Login with the user.
	form = url.Values{}
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := testrequest.SendJSON(t, core, "POST", "/v1/login", form)

	// Verify the response.
	r := new(model.LoginResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "", r.Body.Token)

	testutil.TeardownDatabase(db)
}
