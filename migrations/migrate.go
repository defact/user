package migrations

import (
			"github.com/golang-migrate/migrate"
	  _	"github.com/golang-migrate/migrate/database/postgres"
		_ "github.com/golang-migrate/migrate/source/file"
)

func Migrate() {
	m, err := migrate.New("file://migrations", "postgres://johnny:johnny@localhost:5432/defact?sslmode=disable")
	if err != nil {
		panic(err);
	}
	m.Up()
}