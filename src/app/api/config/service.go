package config

import (
	"time"

	"app/api/endpoint"
	"app/api/internal/bind"
	"app/api/internal/query"
	"app/api/internal/response"
	"app/api/pkg/database"
	"app/api/pkg/logger"
	"app/api/pkg/mock"
	"app/api/pkg/passhash"
	"app/api/pkg/requestcontext"
	"app/api/pkg/router"
	"app/api/pkg/webtoken"
	"app/api/store"
)

// Services sets up the production services.
func Services(l logger.ILog, settings *Settings, dbx *database.DBW,
	mocker *mock.Mocker) endpoint.Core {
	// Configure the services.
	mux := router.New()

	// Return a new core.
	core := endpoint.NewCore(
		l,
		mux,
		bind.New(mocker, mux),
		response.New(),
		webtoken.New([]byte(settings.Secret),
			time.Duration(settings.SessionTimeout)*time.Minute),
		passhash.New(),
		store.NewFactory(mocker,
			dbx,
			query.New(mocker, dbx),
		),
		requestcontext.New(),
	)

	// Set up the router.
	SetupRouter(l, mux)

	return core
}
