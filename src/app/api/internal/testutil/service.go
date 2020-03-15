package testutil

import (
	"app/api/config"
	"app/api/endpoint"
	"app/api/internal/requestcontext"
	"app/api/pkg/database"
	"app/api/pkg/mock"
)

// Mocks contains all the configurable dependencies.
type Mocks struct {
	Log  *MockLogger
	Mock *mock.Mocker
}

// Services sets up the test services.
func Services(db *database.DBW) (endpoint.Core, *Mocks) {
	// Set up the mocked dependencies.
	mockLogger := NewMockLogger()
	mocker := mock.New(true)

	// Load the environment variables from defaults.
	settings := config.LoadEnv(mockLogger, "")

	// Set up the services.
	core := config.Services(mockLogger, settings, db, mocker, requestcontext.New())

	// Add all the configurable mocks.
	m := &Mocks{
		Log:  mockLogger,
		Mock: mocker,
	}

	return core, m
}
