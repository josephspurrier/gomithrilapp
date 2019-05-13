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

// CoreMock contains all the configurable dependencies.
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
	// Set up the dependencies.
	mockLogger := new(testutil.MockLogger)
	mux := router.New()
	mockQuery := query.New(db)
	binder := bind.New()
	resp := response.New()
	mockToken := new(testutil.MockToken)
	pass := passhash.New()

	// Set up the core.
	core := NewCore(
		mockLogger,
		mux,
		db,
		mockQuery,
		binder,
		resp,
		mockToken,
		pass)

	// Add all the configurable mocks.
	m := &CoreMock{
		Log:      mockLogger,
		DB:       db,
		Q:        mockQuery,
		Bind:     binder,
		Response: resp,
		Token:    mockToken,
		Password: pass,
	}

	return core, m
}
