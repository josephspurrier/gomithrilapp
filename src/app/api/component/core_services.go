package component

import (
	"os"

	"app/api"
	"app/api/internal/bind"
	"app/api/internal/response"
	"app/api/pkg/database"
	"app/api/pkg/logger"
	"app/api/pkg/passhash"
	"app/api/pkg/query"
	"app/api/pkg/router"
	"app/api/pkg/webtoken"

	"github.com/josephspurrier/rove/pkg/adapter/mysql"
)

// Services will set up the production services.
func Services(l logger.ILog) Core {
	// If the host env var is set, use it.
	host := os.Getenv("MYSQL_HOST")
	if len(host) == 0 {
		host = "127.0.0.1"
	}

	// If the password env var is set, use it.
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	if len(password) == 0 {
		password = "password"
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

	db, err := database.Migrate(l, con, api.Changesets)
	if err != nil {
		l.Fatalf(err.Error())
	}

	// Configure the services.
	r := router.New()
	db2 := database.New(db, con.Name)
	q := query.New(db2)
	p := passhash.New()
	resp := response.New()
	b := bind.New()

	// Setup middleware.
	secret := "TA8tALZAvLVLo4ToI44xF/nF6IyrRNOR6HSfpno/81M="
	t := webtoken.New([]byte(secret))

	return NewCore(l, r, db2, q, b, resp, t, p)
}
