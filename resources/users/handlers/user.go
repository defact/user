package users

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/defact/user/json"
	service "github.com/defact/user/resources/users/services"
)

type UserHandlers struct{}

func (h UserHandlers) Get(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := service.Finder{}.Get(id)

	if user.Id == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Not found")
	}

	return c.JSON(http.StatusOK, json.Wrap("user", user))
}
