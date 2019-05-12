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

// NewCore returs a core for all the endpoints.
func NewCore(l logger.ILog, r *router.Mux, db api.IDatabase, q api.IQuery,
	b api.IBind, resp api.IResponse, token api.IToken, p api.IPassword) Core {
	return Core{
		Log:      l,
		Router:   r,
		DB:       db,
		Q:        q,
		Bind:     b,
		Response: resp,
		Token:    token,
		Password: p,
	}
}
