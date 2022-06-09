package server

import (
	"database/sql"
	"io"
	"net/http"
	"strings"
	"taal-client/encryption"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (s Server) read(c echo.Context) error {
	authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
	if authHeader == "" {
		return s.sendError(c, http.StatusUnauthorized, errReadTxNotAuthorized, errors.New("missing authorization"))
	}

	apiKey := strings.Replace(authHeader, "Bearer ", "", 1)

	transactionID := c.Param("txid")

	readerCloser, contentType, err := s.taal.ReadTransaction(c.Request().Context(), apiKey, transactionID)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, errReadTx, errors.Wrap(err, "failed to get transaction"))
	}

	tx, err := s.repository.GetTransaction(c.Request().Context(), transactionID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return s.sendError(c, http.StatusBadRequest, errReadTx, errors.Wrap(err, "failed to lookup transaction"))
		}
	}

	if tx != nil && tx.Secret != "" {
		// Read all the payload, decrypt and return the plain data as a ReaderCloser
		bodyBytes, err := io.ReadAll(readerCloser)
		if err != nil {
			return s.sendError(c, http.StatusBadRequest, errReadTx, errors.Wrap(err, "failed to read encrypted data"))
		}

		plainData, err := encryption.Decrypt(bodyBytes, []byte(tx.Secret))
		if err != nil {
			return s.sendError(c, http.StatusBadRequest, errReadTx, errors.Wrap(err, "failed to decrypt data"))
		}

		return c.Blob(http.StatusOK, contentType, plainData)

	}

	return c.Stream(http.StatusOK, contentType, readerCloser)
}
