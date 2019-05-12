package testutil

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/josephspurrier/rove/pkg/adapter/mysql"

	"app/api"
	"app/api/pkg/database"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// LoadDatabase will set up the DB and apply migrations for the tests.
func LoadDatabase() *database.DBW {
	//return LoadDatabaseFromFile("../../../../../migration/mysql-v0.sql", true)

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
		Name:      "webapitest" + unique,
		Port:      3306,
		Parameter: "parseTime=true&allowNativePasswords=true&collation=utf8mb4_unicode_ci&multiStatements=true",
	}

	ml := new(MockLogger)

	db, err := database.Migrate(ml, con, api.Changesets)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return database.New(db, con.Name)
}

// TeardownDatabase will destroy the test database and unset the environment
// variables.
func TeardownDatabase(db *database.DBW) {
	_, err := db.Exec(`DROP DATABASE IF EXISTS ` + db.Name())
	if err != nil {
		fmt.Println("DB DROP TEARDOWN Error:", err)
	}
}
