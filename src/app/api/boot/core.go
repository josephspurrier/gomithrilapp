package boot

import (
	"os"

	"app/api/endpoint"
	"app/api/internal/bind"
	"app/api/internal/query"
	"app/api/internal/requestcontext"
	"app/api/internal/response"
	"app/api/migration"
	"app/api/pkg/database"
	"app/api/pkg/logger"
	"app/api/pkg/mock"
	"app/api/pkg/passhash"
	"app/api/pkg/router"
	"app/api/pkg/webtoken"
	"app/api/store"

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

	// Migration the database.
	dbx, err := database.Migrate(l, con, migration.Changesets)
	if err != nil {
		l.Fatalf(err.Error())
	}

	return database.New(dbx, con.Name)
}

// Services sets up the production services.
func Services(l logger.ILog, dbx *database.DBW, mocker *mock.Mocker) endpoint.Core {
	// FIXME: This needs to be loaded from a config.
	secret := "TA8tALZAvLVLo4ToI44xF/nF6IyrRNOR6HSfpno/81M="

	// Configure the services.
	mux := router.New()

	// Return a new core.
	core := endpoint.NewCore(
		l,
		mux,
		bind.New(mux),
		response.New(),
		webtoken.New([]byte(secret)),
		passhash.New(),
		store.LoadFactory(mocker,
			dbx,
			query.New(mocker, dbx),
		),
		requestcontext.New(mocker),
	)

	// Set up the router.
	SetupRouter(l, mux)

	return core
}
