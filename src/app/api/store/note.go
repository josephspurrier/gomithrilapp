package store

import (
	"database/sql"
	"time"

	"app/api"
	"app/api/pkg/securegen"
)

// Note is a note of a user.
type Note struct {
	ID        string     `db:"id"`
	UserID    string     `db:"user_id"`
	Message   string     `db:"message"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

// New returns an empty item.
func (x *NoteStore) New() Note {
	item := Note{}
	return item
}

// Table returns the table name.
func (x *Note) Table() string {
	return "note"
}

// PrimaryKey returns the primary key field.
func (x *Note) PrimaryKey() string {
	return "id"
}

// NewGroup returns an empty group.
func (x *NoteStore) NewGroup() NoteGroup {
	group := make(NoteGroup, 0)
	return group
}

// NoteGroup represents a group.
type NoteGroup []Note

// Table returns the table name.
func (x NoteGroup) Table() string {
	return "note"
}

// PrimaryKey returns the primary key field.
func (x NoteGroup) PrimaryKey() string {
	return "id"
}

// NewNoteStore returns a new query object.
func NewNoteStore(c Core) NoteStore {
	return NoteStore{
		Core: c,
	}
}

// NoteStore .
type NoteStore struct {
	Core
}

// Create adds a new item.
func (x *NoteStore) Create(userID, message string) (string, error) {
	if x.mock != nil && x.mock.Enabled() {
		return x.mock.String(), x.mock.Error()
	}

	uuid, err := securegen.UUID()
	if err != nil {
		return "", err
	}

	_, err = x.db.Exec(`
		INSERT INTO note
		(id, user_id, message)
		VALUES
		(?,?,?)
		`,
		uuid, userID, message)

	return uuid, err
}

// Update makes changes to an item.
func (x *NoteStore) Update(ID, userID, message string) (err error) {
	if x.mock != nil && x.mock.Enabled() {
		return x.mock.Error()
	}

	_, err = x.db.Exec(`
		UPDATE note
		SET
			message = ?
		WHERE id = ?
		AND user_id = ?
		LIMIT 1
		`,
		message, ID, userID)
	return
}

// FindAllByUser returns items for a user.
func (x *NoteStore) FindAllByUser(dest *NoteGroup, userID string) (err error) {
	if x.mock != nil && x.mock.Enabled() {
		return x.mock.Error()
	}

	err = x.db.Select(dest, `
		SELECT *
		FROM note
		WHERE user_id = ?
		`,
		userID)
	return
}

// FindOneByIDAndUser returns an item for a user.
func (x *NoteStore) FindOneByIDAndUser(dest *Note, ID string, userID string) (err error) {
	if x.mock != nil && x.mock.Enabled() {
		return x.mock.Error()
	}

	err = x.db.Get(dest, `
		SELECT *
		FROM note
		WHERE id = ?
		AND user_id = ?
		LIMIT 1
		`,
		ID, userID)
	return
}

// DeleteOneByIDAndUser removes one item from a user.
func (x *NoteStore) DeleteOneByIDAndUser(dest api.IRecord, ID string,
	userID string) (affected int, err error) {
	if x.mock != nil && x.mock.Enabled() {
		return x.mock.Int(), x.mock.Error()
	}

	result, err := x.db.Exec(`
	DELETE FROM note
	WHERE id = ?
	AND user_id = ?
	LIMIT 1`,
		ID, userID)
	if err != nil {
		return 0, err
	}

	return affectedRows(result), err
}

// affectedRows returns the number of rows affected by the query.
func affectedRows(result sql.Result) int {
	if result == nil {
		return 0
	}

	// If successful, get the number of affected rows.
	count, err := result.RowsAffected()
	if err != nil {
		return 0
	}

	return int(count)
}
