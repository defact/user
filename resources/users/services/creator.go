package users

import (
	model "github.com/defact/user/resources/users/models"
	"github.com/defact/user/resources/users/repository"
)

type Creator struct{}

func (c Creator) Create(user model.User) model.User {
	user.HashedPassword = model.HashAndSalt(user.Password)
	return users.Repository{}.Create(user)
}
