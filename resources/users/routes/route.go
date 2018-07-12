package users

import (
	"github.com/labstack/echo"

	"github.com/defact/user/resources/users/handlers"
)

func Route(e *echo.Echo) {
	e.GET("/users", users.UsersHandlers{}.Get)
	e.GET("/users/:id", users.UserHandlers{}.Get)
}
