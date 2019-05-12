package database

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

var (
	// ErrNoDatabase is when there is no database connection.
	ErrNoDatabase = errors.New("no database connection is set")
)

// New returns a new database wrapper.
func New(db *sqlx.DB, name string) *DBW {
	return &DBW{
		name: name,
		db:   db,
	}
}

// DBW is a database wrapper that provides helpful utilities.
type DBW struct {
	name string
	db   *sqlx.DB
}

// Select using this DB.
// Any placeholder parameters are replaced with supplied args.
func (d *DBW) Select(dest interface{}, query string, args ...interface{}) error {
	if d == nil {
		return ErrNoDatabase
	}
	return d.db.Select(dest, query, args...)
}

// Get using this DB.
// Any placeholder parameters are replaced with supplied args.
// An error is returned if the result set is empty.
func (d *DBW) Get(dest interface{}, query string, args ...interface{}) error {
	if d == nil {
		return ErrNoDatabase
	}
	return d.db.Get(dest, query, args...)
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (d *DBW) Exec(query string, args ...interface{}) (sql.Result, error) {
	if d == nil {
		return nil, ErrNoDatabase
	}
	return d.db.Exec(query, args...)
}

// QueryRowScan returns a single result.
func (d *DBW) QueryRowScan(dest interface{}, query string, args ...interface{}) error {
	if d == nil {
		return ErrNoDatabase
	}
	return d.db.QueryRow(query, args...).Scan(dest)
}

// Name returns the database name.
func (d *DBW) Name() string {
	return d.name
}

/*
// PaginatedResults returns the paginated results of a query.
func (d *DBW) PaginatedResults(i interface{}, fn func() (interface{}, int,
	error)) (int, error) {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Ptr {
		return 0, errors.New("must pass a pointer, not a value")
	}

	results, d2, d3 := fn()
	if results == nil {
		return d2, d3
	}

	arrPtr := reflect.ValueOf(i)
	value := arrPtr.Elem()
	itemPtr := reflect.ValueOf(results)
	value.Set(itemPtr)

	return d2, d3
}*/
