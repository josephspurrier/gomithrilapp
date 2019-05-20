package boot

import (
	"app/api"
	"app/api/pkg/mock"
	"app/api/store"
)

// Store is a list of the stores in the application.
type Store struct {
	User store.UserStore
}

// Stores will load the stores.
func Stores(m *mock.Mocker, db api.IDatabase, q api.IQuery) *Store {
	cs := store.NewCore(m, db, q)

	return &Store{
		User: store.NewUserStore(cs),
	}
}
