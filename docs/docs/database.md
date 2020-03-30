---
id: database
title: Database
---

## Docker

To get a MySQL running on your computer, we recommend you use [Docker](https://www.docker.com/). You need to make sure it's installed. Docker makes it easy to run software on your computer without having to install to. Docker also makes it easy reset or clear your environment so you can bring up new services like MySQL quickly and easily. This makes testing or working on multiple projects very easy.

You can also install MySQL using `brew install mysql@5.7` on MacOS or install from their [website](https://dev.mysql.com/downloads/installer/), but many of the commands below don't apply.

We'll be using the [standard MySQL Docker image](https://hub.docker.com/_/mysql) which you can read about on their Docker Hub website.

### Create Container

Create a new MySQL 5.7 database container. The username will be **root** and the password will be **password**. You can also connect to it via **127.0.0.1:3306** using a tool like [Sequel Pro](https://www.sequelpro.com/) or the [MySQL Workbench](https://dev.mysql.com/downloads/workbench/).

```bash
# Makefile
make db-init

# Manual
docker run -d --name=gomithrilapp_db_1 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password mysql:5.7
```

### Stop Container

Stop a running MySQL database container. This will preserve the database data.

```bash
# Makefile
make db-stop

# Manual
docker stop gomithrilapp_db_1
```

### Start a Stopped Container

Start a stopped MySQL database container.

```bash
# Makefile
make db-start

# Manual
docker start gomithrilapp_db_1
```

### Reset Container

Drop the database table and then re-run migrations.

```bash
# Makefile
make db-reset

# Manual
docker exec gomithrilapp_db_1 sh -c "exec mysql -h 127.0.0.1 -uroot -ppassword -e 'DROP DATABASE IF EXISTS main;'"
docker exec gomithrilapp_db_1 sh -c "exec mysql -h 127.0.0.1 -uroot -ppassword -e 'CREATE DATABASE IF NOT EXISTS main DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;'"
go run ${GOPATH}/src/app/api/cmd/dbmigrate/main.go
```

### Delete the Container

Stop and then remove the container - all database data will be deleted as well.

```bash
# Makefile
make db-rm

# Manual
docker rm -f gomithrilapp_db_1
```

## Migrations

MySQL migrations are performed at boot by [Rove](https://github.com/josephspurrier/rove), a tool very similiar to Liquibase. The migrations are run when API starts up - they are located [here](https://github.com/josephspurrier/gomithrilapp/blob/master/src/app/api/migration/changeset.go).

Database migrations are a great way to manage incremental database changes. The migration state is stored in the same database and recorded in the **rovechangelog** table.

The primary motivation behind [Rove](https://github.com/josephspurrier/rove) tool was to provide a simple and quick Go (instead of Java) based database migration tool that allows loading migrations from anywhere, including from inside your code so you can distribute single binary applications. You can write the migration and rollback SQL, then Rove will apply it for you properly.

### How do migrations work?

Database migrations are necessary when you make changes to your applications. You may need a new table or a new column so you have to write the SQL commands to make the changes. The tricky piece is when you perform an upgrade, how do you manage which SQL queries will run? Do you run all of them again and then the new ones after? Or is there an easy way to track which queries have been run so you only run new ones? What if you have to rollback your database because of a feature that was released too early and is causing problem? How do you manage those queries? You can definitely write your own code to manage the migration process, but Rove makes the process much easier for you. You also don't have to convert your SQL code to a another format like JSON or XML, you can just add a few comments around it and Rove will handle the rest.

###  How does Rove work?

You'll need to write your changes queries and rollback queries in migration files. These are plain SQL files that can be imported directly into MySQL. Rove just uses comments to help break them into smaller manageable pieces. When you run tell Rove to apply your changes, a table called **rovechangelog** is created in the database to track which changesets have been applied and metadata about them. The tool will ensure no changes have been made to the existing changesets that are already in the database. Changeset checksums are then compared against the changelog table checksums. Any new changesets that are not in the changelog are applied to the database and then a new record is inserted into the changelog for each changeset. Rove supports labeling changesets with a **tag** as well as rolling back to specific tags.

### Rove vs Liquibase

Rove and Liquibase use different changelog tables. Rove includes MySQL out of the box, but it supports adding your own adapters to work with any type of data storage. The Rove changesets can use a very similar plain SQL (no XML or JSON) file format for simplicity and portability. For the most teams, you'll be able use your existing SQL migration files with Rove without making any changes.

To assist with switching from Liquibase to Rove, you can use the CLI tool with the `rove convert` argument to convert a Liquibase **DATABASECHANGELOG** table to a Rove **rovechangelog** table. If you don't run the `rove convert` command first on a database that was originally managed by Liquibase, Rove will try to rerun the same migrations over again if you use the same migration files. The tools use different changelog table names, table schemas, and use different methods for calculating their checksums.

### Sample Migration File from Rove

Below is an example of what a migration file looks likes. Writing the files like this provides you with a few benefits:

- Easily see who created the database migration
- Description of what the changeset does.
- Rollback command to reverse a migration if you need to move back to an earlier version of the software if there is a bug.
- Break your changesets into multiple files so it's easy to separate out for instance dev vs prod changesets.

```sql
--changeset josephspurrier:1
--description Create the user status table.
--description Set deleted_at as a timestamp.
CREATE TABLE user_status (
    id TINYINT(1) UNSIGNED NOT NULL AUTO_INCREMENT,
    
    status VARCHAR(25) NOT NULL,
    
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    
    PRIMARY KEY (id)
);
--rollback DROP TABLE user_status;

--include anotherfile.sql

--changeset josephspurrier:2
INSERT INTO user_status (id, status, created_at, updated_at, deleted) VALUES
(1, 'active',   CURRENT_TIMESTAMP,  CURRENT_TIMESTAMP,  0),
(2, 'inactive', CURRENT_TIMESTAMP,  CURRENT_TIMESTAMP,  0);
--rollback TRUNCATE TABLE user_status;
```

You can read more about [Rove on GitHub](https://github.com/josephspurrier/rove).

## Changesets

A changeset is one or more SQL queries that Rove will apply to the database. Each changeset should have the SQL queries to run during the migration and then rollback SQL queries.

The current changesets are here: [src/app/api/migration/changeset.go](https://github.com/josephspurrier/gomithrilapp/blob/master/src/app/api/migration/changeset.go)

### When do migrations apply?

To simply the setup process, the migrations are run on the database when the API runs up.

```go
func main() {
	// Create the logger.
	l := logger.New(log.New(os.Stderr, "", log.Lshortfile))

	// Load the environment variables.
	settings := config.LoadEnv(l, "")

	// Setup the services.
	core := config.Services(l, settings, config.Database(l), requestcontext.New(), nil)
	config.LoadRoutes(core)

	// Start the web server.
	l.Printf("Server started.")
	err := http.ListenAndServe(fmt.Sprintf(":%v", settings.Port), config.Middleware(core))
	if err != nil {
		l.Printf(err.Error())
	}
}
```

View [src/app/api/cmd/api/main.go](https://github.com/josephspurrier/gomithrilapp/blob/master/src/app/api/cmd/api/main.go).

The database is setup and migrations are called here: `config.Database(l)`.

```go
core := config.Services(l, settings, config.Database(l), requestcontext.New(), nil)
```

The `config.Database()` function connects to the database, then runs the migrations on the database.

```go
// Database migrates the database and then returns the database connection.
func Database(l logger.ILog) *database.DBW {
	// If the host env var is set, use it.
	host := os.Getenv("MYSQL_HOST")
	if len(host) == 0 {
		host = "127.0.0.1"
	}

	// If the password env var is set, use it.
	password := os.Getenv("MYSQL_ROOT_PASSWORD")

	// Set the database connection information.
	con := &mysql.Connection{
		Hostname:  host,
		Username:  "root",
		Password:  password,
		Name:      "main",
		Port:      3306,
		Parameter: "collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true",
	}

	// Migrate the database.
	dbx, err := database.Migrate(l, con, migration.Changesets)
	if err != nil {
		l.Fatalf(err.Error())
	}

	return database.New(dbx, con.Name)
}
```

View [src/app/api/config/database.go](https://github.com/josephspurrier/gomithrilapp/blob/master/src/app/api/config/database.go).

Below is the code that references the changets and applies them - if there is an error applying them, then the application will shutdown to prevent the database from being in an invalid state.

```go
// Migrate the database.
dbx, err := database.Migrate(l, con, migration.Changesets)
if err != nil {
    l.Fatalf(err.Error())
}
```

### Add a Changeset

To add a new changeset, open up the [src/app/api/migration/changeset.go](https://github.com/josephspurrier/gomithrilapp/blob/master/src/app/api/migration/changeset.go) file and add a new changeset to the bottom.

Notice `--changeset josephspurrier:2` - each changeset must be unique so you should increment the number at the end by one when you add your changeset.

You should always include a rollback - the rollback can span multiple lines like this:

```sql
--rollback DELETE FROM user_status WHERE id=2;
--rollback DELETE FROM user_status WHERE id=1;
```

Once a migration has been applied to a system, you cannot change the existing migration. This is important - if you can a migration that has already been applied, you will receive an error stating the checksum doesn't match. Changing an existing changeset could break functionality so it's not supported.