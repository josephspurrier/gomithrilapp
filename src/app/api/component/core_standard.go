package component

import (
	"app/api"
	"app/api/pkg/logger"
	"app/api/pkg/router"
)

// Core represents the core for all the endpoints.
type Core struct {
	Log      logger.ILog
	Router   *router.Mux
	DB       api.IDatabase
	Q        api.IQuery
	Bind     api.IBind
	Response api.IResponse
	Token    api.IToken
	Password api.IPassword
}

// NewCore returns a core for all the endpoints.
func NewCore(l logger.ILog,
	mux *router.Mux,
	db api.IDatabase,
	query api.IQuery,
	binder api.IBind,
	resp api.IResponse,
	token api.IToken,
	pass api.IPassword) Core {
	return Core{
		Log:      l,
		Router:   mux,
		DB:       db,
		Q:        query,
		Bind:     binder,
		Response: resp,
		Token:    token,
		Password: pass,
	}
}
