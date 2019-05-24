package store

import (
	"time"

	"app/api/pkg/securegen"
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
}

// New returns an empty user.
func (x *UserStore) New() User {
	user := User{}
	return user
}

// Table returns the table name.
func (x *User) Table() string {
	return "user"
}

// PrimaryKey returns the primary key field.
func (x *User) PrimaryKey() string {
	return "id"
}

// NewGroup returns an empty group.
func (x *UserStore) NewGroup() UserGroup {
	group := make(UserGroup, 0)
	return group
}

// UserGroup represents a group.
type UserGroup []User

// Table returns the table name.
func (x UserGroup) Table() string {
	return "user"
}

// PrimaryKey returns the primary key field.
func (x UserGroup) PrimaryKey() string {
	return "id"
}

// NewUserStore returns a new query object.
func NewUserStore(c Core) UserStore {
	return UserStore{
		Core: c,
	}
}

// UserStore .
type UserStore struct {
	Core
}

// Create adds a new item.
func (x *UserStore) Create(firstName, lastName, email, password string) (string, error) {
	if x.mock != nil && x.mock.Enabled() {
		return x.mock.String(), x.mock.Error()
	}

	uuid, err := securegen.UUID()
	if err != nil {
		return "", err
	}

	_, err = x.db.Exec(`
		INSERT INTO user
		(id, first_name, last_name, email, password, status_id)
		VALUES
		(?,?,?,?,?,?)
		`,
		uuid, firstName, lastName, email, password, 1)

	return uuid, err
}

// Update makes changes to an item.
func (x *UserStore) Update(ID, firstName, lastName, email, password string) (err error) {
	if x.mock != nil && x.mock.Enabled() {
		return x.mock.Error()
	}

	_, err = x.db.Exec(`
		UPDATE user
		SET
			first_name = ?,
			last_name = ?,
			email = ?,
			password = ?
		WHERE id = ?
		`,
		firstName, lastName, email, password, ID)
	return
}
