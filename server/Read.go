package server

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func (s Server) read(c echo.Context) error {
	authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
	if authHeader == "" {
		return s.sendError(c, http.StatusUnauthorized, errReadTxNotAuthorized, errors.New("missing authorization"))
	}

	apiKey := strings.Replace(authHeader, "Bearer ", "", 1)

	readerCloser, contentType, err := s.taal.ReadTransaction(apiKey, c.Param("txid"))
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, errReadTxErr, fmt.Errorf("failed to get transaction: %w", err))
	}

	return c.Stream(http.StatusOK, contentType, readerCloser)
}
