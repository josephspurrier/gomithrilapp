package boot

import (
	"app/api/endpoint"
	"app/api/internal/bind"
	"app/api/internal/query"
	"app/api/internal/requestcontext"
	"app/api/internal/response"
	"app/api/pkg/database"
	"app/api/pkg/logger"
	"app/api/pkg/mock"
	"app/api/pkg/passhash"
	"app/api/pkg/router"
	"app/api/pkg/webtoken"
	"app/api/store"
)

// Services sets up the production services.
func Services(l logger.ILog, dbx *database.DBW, mocker *mock.Mocker) endpoint.Core {
	// FIXME: This needs to be loaded from a config.
	secret := "TA8tALZAvLVLo4ToI44xF/nF6IyrRNOR6HSfpno/81M="

	// Configure the services.
	mux := router.New()

	// Return a new core.
	core := endpoint.NewCore(
		l,
		mux,
		bind.New(mocker, mux),
		response.New(),
		webtoken.New([]byte(secret)),
		passhash.New(),
		store.LoadFactory(mocker,
			dbx,
			query.New(mocker, dbx),
		),
		requestcontext.New(mocker),
	)

	// Set up the router.
	SetupRouter(l, mux)

	return core
}
