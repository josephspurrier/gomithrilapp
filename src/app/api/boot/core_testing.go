package boot

import (
	"app/api/endpoint"
	"app/api/internal/testutil"
	"app/api/pkg/database"
	"app/api/pkg/mock"
)

// CoreTest contains all the configurable dependencies.
type CoreTest struct {
	Log   *testutil.MockLogger
	Token *testutil.MockToken
	Mock  *mock.Mocker
}

// TestServices sets up the test services.
func TestServices(db *database.DBW) (endpoint.Core, *CoreTest) {
	// Set up the mocked dependencies.
	mockLogger := new(testutil.MockLogger)
	mockToken := new(testutil.MockToken)
	mocker := mock.New(true)

	core := Services(mockLogger, db, mocker)

	// Add all the configurable mocks.
	m := &CoreTest{
		Log:   mockLogger,
		Token: mockToken,
		Mock:  mocker,
	}

	return core, m
}
