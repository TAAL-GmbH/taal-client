package server

import (
	"net/http"
	"taal-client/settings"

	"github.com/labstack/echo/v4"
)

func (s Server) getSettings(c echo.Context) error {
	return c.JSON(http.StatusOK, settings.Get())
}
