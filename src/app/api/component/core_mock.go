package component

import (
	"app/api/internal/bind"
	"app/api/internal/response"
	"app/api/internal/testutil"
	"app/api/pkg/database"
	"app/api/pkg/mock"
	"app/api/pkg/passhash"
	"app/api/pkg/query"
	"app/api/pkg/router"
)

// CoreMock contains all the configurable dependencies.
type CoreMock struct {
	Log   *testutil.MockLogger
	Token *testutil.MockToken
}

// NewCoreMock returns all mocked dependencies.
func NewCoreMock(db *database.DBW) (Core, *CoreMock) {
	// Set up the dependencies.
	mocker := mock.New(true)
	mockLogger := new(testutil.MockLogger)
	mux := router.New()
	mockQuery := query.New(mocker, db)
	binder := bind.New(mux)
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
		pass,
		mocker,
	)

	core.Store = LoadStores(core)

	// Add all the configurable mocks.
	m := &CoreMock{
		Log:   mockLogger,
		Token: mockToken,
	}

	return core, m
}
