package server

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/defact/user/config"
	sessions "github.com/defact/user/resources/sessions/routes"
	users "github.com/defact/user/resources/users/routes"
)

func Start() {
	configuration := config.Configuration{}
	config.Load(&configuration)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	// e.Use(middleware.JWT([]byte(configuration.Secret)))

	sessions.Route(e)
	users.Route(e)

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(configuration.Port)))
}
