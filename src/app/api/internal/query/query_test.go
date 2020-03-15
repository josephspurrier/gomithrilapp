package query_test

import (
	"errors"
	"testing"

	"app/api/internal/query"
	"app/api/internal/query/internal/store"
	"app/api/internal/testutil"
	"app/api/pkg/mock"

	"github.com/stretchr/testify/assert"
)

func TestFindOneByID(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	mocker := mock.New(true)
	q := query.New(mocker, c.DB)

	cs := store.NewUserStore(mocker, c.DB, q)
	user := store.NewUser()

	ID, err := cs.Create("1", "a", "b", "c", "d")
	assert.Equal(t, "1", ID)
	assert.Nil(t, err)

	e := errors.New("ok")
	q.Mock.Add("Q.FindOneByID", true, e)
	exists, err := q.FindOneByID(&user, "99")
	assert.Equal(t, true, exists)
	assert.Equal(t, e, err)

	exists, err = q.FindOneByID(&user, "1")
	assert.Equal(t, true, exists)
	assert.Nil(t, err, err)

	user.TableName = "bad"
	user.PrimaryKeyName = "bad"
	exists, err = q.FindOneByID(&user, "1")
	assert.Equal(t, false, exists)
	assert.NotNil(t, err)
}

func TestFindOneByField(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	mocker := mock.New(true)
	q := query.New(mocker, c.DB)

	cs := store.NewUserStore(mocker, c.DB, q)
	user := store.NewUser()

	ID, err := cs.Create("1", "a", "b", "c", "d")
	assert.Equal(t, "1", ID)
	assert.Nil(t, err)

	e := errors.New("ok")
	q.Mock.Add("Q.FindOneByField", true, e)
	exists, err := q.FindOneByField(&user, "id", "99")
	assert.Equal(t, true, exists)
	assert.Equal(t, e, err)

	exists, err = q.FindOneByField(&user, "id", "1")
	assert.Equal(t, true, exists)
	assert.Nil(t, err, err)

	user.TableName = "bad"
	user.PrimaryKeyName = "bad"
	exists, err = q.FindOneByField(&user, "id", "1")
	assert.Equal(t, false, exists)
	assert.NotNil(t, err)
}

func TestFindAll(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	mocker := mock.New(true)
	q := query.New(mocker, c.DB)

	cs := store.NewUserStore(mocker, c.DB, q)
	group := store.NewGroup()
	user := store.NewUser()

	total, err := q.FindAll(&user)
	assert.Equal(t, 0, total)
	assert.NotNil(t, err)

	ID, err := cs.Create("1", "a", "b", "c", "d")
	assert.Equal(t, "1", ID)
	assert.Nil(t, err)

	e := errors.New("ok")
	q.Mock.Add("Q.FindAll", 99, e)
	total, err = q.FindAll(&group)
	assert.Equal(t, 99, total)
	assert.Equal(t, e, err)

	total, err = q.FindAll(&group)
	assert.Equal(t, 1, total)
	assert.Nil(t, err, err)

	user.TableName = "bad"
	user.PrimaryKeyName = "bad"
	total, err = q.FindAll(&user)
	assert.Equal(t, 0, total)
	assert.NotNil(t, err)
}

func TestDeleteOneByID(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	mocker := mock.New(true)
	q := query.New(mocker, c.DB)

	cs := store.NewUserStore(mocker, c.DB, q)
	user := store.NewUser()

	ID, err := cs.Create("1", "a", "b", "c", "d")
	assert.Equal(t, "1", ID)
	assert.Nil(t, err)

	e := errors.New("ok")
	q.Mock.Add("Q.DeleteOneByID", 99, e)
	affected, err := q.DeleteOneByID(&user, "99")
	assert.Equal(t, 99, affected)
	assert.Equal(t, e, err)

	user.TableName = "bad"
	user.PrimaryKeyName = "bad"
	affected, err = q.DeleteOneByID(&user, "1")
	assert.Equal(t, 0, affected)
	assert.NotNil(t, err)

	user.TableName = "user"
	user.PrimaryKeyName = "id"
	affected, err = q.DeleteOneByID(&user, "1")
	assert.Equal(t, 1, affected)
	assert.Nil(t, err, err)
}

func TestDeleteAll(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	mocker := mock.New(true)
	q := query.New(mocker, c.DB)

	cs := store.NewUserStore(mocker, c.DB, q)
	user := store.NewUser()

	ID, err := cs.Create("1", "a", "b", "c", "d")
	assert.Equal(t, "1", ID)
	assert.Nil(t, err)

	ID, err = cs.Create("2", "a", "b", "c2", "d")
	assert.Equal(t, "2", ID)
	assert.Nil(t, err)

	e := errors.New("ok")
	q.Mock.Add("Q.DeleteAll", 99, e)
	affected, err := q.DeleteAll(&user)
	assert.Equal(t, 99, affected)
	assert.Equal(t, e, err)

	user.TableName = "bad"
	user.PrimaryKeyName = "bad"
	affected, err = q.DeleteAll(&user)
	assert.Equal(t, 0, affected)
	assert.NotNil(t, err)

	user.TableName = "user"
	user.PrimaryKeyName = "id"
	affected, err = q.DeleteAll(&user)
	assert.Equal(t, 2, affected)
	assert.Nil(t, err, err)
}

func TestExistsByID(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	mocker := mock.New(true)
	q := query.New(mocker, c.DB)

	cs := store.NewUserStore(mocker, c.DB, q)
	user := store.NewUser()

	ID, err := cs.Create("1", "a", "b", "c", "d")
	assert.Equal(t, "1", ID)
	assert.Nil(t, err)

	e := errors.New("ok")
	q.Mock.Add("Q.ExistsByID", true, e)
	exists, err := q.ExistsByID(&user, "99")
	assert.Equal(t, true, exists)
	assert.Equal(t, e, err)

	exists, err = q.ExistsByID(&user, "1")
	assert.Equal(t, true, exists)
	assert.Nil(t, err, err)

	exists, err = q.ExistsByID(&user, "12")
	assert.Equal(t, false, exists)
	assert.Nil(t, err)

	user.TableName = "bad"
	user.PrimaryKeyName = "bad"
	exists, err = q.ExistsByID(&user, "1")
	assert.Equal(t, false, exists)
	assert.NotNil(t, err)
}

func TestExistsByField(t *testing.T) {
	c := testutil.Setup()
	defer c.Teardown()

	mocker := mock.New(true)
	q := query.New(mocker, c.DB)

	cs := store.NewUserStore(mocker, c.DB, q)
	user := store.NewUser()

	ID, err := cs.Create("1", "a", "b", "c", "d")
	assert.Equal(t, "1", ID)
	assert.Nil(t, err)

	e := errors.New("ok")
	q.Mock.Add("Q.ExistsByField", true, "99", e)
	exists, ID, err := q.ExistsByField(&user, "id", "99")
	assert.Equal(t, true, exists)
	assert.Equal(t, "99", ID)
	assert.Equal(t, e, err)

	exists, ID, err = q.ExistsByField(&user, "id", "1")
	assert.Equal(t, true, exists)
	assert.Equal(t, "1", ID)
	assert.Nil(t, err, err)

	exists, ID, err = q.ExistsByField(&user, "id", "12")
	assert.Equal(t, false, exists)
	assert.Equal(t, "", ID)
	assert.Nil(t, err)

	user.TableName = "bad"
	user.PrimaryKeyName = "bad"
	exists, ID, err = q.ExistsByField(&user, "id", "1")
	assert.Equal(t, false, exists)
	assert.Equal(t, "", ID)
	assert.NotNil(t, err)
}
