package store_test

import (
	"errors"
	"testing"

	"app/api/internal/testutil"

	"github.com/stretchr/testify/assert"
)

func TestNote(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	// Create a user.
	u := c.Core.Store.User
	userID, err := u.Create("first", "last", "email", "password")
	assert.NoError(t, err)

	s := c.Core.Store.Note
	ID, err := s.Create(userID, "foo")
	assert.NoError(t, err)
	assert.Equal(t, 36, len(ID))

	affected, err := s.Update(ID, userID, "bar")
	assert.NoError(t, err)
	assert.Equal(t, 1, affected)

	item := s.New()
	found, err := s.FindOneByID(&item, ID)
	assert.NoError(t, err)
	assert.Equal(t, true, found)

	ID, err = s.Create(userID, "foo2")
	assert.NoError(t, err)
	assert.Equal(t, 36, len(ID))

	group := s.NewGroup()
	total, err := s.FindAll(&group)
	assert.NoError(t, err)
	assert.Equal(t, 2, total)

	group = s.NewGroup()
	total, err = s.FindAllByUser(&group, userID)
	assert.NoError(t, err)
	assert.Equal(t, 2, total)

	item = s.New()
	exists, err := s.FindOneByIDAndUser(&item, ID, userID)
	assert.NoError(t, err)
	assert.Equal(t, true, exists)

	item = s.New()
	exists, err = s.FindOneByIDAndUser(&item, "bad-id", userID)
	assert.NoError(t, err)
	assert.Equal(t, false, exists)

	item = s.New()
	affected, err = s.DeleteOneByIDAndUser(&item, ID, userID)
	assert.NoError(t, err)
	assert.Equal(t, 1, affected)

	item = s.New()
	affected, err = s.DeleteOneByIDAndUser(&item, ID, userID)
	assert.NoError(t, err)
	assert.Equal(t, 0, affected)
}

func TestNoteMock(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	// Create a user.
	u := c.Core.Store.User
	userID, err := u.Create("first", "last", "email", "password")
	assert.NoError(t, err)

	s := c.Core.Store.Note

	e := errors.New("yes")
	c.Test.Mock.Add("NoteStore.Create", "1", e)
	ID, err := s.Create(userID, "foo")
	assert.Equal(t, e, err)
	assert.Equal(t, "1", ID)

	c.Test.Mock.Add("NoteStore.Update", 33, e)
	affected, err := s.Update(ID, userID, "bar")
	assert.Equal(t, e, err)
	assert.Equal(t, 33, affected)

	group := s.NewGroup()
	c.Test.Mock.Add("NoteStore.FindAllByUser", 99, e)
	total, err := s.FindAllByUser(&group, userID)
	assert.Equal(t, e, err)
	assert.Equal(t, 99, total)

	item := s.New()
	c.Test.Mock.Add("NoteStore.FindOneByIDAndUser", true, e)
	exists, err := s.FindOneByIDAndUser(&item, "bad ID", userID)
	assert.Equal(t, e, err)
	assert.Equal(t, true, exists)

	c.Test.Mock.Add("NoteStore.DeleteOneByIDAndUser", 25, e)
	affected, err = s.DeleteOneByIDAndUser(&item, "bad ID", userID)
	assert.Equal(t, e, err)
	assert.Equal(t, 25, affected)

}
