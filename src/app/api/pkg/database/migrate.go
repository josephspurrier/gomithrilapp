package database

import (
	"fmt"

	"app/api/pkg/logger"

	"github.com/jmoiron/sqlx"
	"github.com/josephspurrier/rove"
	"github.com/josephspurrier/rove/pkg/adapter/mysql"
)

// Migrate will run the database migrations and will create the database if it
// does not exist.
func Migrate(l logger.ILog, con *mysql.Connection, changesets string) (*sqlx.DB, error) {
	// Connect to the database.
	db, err := mysql.New(con)
	if err != nil {
		// Attempt to connect without the database name.
		d, err := con.Connect(false)
		if err != nil {
			return nil, err
		}

		// Create the database.
		_, err = d.Query(fmt.Sprintf(`CREATE DATABASE IF NOT EXISTS %v DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;`, con.Name))
		if err != nil {
			return nil, err
		}
		l.Printf("Database created.")

		// Attempt to reconnect with the database name.
		db, err = mysql.New(con)
		if err != nil {
			return nil, err
		}
	}

	// Perform all migrations against the database.
	r := rove.NewChangesetMigration(db, changesets)
	r.Verbose = false
	return db.DB, r.Migrate(0)
}
