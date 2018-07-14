package migrations

import (
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

func Migrate(path, conn string) {
	m, err := migrate.New(path, conn)
	if err != nil {
		panic(err)
	}
	m.Up()
}
