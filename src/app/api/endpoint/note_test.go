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

	// Invalid unmarshal.
	e := errors.New("bad error")
	m.Mock.Add("Binder.Unmarshal", e)
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "POST", "/v1/note", nil)
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

func TestNoteIndexSuccess(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	p, _ := boot.TestServices(db)
	tr := testrequest.New()

	// Get an auth token.
	token := auth(t, tr, p)

	// Get the notes - there should be none.
	tr.Header.Set("Authorization", "Bearer "+token)
	w := tr.SendJSON(t, p, "GET", "/v1/note", nil)
	rr := new(model.NoteIndexResponse)
	err := json.Unmarshal(w.Body.Bytes(), &rr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 0, len(rr.Body.Notes))

	// Create a note.
	form := url.Values{}
	form.Set("message", "foo")
	tr.Header.Set("Authorization", "Bearer "+token)
	tr.SendJSON(t, p, "POST", "/v1/note", form)

	// Create another note.
	form = url.Values{}
	form.Set("message", "foo2")
	tr.Header.Set("Authorization", "Bearer "+token)
	tr.SendJSON(t, p, "POST", "/v1/note", form)

	// Get the notes.
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "GET", "/v1/note", nil)

	// Verify the response.
	rr = new(model.NoteIndexResponse)
	err = json.Unmarshal(w.Body.Bytes(), &rr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 2, len(rr.Body.Notes))
}

func TestNoteIndexFail(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	p, m := boot.TestServices(db)
	tr := testrequest.New()

	// Get an auth token.
	token := auth(t, tr, p)

	// Invalid user.
	m.Mock.Add("CTX.UserID", "", false)
	tr.Header.Set("Authorization", "Bearer "+token)
	w := tr.SendJSON(t, p, "GET", "/v1/note", nil)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Invalid DB.
	m.Mock.Add("NoteStore.FindAllByUser", 0, errors.New("no notes"))
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "GET", "/v1/note", nil)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestNoteShowSuccess(t *testing.T) {
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
	rr := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &rr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, 36, len(rr.Body.RecordID))

	recordID := rr.Body.RecordID

	// Get the note.
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "GET", "/v1/note/"+recordID, nil)
	rrr := new(model.NoteShowResponse)
	err = json.Unmarshal(w.Body.Bytes(), &rrr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "foo", rrr.Body.Message)
}

func TestNoteShowFail(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	p, m := boot.TestServices(db)
	tr := testrequest.New()

	// Get an auth token.
	token := auth(t, tr, p)

	// Create a note.
	form := url.Values{}
	form.Set("message", "foo")
	tr.Header.Set("Authorization", "Bearer "+token)
	w := tr.SendJSON(t, p, "POST", "/v1/note", form)
	rr := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &rr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, 36, len(rr.Body.RecordID))

	recordID := rr.Body.RecordID

	// Note not found.
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "GET", "/v1/note/bad-id", nil)
	rrr := new(model.BadRequestResponse)
	err = json.Unmarshal(w.Body.Bytes(), &rrr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Invalid unmarshal.
	e := errors.New("bad error")
	m.Mock.Add("Binder.Unmarshal", e)
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "GET", "/v1/note/"+recordID, nil)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Invalid validate.
	m.Mock.Add("Binder.Validate", e)
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "GET", "/v1/note/"+recordID, nil)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Invalid user.
	m.Mock.Add("CTX.UserID", "", false)
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "GET", "/v1/note/"+recordID, nil)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Invalid DB.
	m.Mock.Add("NoteStore.FindOneByIDAndUser", false, errors.New("no note"))
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "GET", "/v1/note/"+recordID, nil)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestNoteUpdateSuccess(t *testing.T) {
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
	rr := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &rr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, 36, len(rr.Body.RecordID))

	recordID := rr.Body.RecordID

	// Update the note.
	tr.Header.Set("Authorization", "Bearer "+token)
	form = url.Values{}
	form.Set("message", "bar")
	w = tr.SendJSON(t, p, "PUT", "/v1/note/"+recordID, form)
	rrr := new(model.OKResponse)
	err = json.Unmarshal(w.Body.Bytes(), &rrr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)

	// Get the note.
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "GET", "/v1/note/"+recordID, nil)
	rrrr := new(model.NoteShowResponse)
	err = json.Unmarshal(w.Body.Bytes(), &rrrr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "bar", rrrr.Body.Message)
}

func TestNoteUpdateFail(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	p, m := boot.TestServices(db)
	tr := testrequest.New()

	// Get an auth token.
	token := auth(t, tr, p)

	// Create a note.
	form := url.Values{}
	form.Set("message", "foo")
	tr.Header.Set("Authorization", "Bearer "+token)
	w := tr.SendJSON(t, p, "POST", "/v1/note", form)
	rr := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &rr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, 36, len(rr.Body.RecordID))

	recordID := rr.Body.RecordID

	// Note not found.
	form = url.Values{}
	form.Set("message", "foo")
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "PUT", "/v1/note/bad-id", form)
	rrr := new(model.BadRequestResponse)
	err = json.Unmarshal(w.Body.Bytes(), &rrr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Invalid unmarshal.
	e := errors.New("bad error")
	m.Mock.Add("Binder.Unmarshal", e)
	form = url.Values{}
	form.Set("message", "foo")
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "PUT", "/v1/note/"+recordID, form)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Invalid validate.
	m.Mock.Add("Binder.Validate", e)
	form = url.Values{}
	form.Set("message", "foo")
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "PUT", "/v1/note/"+recordID, form)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Invalid user.
	form = url.Values{}
	form.Set("message", "foo")
	m.Mock.Add("CTX.UserID", "", false)
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "PUT", "/v1/note/"+recordID, form)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Invalid DB.
	m.Mock.Add("NoteStore.FindOneByIDAndUser", false, errors.New("no note"))
	form = url.Values{}
	form.Set("message", "foo")
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "PUT", "/v1/note/"+recordID, form)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Invalid DB.
	m.Mock.Add("NoteStore.Update", 0, errors.New("no note"))
	form = url.Values{}
	form.Set("message", "foo")
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "PUT", "/v1/note/"+recordID, form)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestNoteDestroySuccess(t *testing.T) {
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
	rr := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &rr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, 36, len(rr.Body.RecordID))

	recordID := rr.Body.RecordID

	// Get the note.
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "DELETE", "/v1/note/"+recordID, nil)
	rrr := new(model.OKResponse)
	err = json.Unmarshal(w.Body.Bytes(), &rrr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestNoteDestroyFail(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	p, m := boot.TestServices(db)
	tr := testrequest.New()

	// Get an auth token.
	token := auth(t, tr, p)

	// Create a note.
	form := url.Values{}
	form.Set("message", "foo")
	tr.Header.Set("Authorization", "Bearer "+token)
	w := tr.SendJSON(t, p, "POST", "/v1/note", form)
	rr := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &rr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, 36, len(rr.Body.RecordID))

	recordID := rr.Body.RecordID

	// Note not found.
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "DELETE", "/v1/note/bad-id", nil)
	rrr := new(model.BadRequestResponse)
	err = json.Unmarshal(w.Body.Bytes(), &rrr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Invalid unmarshal.
	e := errors.New("bad error")
	m.Mock.Add("Binder.Unmarshal", e)
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "DELETE", "/v1/note/"+recordID, nil)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Invalid validate.
	m.Mock.Add("Binder.Validate", e)
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "DELETE", "/v1/note/"+recordID, nil)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Invalid user.
	m.Mock.Add("CTX.UserID", "", false)
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "DELETE", "/v1/note/"+recordID, nil)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Invalid DB.
	m.Mock.Add("NoteStore.DeleteOneByIDAndUser", 0, errors.New("no note"))
	tr.Header.Set("Authorization", "Bearer "+token)
	w = tr.SendJSON(t, p, "DELETE", "/v1/note/"+recordID, nil)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
