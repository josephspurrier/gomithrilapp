package endpoint_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"app/api/boot"
	"app/api/endpoint"
	"app/api/internal/testrequest"
	"app/api/internal/testutil"
	"app/api/model"

	"github.com/stretchr/testify/assert"
)

func auth(t *testing.T, tr *testrequest.TR, p endpoint.Core) string {
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

	return r.Body.Token
}

func TestNoteCreateSuccess(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	p, _ := boot.TestServices(db)
	tr := testrequest.New()

	// Get an auth token.
	token := auth(t, tr, p)

	// Create a note.
	form := url.Values{}
	form.Set("message", "foo")
	tr.Header.Set("Authorization", "Bearer "+token)
	w := tr.SendJSON(t, p, "POST", "/v1/note", form)

	// Verify the response.
	rr := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &rr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, 36, len(rr.Body.RecordID))
}

func TestNoteCreateFail(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	p, m := boot.TestServices(db)
	tr := testrequest.New()

	// Get an auth token.
	token := auth(t, tr, p)

	// Missing message.
	form := url.Values{}
	form.Set("message", "")
	tr.Header.Set("Authorization", "Bearer "+token)
	w := tr.SendJSON(t, p, "POST", "/v1/note", form)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Invalid data.
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "POST", "/v1/note", nil)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Invalid user.
	m.Mock.Add("CTX.UserID", "", false)
	form = url.Values{}
	form.Set("message", "foo")
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "POST", "/v1/note", form)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Invalid DB.
	m.Mock.Add("NoteStore.Create", "0", errors.New("error creating note"))
	form = url.Values{}
	form.Set("message", "foo")
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "POST", "/v1/note", form)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
