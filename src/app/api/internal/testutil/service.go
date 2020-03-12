package testutil

import (
	"app/api/config"
	"app/api/endpoint"
	"app/api/pkg/database"
	"app/api/pkg/mock"
)

// CoreTest contains all the configurable dependencies.
type CoreTest struct {
	Log   *MockLogger
	Token *MockToken
	Mock  *mock.Mocker
}

// Services sets up the test services.
func Services(db *database.DBW) (endpoint.Core, *CoreTest) {
	// Set up the mocked dependencies.
	mockLogger := new(MockLogger)
	mockToken := new(MockToken)
	mocker := mock.New(true)

	core := config.Services(mockLogger, db, mocker)

	// Add all the configurable mocks.
	m := &CoreTest{
		Log:   mockLogger,
		Token: mockToken,
		Mock:  mocker,
	}

	return core, m
}
