package store_test

import (
	"errors"
	"testing"

	"app/api/component"
	"app/api/internal/testutil"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	p, _ := component.NewCoreMock(db)

	s := p.Store.User

	ID, err := s.Create("a", "b", "c", "d")
	assert.NoError(t, err)
	assert.Equal(t, 36, len(ID))

	err = s.Update(ID, "aa", "bb", "cc", "dd")
	assert.NoError(t, err)

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
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	p, _ := component.NewCoreMock(db)

	s := p.Store.User

	e := errors.New("yes")
	p.Mock.Add("UserStore.Create", "1", e)
	ID, err := s.Create("aaa", "bbb", "ccc", "ddd")
	assert.Equal(t, e, err)
	assert.Equal(t, "1", ID)

	p.Mock.Add("UserStore.Update", e)
	err = s.Update(ID, "aa", "bb", "cc", "dd")
	assert.Equal(t, e, err)
}
