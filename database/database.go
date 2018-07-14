package database

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/defact/user/config"
	"github.com/defact/user/migrations"
)

var database *sql.DB
var configuration config.Configuration

func Initialize() {
	config.Load(&configuration)

	db, err := sql.Open("postgres", configuration.PgConnection)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	database = db

	migrations.Migrate(configuration.MigrationsPath, configuration.PgConnection)
}

func Instance() *sql.DB {
	return database
}
