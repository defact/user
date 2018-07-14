package server

import (
	users "github.com/defact/user/resources/users/routes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	users.Route(e)

	e.Logger.Fatal(e.Start(":5000"))
}
