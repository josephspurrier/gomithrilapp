package component_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"app/api/component"
	"app/api/internal/testrequest"
	"app/api/internal/testutil"
	"app/api/model"

	"github.com/stretchr/testify/assert"
)

func TestRegisterSuccess(t *testing.T) {
	db := testutil.LoadDatabase()
	core, _ := component.NewCoreMock(db)

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := testrequest.SendJSON(t, core, "POST", "/v1/register", form)

	// Verify the response.
	r := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, 36, len(r.Body.RecordID))

	testutil.TeardownDatabase(db)
}

func TestRegisterFailUserExists(t *testing.T) {
	db := testutil.LoadDatabase()
	core, _ := component.NewCoreMock(db)

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	testrequest.SendJSON(t, core, "POST", "/v1/register", form)

	form = url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := testrequest.SendJSON(t, core, "POST", "/v1/register", form)

	// Verify the response.
	r := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	testutil.TeardownDatabase(db)
}

func TestRegisterFailMissingField(t *testing.T) {
	db := testutil.LoadDatabase()
	core, _ := component.NewCoreMock(db)

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	//form.Set("password", "a")
	w := testrequest.SendJSON(t, core, "POST", "/v1/register", form)

	// Verify the response.
	r := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	testutil.TeardownDatabase(db)
}

func TestRegisterFailInvalidJSON(t *testing.T) {
	db := testutil.LoadDatabase()
	core, _ := component.NewCoreMock(db)

	// Register the user.
	w := testrequest.SendJSON(t, core, "POST", "/v1/register", nil)

	// Verify the response.
	r := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	testutil.TeardownDatabase(db)
}

func TestRegisterFailDatabase(t *testing.T) {
	core, _ := component.NewCoreMock(nil)

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := testrequest.SendJSON(t, core, "POST", "/v1/register", form)

	// Verify the response.
	r := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestRegisterFailDatabase2(t *testing.T) {
	db := testutil.LoadDatabase()
	core, _ := component.NewCoreMock(db)

	core.Mock.Add("UserStore.Create", "0", errors.New("error creating user"))

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := testrequest.SendJSON(t, core, "POST", "/v1/register", form)

	// Verify the response.
	r := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	testutil.TeardownDatabase(db)
}

func TestRegisterFailHash(t *testing.T) {
	db := testutil.LoadDatabase()
	core, _ := component.NewCoreMock(db)

	mpass := new(testutil.MockPasshash)
	core.Password = mpass
	mpass.HashError = errors.New("bad error")

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a@a.com")
	form.Set("last_name", "a@a.com")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := testrequest.SendJSON(t, core, "POST", "/v1/register", form)

	// Verify the response.
	r := new(model.InternalServerErrorResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	testutil.TeardownDatabase(db)
}
