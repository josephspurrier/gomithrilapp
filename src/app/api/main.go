package main

import (
	"app/api/iface"
	"log"
	"net/http"

	"app/api/component"
	"app/api/middleware"
	"app/api/pkg/database"
	"app/api/pkg/passhash"
	"app/api/pkg/query"

	"github.com/husobee/vestigo"
	"github.com/jmoiron/sqlx"
	"github.com/josephspurrier/rove"
	"github.com/josephspurrier/rove/pkg/adapter/mysql"
)

func init() {
	// Verbose logging with file name and line number.
	log.SetFlags(log.Lshortfile)
}

func main() {
	port := "8081"

	db, err := LoadMigrations()
	if err != nil {
		log.Fatalln(err)
	}

	db2 := database.New(db)
	q := query.New(db2)
	p := passhash.New()

	router := LoadRoutes(db, q, p)
	router.SetGlobalCors(&vestigo.CorsAccessControl{
		AllowOrigin:  []string{"*"},
		AllowHeaders: []string{"Content-Type", "Origin", "X-Requested-With", "Accept"},
	})

	log.Println("Server started.")
	err = http.ListenAndServe(":"+port, middleware.Log(router))
	if err != nil {
		log.Println(err)
	}
}

// LoadRoutes will load the endpoints.
func LoadRoutes(db *sqlx.DB, q iface.IQuery, p iface.IPassword) *vestigo.Router {
	core := component.NewCore(db, q, p)

	component.SetupStatic(core)
	component.SetupLogin(core)
	component.SetupRegister(core)

	return core.Router
}

// LoadMigrations will run the database migrations.
func LoadMigrations() (*sqlx.DB, error) {
	// Create a new MySQL database object.
	db, err := mysql.New(&mysql.Connection{
		Hostname:  "127.0.0.1",
		Username:  "root",
		Password:  "password",
		Name:      "main",
		Port:      3306,
		Parameter: "collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true",
	})
	if err != nil {
		return nil, err
	}

	// Perform all migrations against the database.
	r := rove.NewChangesetMigration(db, changesets)
	r.Verbose = true
	return db.DB, r.Migrate(0)
}

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
