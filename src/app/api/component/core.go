package component

import (
	"app/api/iface"
	"app/api/pkg/logger"
	"app/api/pkg/router"

	"github.com/jmoiron/sqlx"
)

// Core represents the core for all the endpoints.
type Core struct {
	Log      logger.ILog
	Router   *router.Mux
	DB       iface.IDatabase
	Q        iface.IQuery
	Bind     iface.IBind
	Response iface.IResponse
	Token    iface.IToken
	Password iface.IPassword
}

// NewCore returs a core for all the endpoints.
func NewCore(l logger.ILog, r *router.Mux, db *sqlx.DB, q iface.IQuery,
	b iface.IBind, resp iface.IResponse, token iface.IToken, p iface.IPassword) Core {
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
