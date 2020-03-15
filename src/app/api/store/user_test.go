package store_test

import (
	"errors"
	"testing"

	"app/api/internal/testutil"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	s := c.Core.Store.User

	ID, err := s.Create("a", "b", "c", "d")
	assert.NoError(t, err)
	assert.Equal(t, 36, len(ID))

	affected, err := s.Update(ID, "aa", "bb", "cc", "dd")
	assert.NoError(t, err)
	assert.Equal(t, 1, affected)

	user := s.New()
	found, err := s.FindOneByID(&user, ID)
	assert.NoError(t, err)
	assert.Equal(t, true, found)

	ID, err = s.Create("aaa", "bbb", "ccc", "ddd")
	assert.NoError(t, err)
	assert.Equal(t, 36, len(ID))

	group := s.NewGroup()
	total, err := s.FindAll(&group)
	assert.NoError(t, err)
	assert.Equal(t, 2, total)
}

func TestUserMock(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	s := c.Core.Store.User

	e := errors.New("yes")
	c.Test.Mock.Add("UserStore.Create", "1", e)
	ID, err := s.Create("aaa", "bbb", "ccc", "ddd")
	assert.Equal(t, e, err)
	assert.Equal(t, "1", ID)

	c.Test.Mock.Add("UserStore.Update", 22, e)
	affected, err := s.Update(ID, "aa", "bb", "cc", "dd")
	assert.Equal(t, e, err)
	assert.Equal(t, 22, affected)
}
