package component

import (
	"app/api/store"
)

// Store is a list of the stores in the application.
type Store struct {
	User store.UserStore
}

// LoadStores will load the stores.
func LoadStores(core Core) *Store {
	s := new(Store)

	s.User = store.NewUserStore(core.Mock, core.DB, core.Q)

	return s
}
