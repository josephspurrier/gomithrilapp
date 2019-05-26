package endpoint_test

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	"app/api/boot"
	"app/api/internal/testrequest"
	"app/api/internal/testutil"
	"app/api/model"

	"github.com/stretchr/testify/assert"
)

func TestNoteCreate(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	p, _ := boot.TestServices(db)
	tr := testrequest.New()

	// Register the user.
	form := url.Values{}
	form.Set("first_name", "a")
	form.Set("last_name", "a")
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	tr.SendJSON(t, p, "POST", "/v1/register", form)

	// Login with the user.
	form = url.Values{}
	form.Set("email", "a@a.com")
	form.Set("password", "a")
	w := tr.SendJSON(t, p, "POST", "/v1/login", form)

	r := new(model.LoginResponse)
	err := json.Unmarshal(w.Body.Bytes(), &r.Body)
	assert.Nil(t, err)

	// Create a note.
	form = url.Values{}
	form.Set("message", "foo")
	tr.Header.Set("Authorization", "Bearer "+r.Body.Token)
	w = tr.SendJSON(t, p, "POST", "/v1/note", form)

	// Verify the response.
	rr := new(model.CreatedResponse)
	err = json.Unmarshal(w.Body.Bytes(), &rr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, 36, len(rr.Body.RecordID))
}
