package users

import (
	"database/sql"

	"github.com/defact/user/database"
	"github.com/defact/user/resources/users/models"
)

type Repository struct{}

func (r Repository) Get(id string) users.User {
	db := database.Instance()

	var user users.User

	err := db.QueryRow(`SELECT id, email FROM users WHERE id=$1;`, id).Scan(&user.Id, &user.Email)

	if err == sql.ErrNoRows {
		return user
	}

	if err != nil {
		panic(err)
	}

	return user
}

func (r Repository) List() []users.User {
	db := database.Instance()

	rows, err := db.Query("SELECT id, email FROM users")

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	u := []users.User{}

	for rows.Next() {
		var user users.User
		err = rows.Scan(&user.Id, &user.Email)

		if err != nil {
			panic(err)
		}

		u = append(u, user)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return u
}
