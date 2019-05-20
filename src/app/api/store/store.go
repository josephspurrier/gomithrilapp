package store

import (
	"app/api"
	"app/api/pkg/mock"
)

// Core has the dependencies for the store package.
type Core struct {
	mock *mock.Mocker
	db   api.IDatabase
	api.IQuery
}

// NewCore returns a new core object.
func NewCore(m *mock.Mocker, db api.IDatabase, q api.IQuery) Core {
	return Core{
		mock:   m,
		db:     db,
		IQuery: q,
	}
}
