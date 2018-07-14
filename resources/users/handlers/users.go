package users

import (
	"net/http"

	"github.com/labstack/echo"

	model "github.com/defact/user/resources/users/models"
	service "github.com/defact/user/resources/users/services"
)

type UsersHandlers struct{}

func (h UsersHandlers) Get(c echo.Context) error {
	users := service.Finder{}.List()

	return c.JSON(http.StatusOK, users)
}

func (h UsersHandlers) Post(c echo.Context) error {
	u := model.User{}
	if err := c.Bind(&u); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	user := service.Creator{}.Create(u)

	return c.JSON(http.StatusCreated, user)
}
