package endpoint_test

import (
	"encoding/base64"
	"errors"
	"net/url"
	"testing"

	"app/api/internal/testutil"

	"github.com/stretchr/testify/assert"
)

func TestLoginSuccess(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	// Mock the token service.
	mt := testutil.NewMockToken()
	mt.GenerateFunc = func(userID string) (string, error) {
		enc := base64.StdEncoding.EncodeToString([]byte("0123456789ABCDEF0123456789ABCDEF"))
		return enc, nil
	}
	c.Core.Token = mt

	// Register and login with the user.
	userToken, _ := testutil.Auth(t, c.Request, c.Core)

	// Verify the response.
	assert.Equal(t, "MDEyMzQ1Njc4OUFCQ0RFRjAxMjM0NTY3ODlBQkNERUY=", userToken)
}

func TestLoginFail(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	// Register the user.
	testutil.Register(t, c.Request, c.Core)

	// Wrong password.
	form := url.Values{}
	form.Set("email", "fbar@example.com")
	form.Set("password", "wrong-password")
	w := c.Request.SendJSON(t, c.Core, "POST", "/api/v1/login", form)
	testutil.EnsureBadRequest(t, w)

	// Missing password.
	form = url.Values{}
	form.Set("email", "fbar@example.com")
	w = c.Request.SendJSON(t, c.Core, "POST", "/api/v1/login", form)
	testutil.EnsureBadRequest(t, w)

	// Wrong user.
	form = url.Values{}
	form.Set("email", "other@example.com")
	form.Set("password", "guess123")
	w = c.Request.SendJSON(t, c.Core, "POST", "/api/v1/login", form)
	testutil.EnsureBadRequest(t, w)

	// Invalid unmarshal.
	w = c.Request.SendJSON(t, c.Core, "POST", "/api/v1/login", nil)
	testutil.EnsureBadRequest(t, w)
}

func TestLoginTokenBad(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	// Mock the token service.
	mt := testutil.NewMockToken()
	mt.GenerateFunc = func(userID string) (string, error) {
		return "", errors.New("bad token generation")
	}
	c.Core.Token = mt

	// Register the user.
	testutil.Register(t, c.Request, c.Core)

	// Login with the user.
	form := url.Values{}
	form.Set("email", "fbar@example.com")
	form.Set("password", "guess123")
	w := c.Request.SendJSON(t, c.Core, "POST", "/api/v1/login", form)
	testutil.EnsureInternalServerError(t, w)
}

func TestLoginFailDatabase(t *testing.T) {
	c := testutil.Setup()
	c.Teardown() // Teardown now to test a bad DB connection.

	// Login with the user.
	form := url.Values{}
	form.Set("email", "fbar@example.com")
	form.Set("password", "guess123")
	w := c.Request.SendJSON(t, c.Core, "POST", "/api/v1/login", form)
	testutil.EnsureInternalServerError(t, w)
}
