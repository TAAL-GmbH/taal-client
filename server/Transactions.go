package server

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

const defaultHoursBack = 720 // 30 * 24 = 720 hours in 30 days

func (s Server) getTransactions(c echo.Context) error {
	ctx := context.Background()

	hoursBackParam := c.QueryParam("hours_back")
	var hoursBack int = defaultHoursBack

	if hoursBackParam != "" {
		hoursBackParsed, err := strconv.Atoi(hoursBackParam)
		if err != nil {
			s.sendError(c, http.StatusBadRequest, errGetTransactionsHoursBackMustBeInteger, errors.Wrapf(err, "given value for hours back parameter is %s, but must be integer", hoursBackParam))
		}
		if hoursBackParsed < 0 {
			s.sendError(c, http.StatusBadRequest, errGetTransactionsHoursBackMustBePositive, errors.Wrapf(err, "given value for hours back parameter is %d, but must be positive", hoursBackParsed))
		}
		hoursBack = hoursBackParsed
	}

	txs, err := s.repository.GetAllTransactions(ctx, hoursBack)
	if err != nil {
		return s.sendError(c, http.StatusInternalServerError, errGetTransactions, errors.Wrap(err, "failed to get transaction information"))
	}

	return c.JSON(http.StatusOK, Transactions{Transactions: txs})
}
