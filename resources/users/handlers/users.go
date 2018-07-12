package users

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/defact/user/resources/users/services"
)

type UsersHandlers struct{}

func (h UsersHandlers) Get(c echo.Context) error {
	u := users.Finder{}.List()

	return c.JSON(http.StatusOK, u)
}
