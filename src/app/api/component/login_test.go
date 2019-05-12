package component_test

import (
	"app/api/component"
	"app/api/internal/testrequest"
	"app/api/internal/testutil"
	"app/api/model"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
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
	testrequest.SendForm(t, core, "POST", "/v1/register", form)

	// Login with the user.
	form = url.Values{}
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := testrequest.SendForm(t, core, "POST", "/v1/login", form)

	// Verify the response.
	r := new(model.LoginResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "OK", r.Body.Status)
	assert.Equal(t, "MDEyMzQ1Njc4OUFCQ0RFRjAxMjM0NTY3ODlBQkNERUY=", r.Body.Token)

	testutil.TeardownDatabase(db)
}
