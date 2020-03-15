package endpoint_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"app/api/internal/testutil"

	"github.com/stretchr/testify/assert"
)

func TestNote(t *testing.T) {
	c := testutil.Setup()

	// Get an auth token.
	token, _ := testutil.Auth(t, c.Request, c.Core)

	tests := []struct {
		Method string
		Path   string
	}{
		{"POST", "/api/v1/note"},
		{"GET", "/api/v1/note"},
		{"GET", "/api/v1/note/"},
		{"PUT", "/api/v1/note/"},
		{"DELETE", "/api/v1/note/"},
	}

	noteID := ""
	limit := 2
	for i := 0; i < limit; i++ {
		if i > 0 {
			c.Teardown()
		}
		for _, tt := range tests {
			t.Run(fmt.Sprint(i)+tt.Method+tt.Path, func(t *testing.T) {
				c.Request.Header.Set("Authorization", "Bearer "+token)
				add := ""
				if strings.HasSuffix(tt.Path, "/") {
					add = noteID
				}
				w := c.Request.SendJSON(t, c.Core, tt.Method, tt.Path+add, nil)

				if i > 0 {
					testutil.EnsureInternalServerError(t, w)
				} else {
					// Create a note.
					if tt.Method == "POST" {
						r := testutil.EnsureCreated(t, w)
						noteID = r.Body.RecordID
						assert.Equal(t, 36, len(noteID))
					} else {
						testutil.EnsureOK(t, w)
					}
				}
			})
		}
	}
}

func TestNoteUnauthorized(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	tests := []struct {
		Method string
		Path   string
	}{
		{"POST", "/api/v1/note"},
		{"GET", "/api/v1/note"},
		{"GET", "/api/v1/note/1"},
		{"PUT", "/api/v1/note/1"},
		{"DELETE", "/api/v1/note/1"},
	}

	for _, tt := range tests {
		t.Run(tt.Method+tt.Path, func(t *testing.T) {
			w := c.Request.SendJSON(t, c.Core, tt.Method, tt.Path, nil)
			testutil.EnsureUnauthorized(t, w)
		})
	}
}

func TestNoteFailContext(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	// Get an auth token.
	token, userID := testutil.Auth(t, c.Request, c.Core)

	ctx := new(testutil.Context)
	ctx.On("SetUserID", userID).Return()
	ctx.On("UserID").Return("2", true)
	c.Core.Context = ctx

	tests := []struct {
		Method string
		Path   string
	}{
		{"POST", "/api/v1/note"},
		{"GET", "/api/v1/note"},
		{"GET", "/api/v1/note/1"},
		{"PUT", "/api/v1/note/1"},
		{"DELETE", "/api/v1/note/1"},
	}

	for _, tt := range tests {
		t.Run(tt.Method+tt.Path, func(t *testing.T) {
			c.Request.Header.Set("Authorization", "Bearer "+token)
			w := c.Request.SendJSON(t, c.Core, tt.Method, tt.Path, nil)
			testutil.EnsureInternalServerError(t, w)
		})
	}
}

func TestNoteFailDatabase(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	// Get an auth token.
	token, _ := testutil.Auth(t, c.Request, c.Core)

	// Create a note.
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w := c.Request.SendJSON(t, c.Core, "POST", "/api/v1/note", nil)
	r := testutil.EnsureCreated(t, w)

	c.Test.Mock.Add("NoteStore.FindOneByIDAndUser", false, nil)
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w = c.Request.SendJSON(t, c.Core, "GET", "/api/v1/note/"+r.Body.RecordID, nil)
	testutil.EnsureBadRequest(t, w)

	c.Test.Mock.Add("NoteStore.FindOneByIDAndUser", false, nil)
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w = c.Request.SendJSON(t, c.Core, "PUT", "/api/v1/note/"+r.Body.RecordID, nil)
	testutil.EnsureBadRequest(t, w)

	err := errors.New("bad error")
	c.Test.Mock.Add("NoteStore.Update", 0, err)
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w = c.Request.SendJSON(t, c.Core, "PUT", "/api/v1/note/"+r.Body.RecordID, nil)
	testutil.EnsureInternalServerError(t, w)

	c.Test.Mock.Add("NoteStore.DeleteOneByIDAndUser", 0, nil)
	c.Request.Header.Set("Authorization", "Bearer "+token)
	w = c.Request.SendJSON(t, c.Core, "DELETE", "/api/v1/note/"+r.Body.RecordID, nil)
	testutil.EnsureBadRequest(t, w)
}
