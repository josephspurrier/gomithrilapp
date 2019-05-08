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
	DB       iface.IDatabase
	Q        iface.IQuery
	Router   *router.Mux
	Response iface.IResponse
	Password iface.IPassword
}

// NewCore returs a core for all the endpoints.
func NewCore(l logger.ILog, r *router.Mux, db *sqlx.DB, q iface.IQuery,
	resp iface.IResponse, p iface.IPassword) Core {
	return Core{
		DB:       db,
		Q:        q,
		Router:   r,
		Response: resp,
		Password: p,
	}
}
