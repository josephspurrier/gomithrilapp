package testutil

import (
	"app/api/config"
	"app/api/endpoint"
	"app/api/pkg/database"
	"app/api/pkg/mock"
	"app/api/pkg/requestcontext"
)

// Mocks contains all the configurable dependencies.
type Mocks struct {
	Log  *MockLogger
	Mock *mock.Mocker
	//Context *Context
}

// Services sets up the test services.
func Services(db *database.DBW) (endpoint.Core, *Mocks) {
	// Set up the mocked dependencies.
	mockLogger := new(MockLogger)
	mocker := mock.New(true)

	// Set up the mocked dependencies.
	//ctx := new(Context)
	ctx := requestcontext.New()

	// Load the environment variables from defaults.
	settings := config.LoadEnv(mockLogger, "")

	// Set up the services.
	core := config.Services(mockLogger, settings, db, mocker, ctx)

	// Add all the configurable mocks.
	m := &Mocks{
		Log:  mockLogger,
		Mock: mocker,
		//Context: ctx,
	}

	return core, m
}
