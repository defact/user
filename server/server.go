package server

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/defact/user/config"
	users "github.com/defact/user/resources/users/routes"
)

func Start() {
	configuration := config.Configuration{}
	config.Load(&configuration)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	users.Route(e)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(configuration.Port)))
}
