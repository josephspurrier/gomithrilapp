package endpoint_test

import (
	"errors"
	"net/url"
	"testing"

	"app/api/internal/testutil"

	"github.com/stretchr/testify/assert"
)

func TestRegisterSuccess(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "Foo")
	form.Set("last_name", "Bar")
	form.Set("email", "fbar@example.com")
	form.Set("password", "password")
	w := c.Request.SendJSON(t, c.Core, "POST", "/api/v1/register", form)
	r := testutil.EnsureCreated(t, w)
	assert.Equal(t, 36, len(r.Body.RecordID))
}

func TestRegisterFailUserExists(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	// Register the user.
	testutil.Register(t, c.Request, c.Core)

	// Try to register the same user.
	form := url.Values{}
	form.Set("first_name", "Foo")
	form.Set("last_name", "Bar")
	form.Set("email", "fbar@example.com")
	form.Set("password", "guess123")
	w := c.Request.SendJSON(t, c.Core, "POST", "/api/v1/register", form)
	testutil.EnsureBadRequest(t, w)
}

func TestRegisterFailInvalidRequest(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	// Missing password.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "b@a.com")
	//form.Set("password", "a")
	w := c.Request.SendJSON(t, c.Core, "POST", "/api/v1/register", form)
	testutil.EnsureBadRequest(t, w)

	// Invalid unmarshal.
	w = c.Request.SendJSON(t, c.Core, "POST", "/api/v1/register", nil)
	testutil.EnsureBadRequest(t, w)
}

func TestRegisterFailDatabase(t *testing.T) {
	c := testutil.Setup()
	c.Teardown() // Teardown now to test a bad DB connection.

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := c.Request.SendJSON(t, c.Core, "POST", "/api/v1/register", form)
	testutil.EnsureInternalServerError(t, w)
}

func TestRegisterFailHash(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	// Mock the password hash library.
	mpass := new(testutil.MockPasshash)
	c.Core.Password = mpass
	mpass.HashError = errors.New("bad error")

	// Fail hash.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := c.Request.SendJSON(t, c.Core, "POST", "/api/v1/register", form)
	testutil.EnsureInternalServerError(t, w)
}

func TestRegisterFailCreateUser(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	// Force the operation to fail.
	c.Test.Mock.Add("UserStore.Create", "0", errors.New("error creating user"))

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := c.Request.SendJSON(t, c.Core, "POST", "/api/v1/register", form)
	testutil.EnsureInternalServerError(t, w)
}
