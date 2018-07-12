package users

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/defact/user/resources/users/services"
)

type UserHandlers struct{}

func (h UserHandlers) Get(c echo.Context) error {
	user := users.Finder{}.Get(c.Param("id"))

	if user.Id == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Not found")
	}

	return c.JSON(http.StatusOK, user)
}
