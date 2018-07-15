package sessions

import (
	"github.com/labstack/echo"

	"github.com/defact/user/resources/sessions/handlers"
)

func Route(e *echo.Echo) {
	e.POST("/sessions", sessions.SessionsHandlers{}.Post)
}
