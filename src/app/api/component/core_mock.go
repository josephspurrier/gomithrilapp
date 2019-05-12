package component

import (
	"app/api"
	"app/api/internal/bind"
	"app/api/internal/response"
	"app/api/internal/testutil"
	"app/api/pkg/database"
	"app/api/pkg/passhash"
	"app/api/pkg/query"
	"app/api/pkg/router"
)

// CoreMock contains all the mocked dependencies.
type CoreMock struct {
	Log      *testutil.MockLogger
	DB       api.IDatabase
	Q        api.IQuery
	Bind     api.IBind
	Response api.IResponse
	Token    *testutil.MockToken
	Password api.IPassword
}

// NewCoreMock returns all mocked dependencies.
func NewCoreMock(db *database.DBW) (Core, *CoreMock) {
	ml := new(testutil.MockLogger)
	mq := query.New(db)
	mt := new(testutil.MockToken)
	resp := response.New()
	binder := bind.New()
	p := passhash.New()

	r := router.New()

	core := NewCore(ml, r, db, mq, binder, resp, mt, p)
	m := &CoreMock{
		Log:      ml,
		DB:       db,
		Q:        mq,
		Bind:     binder,
		Response: resp,
		Token:    mt,
		Password: p,
	}
	return core, m
}
