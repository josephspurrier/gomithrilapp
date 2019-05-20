package component

import (
	"app/api/store"
)

// Store is a list of the stores in the application.
type Store struct {
	User store.User
}

// LoadStores will load the stores.
func LoadStores(core Core) *Store {
	s := new(Store)

	s.User = store.NewUser(core.Mock, core.DB, core.Q)

	return s
}
