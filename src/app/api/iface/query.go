package iface

import "app/api/pkg/query"

// IQuery provides default queries.
type IQuery interface {
	FindOneByID(dest query.IRecord, ID string) (found bool, err error)
	FindAll(dest query.IRecord) (total int, err error)
	ExistsByID(db query.IRecord, s string) (found bool, err error)
	ExistsByField(db query.IRecord, field string, value string) (found bool, ID string, err error)
	DeleteOneByID(dest query.IRecord, ID string) (affected int, err error)
	DeleteAll(dest query.IRecord) (affected int, err error)
}
