package database

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
		
	"github.com/defact/user/migrations"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "johnny"
  password = "johnny"
  dbname   = "defact"
)

var database *sql.DB

func Initialize() {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
		
	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	database = db

	migrations.Migrate()
}

func Instance() *sql.DB {
	return database
}