package store_test

import (
	"errors"
	"testing"

	"app/api/component"
	"app/api/internal/testutil"
	"app/api/store"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	core, _ := component.NewCoreMock(db)

	user := store.NewUser(core.Mock, core.DB, core.Q)
	ID, err := user.Create("a", "b", "c", "d")
	assert.NoError(t, err)
	assert.Equal(t, 36, len(ID))

	err = user.Update(ID, "aa", "bb", "cc", "dd")
	assert.NoError(t, err)

	found, err := user.FindOneByID(user, ID)
	assert.NoError(t, err)
	assert.Equal(t, true, found)

	ID, err = user.Create("aaa", "bbb", "ccc", "ddd")
	assert.NoError(t, err)
	assert.Equal(t, 36, len(ID))

	group := user.NewGroup()
	total, err := user.FindAll(group)
	assert.NoError(t, err)
	assert.Equal(t, 2, total)
}

func TestUserMock(t *testing.T) {
	db := testutil.LoadDatabase()
	defer testutil.TeardownDatabase(db)
	core, m := component.NewCoreMock(db)

	user := store.NewUser(core.Mock, core.DB, core.Q)

	e := errors.New("yes")
	m.Mock.Add("User.Create", "1", e)
	ID, err := user.Create("aaa", "bbb", "ccc", "ddd")
	assert.Equal(t, e, err)
	assert.Equal(t, "1", ID)

	m.Mock.Add("User.Update", e)
	err = user.Update(ID, "aa", "bb", "cc", "dd")
	assert.Equal(t, e, err)
}
