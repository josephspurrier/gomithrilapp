package api

import (
	"app/api/pkg/query"
	"database/sql"
	"net/http"
	"time"
)

// IBind provides bind and validation for requests.
type IBind interface {
	Unmarshal(i interface{}, r *http.Request) (err error)
	Validate(s interface{}) error
}

// IDatabase provides data query capabilities.
type IDatabase interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	Name() string
}

// IPassword provides password hashing.
type IPassword interface {
	HashString(password string) (string, error)
	MatchString(hash, password string) bool
}

// IQuery provides default queries.
type IQuery interface {
	FindOneByID(dest query.IRecord, ID string) (found bool, err error)
	FindAll(dest query.IRecord) (total int, err error)
	ExistsByID(db query.IRecord, s string) (found bool, err error)
	ExistsByField(db query.IRecord, field string, value string) (found bool, ID string, err error)
	DeleteOneByID(dest query.IRecord, ID string) (affected int, err error)
	DeleteAll(dest query.IRecord) (affected int, err error)
}

// IResponse provides outputs for data.
type IResponse interface {
	JSON(w http.ResponseWriter, body interface{}) (int, error)
	Created(w http.ResponseWriter, recordID string) (int, error)
	OK(w http.ResponseWriter, message string) (int, error)
}

// IToken provides outputs for the JWT.
type IToken interface {
	Generate(userID string, duration time.Duration) (string, error)
	Secret() []byte
}
