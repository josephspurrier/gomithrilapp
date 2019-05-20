package component

import (
	"os"

	"app/api"
	"app/api/internal/bind"
	"app/api/internal/response"
	"app/api/pkg/database"
	"app/api/pkg/logger"
	"app/api/pkg/mock"
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

	// Migration the database.
	dbx, err := database.Migrate(l, con, api.Changesets)
	if err != nil {
		l.Fatalf(err.Error())
	}

	// FIXME: This needs to be loaded from a config.
	secret := "TA8tALZAvLVLo4ToI44xF/nF6IyrRNOR6HSfpno/81M="

	// Configure the services.
	mocker := mock.New(false)
	mux := router.New()
	db := database.New(dbx, con.Name)
	q := query.New(mocker, db)
	binder := bind.New(mux)
	resp := response.New()
	token := webtoken.New([]byte(secret))
	pass := passhash.New()

	// Return a new core.
	core := NewCore(
		l,
		mux,
		db,
		q,
		binder,
		resp,
		token,
		pass,
		mocker,
	)

	core.Store = LoadStores(core)

	return core
}
