package store

import (
	"app/api"
	"app/api/pkg/mock"
)

// Factory is all of the stores in the application.
type Factory struct {
	User UserStore
}

// LoadFactory will return the factory.
func LoadFactory(m *mock.Mocker, db api.IDatabase, q api.IQuery) *Factory {
	cs := NewCore(m, db, q)

	return &Factory{
		User: NewUserStore(cs),
	}
}
