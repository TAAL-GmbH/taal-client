package server

import (
	"context"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type Granularity int64

const (
	None       Granularity = 0
	Minute                 = 1
	TenMinutes             = 2
	Hour                   = 3
	TenHour                = 4
	Day                    = 5
)

const defaultHoursBack = 720 // 30 * 24 = 720 hours in 30 days

func (s Server) getTransactionInfo(c echo.Context) error {
	ctx := context.Background()

	hoursBackParam := c.QueryParam("hours_back")
	var hoursBack int = defaultHoursBack

	if hoursBackParam != "" {
		hoursBackParsed, err := strconv.Atoi(hoursBackParam)
		if err != nil {
			s.sendError(c, http.StatusBadRequest, errGetTransactionInfoHoursBackMustBeInteger, errors.Wrapf(err, "given value for hours back parameter is %s, but must be integer", hoursBackParam))
		}
		if hoursBackParsed < 0 {
			s.sendError(c, http.StatusBadRequest, errGetTransactionInfoHoursBackMustBePositive, errors.Wrapf(err, "given value for hours back parameter is %d, but must be positive", hoursBackParsed))
		}
		hoursBack = hoursBackParsed
	}

	to := time.Now()
	from := to.Add(-1 * time.Duration(hoursBack) * time.Hour).UTC()

	granularity, err := getGranularity(from, to)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, errGetTransactionInfoGetGranularity, errors.Wrapf(err, "failed getting granularity for inputs from: %s, to: %s", from, to))
	}
	txs, err := s.repository.GetTransactionInfo(ctx, from, to, granularity)
	if err != nil {
		return s.sendError(c, http.StatusInternalServerError, errGetTransactionInfo, errors.Wrap(err, "failed to get transaction information"))
	}

	txsWithGaps, err := fillDates(from, to, granularity, txs)
	if err != nil {
		return s.sendError(c, http.StatusInternalServerError, errGetTransactionInfoFillDates, errors.Wrap(err, "failed filling dates"))
	}
	sort.SliceStable(txsWithGaps, func(i, j int) bool { return txsWithGaps[i].Timestamp.Before(txsWithGaps[j].Timestamp) })

	return c.JSON(http.StatusOK, TransactionInfos{Transactions: txsWithGaps})
}

func getGranularity(from, to time.Time) (Granularity, error) {
	if from.After(to) {
		return 0, errors.New("From must be before To")
	}

	difference := to.Sub(from)
	fuzziness := time.Duration(1 * time.Minute)

	if difference == 0 {
		return None, errors.New("From and To must be different")
	}

	if difference < 1*time.Hour+fuzziness {
		return Minute, nil // 1 minute
	}

	if difference < 72*time.Hour+fuzziness {
		return Hour, nil // 1 hour
	}

	return Day, nil // 1 day
}

func fillDates(from time.Time, to time.Time, granularity Granularity, stats []TransactionInfo) ([]TransactionInfo, error) {

	var step time.Duration

	switch granularity {
	case Minute:
		step = time.Minute
	case TenMinutes:
		step = time.Minute * 10
	case Hour:
		step = time.Hour
	case TenHour:
		step = time.Hour * 10
	case Day:
		step = time.Hour * 24
	default:
		return nil, errors.New("invalid granularity")
	}

	from = from.UTC().Truncate(step)

	for d := from; d.Before(to); d = d.Add(step) {
		if !dateExists(stats, d) {
			stats = append(stats, TransactionInfo{
				Timestamp: d,
			})
		}
	}

	return stats, nil
}

func dateExists(stats []TransactionInfo, ts time.Time) bool {
	for _, s := range stats {
		if s.Timestamp == ts {
			return true
		}
	}
	return false
}
