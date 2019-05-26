package endpoint_test

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"testing"
	"time"

	"app/api/boot"
	"app/api/internal/testrequest"
	"app/api/internal/testutil"
	"app/api/model"

	"github.com/stretchr/testify/assert"
)

func TestLoginSuccess(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	core, m := boot.TestServices(db)
	tr := testrequest.New()

	core.Token = m.Token
	m.Token.SecretValue = "0123456789ABCDEF0123456789ABCDEF"
	m.Token.GenerateFunc = func(userID string, duration time.Duration) (string, error) {
		enc := base64.StdEncoding.EncodeToString(m.Token.Secret())
		return enc, nil
	}

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	tr.SendJSON(t, core, "POST", "/v1/register", form)

	// Login with the user.
	form = url.Values{}
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := tr.SendJSON(t, core, "POST", "/v1/login", form)

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
	core, _ := boot.TestServices(db)
	tr := testrequest.New()

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	tr.SendJSON(t, core, "POST", "/v1/register", form)

	// Login with the user.
	form = url.Values{}
	form.Set("email", "a@a.com")
	form.Set("password", "wrong-password")
	w := tr.SendJSON(t, core, "POST", "/v1/login", form)

	// Verify the response.
	r := new(model.LoginResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", r.Body.Token)
}

func TestLoginFailMissingField(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	core, _ := boot.TestServices(db)
	tr := testrequest.New()

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	tr.SendJSON(t, core, "POST", "/v1/register", form)

	// Login with the user.
	form = url.Values{}
	form.Set("email", "a@a.com")
	//form.Set("password", "wrong-password")
	w := tr.SendJSON(t, core, "POST", "/v1/login", form)

	// Verify the response.
	r := new(model.LoginResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", r.Body.Token)
}

func TestLoginFailMissingUser(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	core, _ := boot.TestServices(db)
	tr := testrequest.New()

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	tr.SendJSON(t, core, "POST", "/v1/register", form)

	// Login with the user.
	form = url.Values{}
	form.Set("email", "b@b.com")
	form.Set("password", "a")
	w := tr.SendJSON(t, core, "POST", "/v1/login", form)

	// Verify the response.
	r := new(model.LoginResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", r.Body.Token)
}

func TestLoginFailMissingBody(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	core, _ := boot.TestServices(db)
	tr := testrequest.New()

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	tr.SendJSON(t, core, "POST", "/v1/register", form)

	// Login with the user.
	w := tr.SendJSON(t, core, "POST", "/v1/login", nil)

	// Verify the response.
	r := new(model.LoginResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "", r.Body.Token)
}

func TestLoginTokenBad(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	core, m := boot.TestServices(db)
	tr := testrequest.New()

	core.Token = m.Token
	m.Token.SecretValue = "0123456789ABCDEF0123456789ABCDEF"
	m.Token.GenerateFunc = func(userID string, duration time.Duration) (string, error) {
		return "", errors.New("bad token generation")
	}

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	tr.SendJSON(t, core, "POST", "/v1/register", form)

	// Login with the user.
	form = url.Values{}
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := tr.SendJSON(t, core, "POST", "/v1/login", form)

	// Verify the response.
	r := new(model.LoginResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "", r.Body.Token)
}

func TestLoginFailDatabase(t *testing.T) {
	core, _ := boot.TestServices(nil)
	tr := testrequest.New()

	// Login with the user.
	form := url.Values{}
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := tr.SendJSON(t, core, "POST", "/v1/login", form)

	// Verify the response.
	r := new(model.InternalServerErrorResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
