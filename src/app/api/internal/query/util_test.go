package query_test

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"app/api/pkg/database"

	"github.com/josephspurrier/rove/pkg/adapter/mysql"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// LoadDatabase will set up the DB and apply migrations for the tests.
func LoadDatabase() *database.DBW {
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

	db, err := database.Migrate(nil, con, changesets)
	if err != nil {
		log.Fatalln(err.Error())
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

// changesets is the database migrations.
var changesets = `
--changeset josephspurrier:1
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';
CREATE TABLE user_status (
    id TINYINT(1) UNSIGNED NOT NULL AUTO_INCREMENT,
    
    status VARCHAR(25) NOT NULL,
    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted TINYINT(1) UNSIGNED NOT NULL DEFAULT 0,
    
    PRIMARY KEY (id)
);
--rollback DROP TABLE user_status;

--changeset josephspurrier:2
INSERT INTO user_status (id, status, created_at, updated_at, deleted) VALUES
(1, 'active',   CURRENT_TIMESTAMP,  CURRENT_TIMESTAMP,  0),
(2, 'inactive', CURRENT_TIMESTAMP,  CURRENT_TIMESTAMP,  0);
--rollback TRUNCATE TABLE user_status;

--changeset josephspurrier:3
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';
CREATE TABLE user (
    id VARCHAR(36) NOT NULL,
    
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password CHAR(60) NOT NULL,
    
    status_id TINYINT(1) UNSIGNED NOT NULL DEFAULT 1,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT 0,
    
    UNIQUE KEY (email),
    CONSTRAINT f_user_status FOREIGN KEY (status_id) REFERENCES user_status (id) ON DELETE CASCADE ON UPDATE CASCADE,
    
    PRIMARY KEY (id)
);
--rollback DROP TABLE user;
`
