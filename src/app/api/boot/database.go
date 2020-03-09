package boot

import (
	"os"

	"app/api/migration"
	"app/api/pkg/database"
	"app/api/pkg/logger"

	"github.com/josephspurrier/rove/pkg/adapter/mysql"
)

// Database migrates the database and then returns the database connection.
func Database(l logger.ILog) *database.DBW {
	// If the host env var is set, use it.
	host := os.Getenv("MYSQL_HOST")
	if len(host) == 0 {
		host = "127.0.0.1"
	}

	// If the password env var is set, use it.
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	if len(password) == 0 {
		if os.Getenv("TRAVIS") == "true" {
			password = ""
		} else {
			password = "password"
		}
	}

	// Set the database connection information.
	con := &mysql.Connection{
		Hostname:  host,
		Username:  "root",
		Password:  password,
		Name:      "main",
		Port:      3306,
		Parameter: "collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true",
	}

	// Migration the database.
	dbx, err := database.Migrate(l, con, migration.Changesets)
	if err != nil {
		l.Fatalf(err.Error())
	}

	return database.New(dbx, con.Name)
}
