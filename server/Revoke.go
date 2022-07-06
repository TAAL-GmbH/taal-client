package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (s Server) deactivate(c echo.Context) error {
	apiKey := c.Param("apikey")

	ctx := context.Background()
	err := s.repository.DeactivateKey(ctx, apiKey)
	if err != nil {
		s.sendError(c, http.StatusInternalServerError, errAPIKeysRevoke, errors.Wrap(err, "revoke failed"))
	}

	return c.JSON(http.StatusOK, "")
}
