package testutil

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"app/api/migration"
	"app/api/pkg/database"
	"app/api/pkg/logger"

	"github.com/josephspurrier/rove/pkg/adapter/mysql"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// LoadDatabase will set up the DB and apply migrations for the tests.
func LoadDatabase(ml logger.ILog) *database.DBW {
	unique := "T" + fmt.Sprint(rand.Intn(500))

	password := "password"
	if os.Getenv("TRAVIS") == "true" {
		password = ""
	}

	// Set the database connection information.
	con := &mysql.Connection{
		Hostname:  "127.0.0.1",
		Username:  "root",
		Password:  password,
		Name:      "govueapptest" + unique,
		Port:      3306,
		Parameter: "parseTime=true&allowNativePasswords=true&collation=utf8mb4_unicode_ci&multiStatements=true",
	}

	db, err := database.Migrate(ml, con, migration.Changesets)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return database.New(db, con.Name)
}

// TeardownDatabase will destroy the test database.
func TeardownDatabase(db *database.DBW) {
	_, err := db.Exec(`DROP DATABASE IF EXISTS ` + db.Name())
	if err != nil {
		fmt.Println("DB DROP TEARDOWN Error:", err)
	}
}
