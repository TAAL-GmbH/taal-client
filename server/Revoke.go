package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (s Server) revoke(c echo.Context) error {
	return s.sendError(c, http.StatusNotImplemented, errAPIKeysRevoke, errors.New("revoke failed - functionality not yet implemented"))
}
