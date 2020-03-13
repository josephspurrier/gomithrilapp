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
	Context  api.IContext
}

// NewCore returns a core for all the endpoints.
func NewCore(l logger.ILog,
	mux *router.Mux,
	bind api.IBind,
	resp api.IResponse,
	token api.IToken,
	pass api.IPassword,
	storeFactory *store.Factory,
	ctx api.IContext) Core {
	c := Core{
		Log:      l,
		Router:   mux,
		Bind:     bind,
		Response: resp,
		Token:    token,
		Password: pass,
		Store:    storeFactory,
		Context:  ctx,
	}

	return c
}
