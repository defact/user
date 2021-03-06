package users

import (
	"github.com/labstack/echo"

	"github.com/defact/user/resources/users/handlers"
)

func Route(e *echo.Echo) {
	e.POST("/users", users.UsersHandlers{}.Post)
	e.GET("/users", users.UsersHandlers{}.Get)
	e.GET("/users/:id", users.UserHandlers{}.Get)
}
