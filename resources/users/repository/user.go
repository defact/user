package users

import (
	"database/sql"
	"strconv"

	"github.com/defact/user/database"
	"github.com/defact/user/resources/users/models"
)

type Repository struct{}

func (r Repository) Create(user users.User) users.User {
	db := database.Instance()

	id := 0
	query := `INSERT INTO users (email, hashed_password) VALUES ($1, $2) RETURNING id;`

	err := db.QueryRow(query, user.Email, user.HashedPassword).Scan(&id)

	if err != nil {
		panic(err)
	}

	return r.Get(id)
}

func get(query, param string) users.User {
	db := database.Instance()

	user := users.User{}

	err := db.QueryRow(query, param).Scan(&user.Id, &user.Email, &user.HashedPassword)

	if err == sql.ErrNoRows {
		return user
	}

	if err != nil {
		panic(err)
	}

	return user
}

func (r Repository) Get(id int) users.User {
	return get(`SELECT id, email, COALESCE(hashed_password, '') FROM users WHERE id=$1;`, strconv.Itoa(id))
}

func (r Repository) GetByEmail(email string) users.User {
	return get(`SELECT id, email, COALESCE(hashed_password, '') FROM users WHERE email=$1;`, email)
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
