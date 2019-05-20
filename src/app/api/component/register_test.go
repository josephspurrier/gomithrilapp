package component_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"app/api/boot"
	"app/api/internal/testrequest"
	"app/api/internal/testutil"
	"app/api/model"

	"github.com/stretchr/testify/assert"
)

func TestRegisterSuccess(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	p, _ := boot.TestServices(db)

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := testrequest.SendJSON(t, p, "POST", "/v1/register", form)

	// Verify the response.
	r := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, 36, len(r.Body.RecordID))
}

func TestRegisterFailUserExists(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	p, _ := boot.TestServices(db)

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	testrequest.SendJSON(t, p, "POST", "/v1/register", form)

	form = url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := testrequest.SendJSON(t, p, "POST", "/v1/register", form)

	// Verify the response.
	r := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestRegisterFailMissingField(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	p, _ := boot.TestServices(db)

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	//form.Set("password", "a")
	w := testrequest.SendJSON(t, p, "POST", "/v1/register", form)

	// Verify the response.
	r := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestRegisterFailInvalidJSON(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	p, _ := boot.TestServices(db)

	// Register the user.
	w := testrequest.SendJSON(t, p, "POST", "/v1/register", nil)

	// Verify the response.
	r := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestRegisterFailDatabase(t *testing.T) {
	p, _ := boot.TestServices(nil)

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := testrequest.SendJSON(t, p, "POST", "/v1/register", form)

	// Verify the response.
	r := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestRegisterFailDatabase2(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	p, m := boot.TestServices(db)

	m.Mock.Add("UserStore.Create", "0", errors.New("error creating user"))

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := testrequest.SendJSON(t, p, "POST", "/v1/register", form)

	// Verify the response.
	r := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestRegisterFailHash(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	p, _ := boot.TestServices(db)

	mpass := new(testutil.MockPasshash)
	p.Password = mpass
	mpass.HashError = errors.New("bad error")

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := testrequest.SendJSON(t, p, "POST", "/v1/register", form)

	// Verify the response.
	r := new(model.InternalServerErrorResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
