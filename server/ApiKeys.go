package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (s Server) getApiKeys(c echo.Context) error {
	ctx := context.Background()
	keys, err := s.repository.GetAllKeys(ctx)
	if err != nil {
		return s.sendError(c, http.StatusInternalServerError, errAPIKeysFailedToGetKeys, errors.Wrap(err, "failed to get api keys"))
	}

	return c.JSON(http.StatusOK, Keys{Keys: keys})
}
