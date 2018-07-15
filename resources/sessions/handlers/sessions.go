package sessions

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/defact/user/json"
	model "github.com/defact/user/resources/sessions/models"
	service "github.com/defact/user/resources/sessions/services"
)

type SessionsHandlers struct{}

func (h SessionsHandlers) Post(c echo.Context) error {
	auth := model.Authentication{}
	if err := c.Bind(&auth); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	token, err := service.Authenticator{}.Authenticate(auth)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, json.Wrap("token", token))
}
