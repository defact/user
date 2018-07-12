package users

import (
	model "github.com/defact/user/resources/users/models"
	"github.com/defact/user/resources/users/repository"
)

type Finder struct{}

func (f Finder) Get(id string) model.User {
	return users.Repository{}.Get(id)
}

func (f Finder) List() []model.User {
	return users.Repository{}.List()
}
