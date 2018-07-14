package users

import (
	"net/http"

	"github.com/labstack/echo"

	service "github.com/defact/user/resources/users/services"
)

type UsersHandlers struct{}

func (h UsersHandlers) Get(c echo.Context) error {
	users := service.Finder{}.List()

	return c.JSON(http.StatusOK, users)
}
