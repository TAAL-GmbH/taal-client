package server

import (
	"context"
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (s Server) write(c echo.Context) error {
	authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
	if authHeader == "" {
		return s.sendError(c, http.StatusUnauthorized, errWriteTxNotAuthorized, errors.New("missing authorization"))
	}

	apiKey := strings.Replace(authHeader, "Bearer ", "", 1)
	ctx := context.Background()
	privateKey, err := s.repository.GetKey(ctx, apiKey)
	if err != nil {
		if err == sql.ErrNoRows {
			return s.sendError(c, http.StatusUnauthorized, errWriteApiKeyUnknown, errors.New("unknown apikey"))
		}
		return s.sendError(c, http.StatusInternalServerError, errWriteFailedToReadApiKey, errors.New("failed to read apikey data"))
	}

	mimeType := c.Request().Header.Get(echo.HeaderContentType)
	if mimeType == "" {
		return s.sendError(c, http.StatusBadRequest, errWriteEmptyMimeType, errors.New("empty mimetype"))
	}

	if c.Request().Body == nil {
		return s.sendError(c, http.StatusBadRequest, errWriteEmptyBody, errors.New("empty body"))
	}

	reqBody, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, errWriteCouldNotReadBody, errors.Wrap(err, "could not read body"))
	}
	defer func() {
		err = c.Request().Body.Close()
		if err != nil {
			log.Printf("WARN: failed to close request body: %s", err.Error())
		}
	}()

	opReturnOutput, err := buildOpReturnOutput(c.Request().Header.Get("x-tag"), mimeType, reqBody)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, errWriteFailedToReturnOpReturnOutput, errors.Wrap(err, "failed to create op return output"))
	}

	feesRequired := true
	if c.Request().Header.Get("x-feesrequired") == "false" {
		feesRequired = false
	}

	feeTx, dataTx, err := s.taal.GetTransactionsTemplate(apiKey, len(opReturnOutput.ToBytes()), feesRequired)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, errWriteFailedToGetTxs, errors.Wrap(err, "failed to get transactions"))
	}

	dataTx.Inputs[0].PreviousTxSatoshis = feeTx.GetOutputs()[0].Satoshis
	dataTx.Inputs[0].PreviousTxScript = feeTx.GetOutputs()[0].LockingScript

	dataTx.AddOutput(opReturnOutput)

	privateKeyCurve, err := GetPrivateKey(privateKey.PrivateKey)
	err = signTx(privateKeyCurve, dataTx)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, errWriteFailedToSignTx, errors.Wrap(err, "failed to sign transaction"))
	}

	err = s.taal.SubmitTransactions(apiKey, feeTx, dataTx)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, errWriteFailedToSubmitTxs, errors.Wrap(err, "failed to submit transactions"))
	}

	ctx := context.Background()
	tx := Transaction{ID: dataTx.GetTxID(), ApiKey: apiKey}
	err = s.repository.InsertTransaction(ctx, tx)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, errWriteInsertTransaction, errors.Wrap(err, "failed to write transaction to DB"))
	}

	return s.sendSuccess(c, http.StatusCreated, dataTx.GetTxID())
}
