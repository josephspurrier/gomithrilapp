package api

import (
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
	QueryRowScan(dest interface{}, query string, args ...interface{}) error
	Name() string
}

// IPassword provides password hashing.
type IPassword interface {
	Hash(password string) (string, error)
	Match(hash, password string) bool
}

// IQuery provides default queries.
type IQuery interface {
	FindOneByID(dest IRecord, ID string) (found bool, err error)
	FindOneByField(dest IRecord, field string, value string) (exists bool, err error)
	FindAll(dest IRecord) (total int, err error)
	ExistsByID(db IRecord, s string) (found bool, err error)
	ExistsByField(db IRecord, field string, value string) (found bool, ID string, err error)
	DeleteOneByID(dest IRecord, ID string) (affected int, err error)
	DeleteAll(dest IRecord) (affected int, err error)
}

// IRecord provides table information.
type IRecord interface {
	Table() string
	PrimaryKey() string
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

// IContext provides handlers for type request context.
type IContext interface {
	SetUserID(r *http.Request, val string)
	UserID(r *http.Request) (string, bool)
}
