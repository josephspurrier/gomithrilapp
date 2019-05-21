package boot

import (
	"app/api/component"
	"app/api/internal/bind"
	"app/api/internal/response"
	"app/api/internal/testutil"
	"app/api/pkg/database"
	"app/api/pkg/mock"
	"app/api/pkg/passhash"
	"app/api/pkg/query"
	"app/api/pkg/router"
	"app/api/store"
)

// CoreTest contains all the configurable dependencies.
type CoreTest struct {
	Log   *testutil.MockLogger
	Token *testutil.MockToken
	Mock  *mock.Mocker
}

// TestServices sets up the test services.
func TestServices(db *database.DBW) (component.Core, *CoreTest) {
	// Set up the dependencies.
	mux := router.New()
	mocker := mock.New(true)

	// Set up the mocked dependencies.
	mockLogger := new(testutil.MockLogger)
	mockToken := new(testutil.MockToken)

	// Set up the core.
	core := component.NewCore(
		mockLogger,
		mux,
		bind.New(mux),
		response.New(),
		mockToken,
		passhash.New(),
		store.LoadFactory(mocker, db, query.New(mocker, db)),
	)

	// Set up the router.
	SetupRouter(core.Log, mux)

	// Add all the configurable mocks.
	m := &CoreTest{
		Log:   mockLogger,
		Token: mockToken,
		Mock:  mocker,
	}

	return core, m
}