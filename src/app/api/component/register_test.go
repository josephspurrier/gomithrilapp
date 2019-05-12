package component_test

import (
	"encoding/base64"
	"encoding/json"
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

func TestRegisterSuccess(t *testing.T) {
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
