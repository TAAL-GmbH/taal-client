package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (s Server) getTransactions(c echo.Context) error {
	ctx := context.Background()
	txs, err := s.repository.GetAllTransactions(ctx)
	if err != nil {
		return s.sendError(c, http.StatusInternalServerError, errGetTransactions, errors.Wrap(err, "failed to get transaction information"))
	}

	return c.JSON(http.StatusOK, Transactions{Transactions: txs})
}
