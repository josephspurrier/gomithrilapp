package component

import (
	"app/api"
	"app/api/boot"
	"app/api/pkg/logger"
	"app/api/pkg/mock"
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
	Mock     *mock.Mocker
	Store    *boot.Store
}

// NewCore returns a core for all the endpoints.
func NewCore(l logger.ILog,
	mux *router.Mux,
	db api.IDatabase,
	query api.IQuery,
	binder api.IBind,
	resp api.IResponse,
	token api.IToken,
	pass api.IPassword,
	mocker *mock.Mocker) Core {
	c := Core{
		Log:      l,
		Router:   mux,
		DB:       db,
		Q:        query,
		Bind:     binder,
		Response: resp,
		Token:    token,
		Password: pass,
		Mock:     mocker,
	}

	return c
}
