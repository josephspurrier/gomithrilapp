package component

import (
	"app/api/iface"

	"github.com/husobee/vestigo"
	"github.com/jmoiron/sqlx"
)

// Core represents the core for all the endpoints.
type Core struct {
	DB       iface.IDatabase
	Q        iface.IQuery
	Router   *vestigo.Router
	Password iface.IPassword
}

// NewCore returs a core for all the endpoints.
func NewCore(db *sqlx.DB, q iface.IQuery, p iface.IPassword) Core {
	return Core{
		DB:       db,
		Q:        q,
		Router:   vestigo.NewRouter(),
		Password: p,
	}
}
