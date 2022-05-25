package server

import (
	"net/http"
	"taal-client/settings"

	"github.com/labstack/echo/v4"
)

func (s Server) getSettings(c echo.Context) error {
	b, err := settings.GetJSON()
	if err != nil {
		return s.sendError(c, http.StatusInternalServerError, 22, err)
	}

	return c.Blob(http.StatusOK, echo.MIMEApplicationJSON, b)
}
