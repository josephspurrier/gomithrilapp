package store

import (
	"time"

	"app/api"
	"app/api/internal/query"
	"app/api/pkg/mock"
)

// User is a user of the system.
type User struct {
	ID        string     `db:"id"`
	FirstName string     `db:"first_name"`
	LastName  string     `db:"last_name"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	StatusID  uint8      `db:"status_id"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`

	TableName      string
	PrimaryKeyName string
}

// NewUser returns an empty user.
func NewUser() User {
	user := User{
		TableName:      "user",
		PrimaryKeyName: "id",
	}
	return user
}

// Table returns the table name.
func (x *User) Table() string {
	return x.TableName
}

// PrimaryKey returns the primary key field.
func (x *User) PrimaryKey() string {
	return x.PrimaryKeyName
}

// NewGroup returns an empty group.
func NewGroup() UserGroup {
	group := make(UserGroup, 0)
	return group
}

// UserGroup represents a group of users.
type UserGroup []User

// Table returns the table name.
func (x UserGroup) Table() string {
	return "user"
}

// PrimaryKey returns the primary key field.
func (x UserGroup) PrimaryKey() string {
	return "id"
}

// UserStore is a user of the system.
type UserStore struct {
	mock *mock.Mocker
	db   api.IDatabase
	*query.Q
}

// NewUserStore returns a new query object.
func NewUserStore(m *mock.Mocker, db api.IDatabase, q *query.Q) UserStore {
	return UserStore{
		mock: m,
		db:   db,
		Q:    q,
	}
}

// Create adds a new user.
func (x *UserStore) Create(uuid, firstName, lastName, email, password string) (string, error) {
	if x.mock != nil && x.mock.Enabled() {
		return x.mock.String(), x.mock.Error()
	}

	_, err := x.db.Exec(`
		INSERT INTO user
		(id, first_name, last_name, email, password, status_id)
		VALUES
		(?,?,?,?,?,?)
		`,
		uuid, firstName, lastName, email, password, 1)

	return uuid, err
}
