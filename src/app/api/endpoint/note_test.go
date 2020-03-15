package endpoint_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"app/api/model"

	"github.com/stretchr/testify/assert"
)

func TestNoteCreateSuccess(t *testing.T) {
	c := Setup()
	defer c.Teardown()

	// Get an auth token.
	token, _ := Auth(t, c.Request, c.Core)

	// Mock the context.
	//m.Context.On("SetUserID", userID).Return()
	//m.Context.On("UserID").Return(userID, true)

	// Create a note.
	form := url.Values{}
	form.Set("message", "foo")
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w := c.Request.SendJSON(t, c.Core, "POST", "/api/v1/note", form)
	r := EnsureCreated(t, w)
	assert.Equal(t, 36, len(r.Body.RecordID))
}

func TestNoteCreateFail(t *testing.T) {
	c := Setup()
	defer c.Teardown()

	// Get an auth token.
	token, _ := Auth(t, c.Request, c.Core)

	// SUCCESS: Allow no message.
	form := url.Values{}
	form.Set("message", "")
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w := c.Request.SendJSON(t, c.Core, "POST", "/api/v1/note", form)
	EnsureCreated(t, w)

	// Invalid unmarshal.
	//e := errors.New("bad error")
	//c.Test.Mock.Add("Binder.Unmarshal", e)

	// Mock the context.
	//m.Context.On("SetUserID", tr, "foo").Return()
	//m.Context.On("UserID", tr).Return("foo")

	// c.Request.Header.Set("Authorization", "Bearer "+token)
	// w = c.Request.SendJSON(t, c.Core, "POST", "/api/v1/note", nil)
	// assert.Equal(t, http.StatusBadRequest, w.Code)

	// SUCCESS: Allow no message.
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w = c.Request.SendJSON(t, c.Core, "POST", "/api/v1/note", nil)
	EnsureCreated(t, w)

	// // Invalid user.
	// c.Test.Mock.Add("CTX.UserID", "", false)
	// form = url.Values{}
	// form.Set("message", "foo")
	// c.Request.Header.Set("Authorization", "Bearer "+token)
	// w = c.Request.SendJSON(t, c.Core, "POST", "/api/v1/note", form)
	// assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Invalid DB.
	c.Test.Mock.Add("NoteStore.Create", "0", errors.New("error creating note"))
	form = url.Values{}
	form.Set("message", "foo")
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w = c.Request.SendJSON(t, c.Core, "POST", "/api/v1/note", form)
	EnsureInternalServerError(t, w)
}

func TestNoteIndexSuccess(t *testing.T) {
	c := Setup()
	defer c.Teardown()

	// Get an auth token.
	token, _ := Auth(t, c.Request, c.Core)

	// Get the notes - there should be none.
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w := c.Request.SendJSON(t, c.Core, "GET", "/api/v1/note", nil)
	rr := new(model.NoteIndexResponse)
	err := json.Unmarshal(w.Body.Bytes(), &rr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 0, len(rr.Body.Notes))

	// Create a note.
	form := url.Values{}
	form.Set("message", "foo")
	c.Request.Header.Set("Authorization", "Bearer "+token)
	c.Request.SendJSON(t, c.Core, "POST", "/api/v1/note", form)

	// Create another note.
	form = url.Values{}
	form.Set("message", "foo2")
	c.Request.Header.Set("Authorization", "Bearer "+token)
	c.Request.SendJSON(t, c.Core, "POST", "/api/v1/note", form)

	// Get the notes.
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w = c.Request.SendJSON(t, c.Core, "GET", "/api/v1/note", nil)

	// Verify the response.
	rr = new(model.NoteIndexResponse)
	err = json.Unmarshal(w.Body.Bytes(), &rr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 2, len(rr.Body.Notes))
}

func TestNoteIndexFail(t *testing.T) {
	c := Setup()
	defer c.Teardown()

	// Get an auth token.
	token, _ := Auth(t, c.Request, c.Core)

	// // Invalid user.
	// c.Test.Mock.Add("CTX.UserID", "", false)
	// c.Request.Header.Set("Authorization", "Bearer "+token)
	// w := c.Request.SendJSON(t, c.Core, "GET", "/api/v1/note", nil)
	// assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Invalid DB.
	c.Test.Mock.Add("NoteStore.FindAllByUser", 0, errors.New("no notes"))
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w := c.Request.SendJSON(t, c.Core, "GET", "/api/v1/note", nil)
	EnsureInternalServerError(t, w)
}

func TestNoteShowSuccess(t *testing.T) {
	c := Setup()
	defer c.Teardown()

	// Get an auth token.
	token, _ := Auth(t, c.Request, c.Core)

	// Create a note.
	form := url.Values{}
	form.Set("message", "foo")
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w := c.Request.SendJSON(t, c.Core, "POST", "/api/v1/note", form)
	r := EnsureCreated(t, w)
	assert.Equal(t, 36, len(r.Body.RecordID))

	recordID := r.Body.RecordID

	// Get the note.
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w = c.Request.SendJSON(t, c.Core, "GET", "/api/v1/note/"+recordID, nil)
	rr := new(model.NoteShowResponse)
	err := json.Unmarshal(w.Body.Bytes(), &rr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "foo", rr.Body.Message)
}

func TestNoteShowFail(t *testing.T) {
	c := Setup()
	defer c.Teardown()

	// Get an auth token.
	token, _ := Auth(t, c.Request, c.Core)

	// Create a note.
	form := url.Values{}
	form.Set("message", "foo")
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w := c.Request.SendJSON(t, c.Core, "POST", "/api/v1/note", form)
	rr := new(model.CreatedResponse)
	err := json.Unmarshal(w.Body.Bytes(), &rr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, 36, len(rr.Body.RecordID))

	recordID := rr.Body.RecordID

	// Note not found.
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w = c.Request.SendJSON(t, c.Core, "GET", "/api/v1/note/bad-id", nil)
	EnsureBadRequest(t, w)

	// // Invalid unmarshal.
	// e := errors.New("bad error")
	// c.Test.Mock.Add("Binder.Unmarshal", e)
	// c.Request.Header.Set("Authorization", "Bearer "+token)
	// w = c.Request.SendJSON(t, c.Core, "GET", "/api/v1/note/"+recordID, nil)
	// assert.Equal(t, http.StatusBadRequest, w.Code)

	// // Invalid validate.
	// c.Test.Mock.Add("Binder.Validate", e)
	// c.Request.Header.Set("Authorization", "Bearer "+token)
	// w = c.Request.SendJSON(t, c.Core, "GET", "/api/v1/note/"+recordID, nil)
	// assert.Equal(t, http.StatusBadRequest, w.Code)

	// // Invalid user.
	// c.Test.Mock.Add("CTX.UserID", "", false)
	// c.Request.Header.Set("Authorization", "Bearer "+token)
	// w = c.Request.SendJSON(t, c.Core, "GET", "/api/v1/note/"+recordID, nil)
	// assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Invalid DB.
	c.Test.Mock.Add("NoteStore.FindOneByIDAndUser", false, errors.New("no note"))
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w = c.Request.SendJSON(t, c.Core, "GET", "/api/v1/note/"+recordID, nil)
	EnsureInternalServerError(t, w)
}

func TestNoteUpdateSuccess(t *testing.T) {
	c := Setup()
	defer c.Teardown()

	// Get an auth token.
	token, _ := Auth(t, c.Request, c.Core)

	// Create a note.
	form := url.Values{}
	form.Set("message", "foo")
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w := c.Request.SendJSON(t, c.Core, "POST", "/api/v1/note", form)
	rr := EnsureCreated(t, w)
	assert.Equal(t, 36, len(rr.Body.RecordID))

	recordID := rr.Body.RecordID

	// Update the note.
	c.Request.Header.Set("Authorization", "Bearer "+token)
	form = url.Values{}
	form.Set("message", "bar")
	w = c.Request.SendJSON(t, c.Core, "PUT", "/api/v1/note/"+recordID, form)
	EnsureOK(t, w)

	// Get the note.
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w = c.Request.SendJSON(t, c.Core, "GET", "/api/v1/note/"+recordID, nil)
	rrrr := new(model.NoteShowResponse)
	err := json.Unmarshal(w.Body.Bytes(), &rrrr.Body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "bar", rrrr.Body.Message)
}

func TestNoteUpdateFail(t *testing.T) {
	c := Setup()
	defer c.Teardown()

	// Get an auth token.
	token, _ := Auth(t, c.Request, c.Core)

	// Create a note.
	form := url.Values{}
	form.Set("message", "foo")
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w := c.Request.SendJSON(t, c.Core, "POST", "/api/v1/note", form)
	rr := EnsureCreated(t, w)
	assert.Equal(t, 36, len(rr.Body.RecordID))

	recordID := rr.Body.RecordID

	// Note not found.
	form = url.Values{}
	form.Set("message", "foo")
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w = c.Request.SendJSON(t, c.Core, "PUT", "/api/v1/note/bad-id", form)
	EnsureBadRequest(t, w)

	// // Invalid unmarshal.
	// e := errors.New("bad error")
	// c.Test.Mock.Add("Binder.Unmarshal", e)
	// form = url.Values{}
	// form.Set("message", "foo")
	// c.Request.Header.Set("Authorization", "Bearer "+token)
	// w = c.Request.SendJSON(t, c.Core, "PUT", "/api/v1/note/"+recordID, form)
	// assert.Equal(t, http.StatusBadRequest, w.Code)

	// // Invalid validate.
	// c.Test.Mock.Add("Binder.Validate", e)
	// form = url.Values{}
	// form.Set("message", "foo")
	// c.Request.Header.Set("Authorization", "Bearer "+token)
	// w = c.Request.SendJSON(t, c.Core, "PUT", "/api/v1/note/"+recordID, form)
	// assert.Equal(t, http.StatusBadRequest, w.Code)

	// // Invalid user.
	// form = url.Values{}
	// form.Set("message", "foo")
	// c.Test.Mock.Add("CTX.UserID", "", false)
	// c.Request.Header.Set("Authorization", "Bearer "+token)
	// w = c.Request.SendJSON(t, c.Core, "PUT", "/api/v1/note/"+recordID, form)
	// assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Invalid DB.
	c.Test.Mock.Add("NoteStore.FindOneByIDAndUser", false, errors.New("no note"))
	form = url.Values{}
	form.Set("message", "foo")
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w = c.Request.SendJSON(t, c.Core, "PUT", "/api/v1/note/"+recordID, form)
	EnsureInternalServerError(t, w)

	// Invalid DB.
	c.Test.Mock.Add("NoteStore.Update", 0, errors.New("no note"))
	form = url.Values{}
	form.Set("message", "foo")
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w = c.Request.SendJSON(t, c.Core, "PUT", "/api/v1/note/"+recordID, form)
	EnsureInternalServerError(t, w)
}

func TestNoteDestroySuccess(t *testing.T) {
	c := Setup()
	defer c.Teardown()

	// Get an auth token.
	token, _ := Auth(t, c.Request, c.Core)

	// Create a note.
	form := url.Values{}
	form.Set("message", "foo")
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w := c.Request.SendJSON(t, c.Core, "POST", "/api/v1/note", form)
	rr := EnsureCreated(t, w)
	assert.Equal(t, 36, len(rr.Body.RecordID))

	recordID := rr.Body.RecordID

	// Get the note.
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w = c.Request.SendJSON(t, c.Core, "DELETE", "/api/v1/note/"+recordID, nil)
	EnsureOK(t, w)
}

func TestNoteDestroyFail(t *testing.T) {
	c := Setup()
	defer c.Teardown()

	// Get an auth token.
	token, _ := Auth(t, c.Request, c.Core)

	// Create a note.
	form := url.Values{}
	form.Set("message", "foo")
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w := c.Request.SendJSON(t, c.Core, "POST", "/api/v1/note", form)
	rr := EnsureCreated(t, w)
	assert.Equal(t, 36, len(rr.Body.RecordID))

	recordID := rr.Body.RecordID

	// Note not found.
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w = c.Request.SendJSON(t, c.Core, "DELETE", "/api/v1/note/bad-id", nil)
	EnsureBadRequest(t, w)

	// // Invalid unmarshal.
	// e := errors.New("bad error")
	// c.Test.Mock.Add("Binder.Unmarshal", e)
	// c.Request.Header.Set("Authorization", "Bearer "+token)
	// w = c.Request.SendJSON(t, c.Core, "DELETE", "/api/v1/note/"+recordID, nil)
	// assert.Equal(t, http.StatusBadRequest, w.Code)

	// // Invalid validate.
	// c.Test.Mock.Add("Binder.Validate", e)
	// c.Request.Header.Set("Authorization", "Bearer "+token)
	// w = c.Request.SendJSON(t, c.Core, "DELETE", "/api/v1/note/"+recordID, nil)
	// assert.Equal(t, http.StatusBadRequest, w.Code)

	// // Invalid user.
	// c.Test.Mock.Add("CTX.UserID", "", false)
	// c.Request.Header.Set("Authorization", "Bearer "+token)
	// w = c.Request.SendJSON(t, c.Core, "DELETE", "/api/v1/note/"+recordID, nil)
	// assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Invalid DB.
	c.Test.Mock.Add("NoteStore.DeleteOneByIDAndUser", 0, errors.New("no note"))
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w = c.Request.SendJSON(t, c.Core, "DELETE", "/api/v1/note/"+recordID, nil)
	EnsureInternalServerError(t, w)
}
