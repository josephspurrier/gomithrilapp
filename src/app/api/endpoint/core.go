package endpoint

import (
	"app/api"
	"app/api/pkg/logger"
	"app/api/pkg/router"
	"app/api/store"
)

// Core represents the core for all the endpoints.
type Core struct {
	Log      logger.ILog
	Router   *router.Mux
	Bind     api.IBind
	Response api.IResponse
	Token    api.IToken
	Password api.IPassword
	Store    *store.Factory
}

// NewCore returns a core for all the endpoints.
func NewCore(l logger.ILog,
	mux *router.Mux,
	binder api.IBind,
	resp api.IResponse,
	token api.IToken,
	pass api.IPassword,
	storeFactory *store.Factory) Core {
	c := Core{
		Log:      l,
		Router:   mux,
		Bind:     binder,
		Response: resp,
		Token:    token,
		Password: pass,
		Store:    storeFactory,
	}

	return c
}
