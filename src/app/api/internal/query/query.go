package query

import (
	"fmt"

	"app/api"
	"app/api/pkg/mock"
)

// New returns a new query object.
func New(m *mock.Mocker, db api.IDatabase) *Q {
	return &Q{
		Mock: m,
		db:   db,
	}
}

// Q is a database wrapper that provides helpful utilities.
type Q struct {
	Mock *mock.Mocker
	db   api.IDatabase
}

// *****************************************************************************
// Find
// *****************************************************************************

// FindOneByID will find a record by string ID.
func (q *Q) FindOneByID(dest api.IRecord, ID string) (exists bool, err error) {
	if q.Mock != nil && q.Mock.Enabled() {
		return q.Mock.Bool(), q.Mock.Error()
	}

	err = q.db.Get(dest, fmt.Sprintf(`
		SELECT * FROM %s
		WHERE %s = ?
		LIMIT 1`, dest.Table(), dest.PrimaryKey()),
		ID)
	return q.db.RecordExists(err)
}

// FindOneByField will find a record by a specified field.
func (q *Q) FindOneByField(dest api.IRecord, field string, value string) (exists bool, err error) {
	if q.Mock != nil && q.Mock.Enabled() {
		return q.Mock.Bool(), q.Mock.Error()
	}

	err = q.db.Get(dest, fmt.Sprintf(`
		SELECT * FROM %s
		WHERE %s = ?
		LIMIT 1`, dest.Table(), field),
		value)
	return q.db.RecordExists(err)
}

// FindAll returns all users.
func (q *Q) FindAll(dest api.IRecord) (total int, err error) {
	if q.Mock != nil && q.Mock.Enabled() {
		return q.Mock.Int(), q.Mock.Error()
	}

	err = q.db.QueryRowScan(&total, fmt.Sprintf(`
		SELECT COUNT(DISTINCT %s)
		FROM %s
		`, dest.PrimaryKey(), dest.Table()))

	if err != nil {
		return total, q.db.SuppressNoRowsError(err)
	}

	err = q.db.Select(dest, fmt.Sprintf(`SELECT * FROM %s`, dest.Table()))
	return total, err
}

// *****************************************************************************
// Delete
// *****************************************************************************

// DeleteOneByID removes one record by ID.
func (q *Q) DeleteOneByID(dest api.IRecord, ID string) (affected int, err error) {
	if q.Mock != nil && q.Mock.Enabled() {
		return q.Mock.Int(), q.Mock.Error()
	}

	result, err := q.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE %s = ? LIMIT 1",
		dest.Table(), dest.PrimaryKey()), ID)
	if err != nil {
		return 0, err
	}

	return q.db.AffectedRows(result), err
}

// DeleteAll removes all records.
func (q *Q) DeleteAll(dest api.IRecord) (affected int, err error) {
	if q.Mock != nil && q.Mock.Enabled() {
		return q.Mock.Int(), q.Mock.Error()
	}

	result, err := q.db.Exec(fmt.Sprintf(`DELETE FROM %s`, dest.Table()))
	if err != nil {
		return 0, err
	}

	return q.db.AffectedRows(result), err
}

// *****************************************************************************
// Exists
// *****************************************************************************

// ExistsByID determines if a records exists by ID.
func (q *Q) ExistsByID(db api.IRecord, value string) (found bool, err error) {
	if q.Mock != nil && q.Mock.Enabled() {
		return q.Mock.Bool(), q.Mock.Error()
	}

	err = q.db.Get(db, fmt.Sprintf(`
		SELECT %s FROM %s
		WHERE %s = ?
		LIMIT 1`, db.PrimaryKey(), db.Table(), db.PrimaryKey()),
		value)
	return q.db.RecordExists(err)
}

// ExistsByField determines if a records exists by a specified field and
// returns the ID.
func (q *Q) ExistsByField(db api.IRecord, field string, value string) (found bool, ID string, err error) {
	if q.Mock != nil && q.Mock.Enabled() {
		return q.Mock.Bool(), q.Mock.String(), q.Mock.Error()
	}

	err = q.db.QueryRowScan(&ID, fmt.Sprintf(`
		SELECT %s FROM %s
		WHERE %s = ?
		LIMIT 1`, db.PrimaryKey(), db.Table(), field),
		value)

	return q.db.RecordExistsString(err, ID)
}
