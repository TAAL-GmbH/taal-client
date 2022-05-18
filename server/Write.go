package server

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"taal-client/config"

	"github.com/labstack/echo/v4"
)

func (s Server) write(c echo.Context) error {
	authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
	if authHeader == "" {
		return s.sendError(c, http.StatusUnauthorized, 1, errors.New("missing authorization"))
	}

	apiKey := strings.Replace(authHeader, "Bearer ", "", 1)

	privateKey, err := config.GetPrivateKey(apiKey)
	if err != nil {
		if os.IsNotExist(err) {
			return s.sendError(c, http.StatusUnauthorized, 1, errors.New("unknown apikey"))
		}
		return s.sendError(c, http.StatusInternalServerError, 1, errors.New("failed to read apikey data"))
	}

	mimeType := c.Request().Header.Get(echo.HeaderContentType)
	if mimeType == "" {
		return s.sendError(c, http.StatusBadRequest, 10, errors.New("empty mimetype"))
	}

	if c.Request().Body == nil {
		return s.sendError(c, http.StatusBadRequest, 11, errors.New("empty body"))
	}

	reqBody, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, 12, fmt.Errorf("could not read body: %w", err))
	}
	defer func() {
		err = c.Request().Body.Close()
		if err != nil {
			log.Printf("WARN: failed to close request body: %s", err.Error())
		}
	}()

	opReturnOutput, err := buildOpReturnOutput(c.Request().Header.Get("x-tag"), mimeType, reqBody)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, 14, fmt.Errorf("failed to create op return output: %w", err))
	}

	feesRequired := true
	if c.Request().Header.Get("x-feesrequired") == "false" {
		feesRequired = false
	}

	feeTx, dataTx, err := s.taal.GetTransactionsTemplate(apiKey, len(opReturnOutput.ToBytes()), feesRequired)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, 13, fmt.Errorf("failed to get transactions: %w", err))
	}

	dataTx.Inputs[0].PreviousTxSatoshis = feeTx.GetOutputs()[0].Satoshis
	dataTx.Inputs[0].PreviousTxScript = feeTx.GetOutputs()[0].LockingScript

	dataTx.AddOutput(opReturnOutput)

	err = signTx(privateKey, dataTx)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, 15, fmt.Errorf("failed to sign transaction: %w", err))
	}

	// fmt.Printf("DATA TX:\n%x\n", dataTx.ToBytes())

	err = s.taal.SubmitTransactions(apiKey, feeTx, dataTx)
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, 16, fmt.Errorf("failed to submit transactions: %w", err))
	}

	return s.sendSuccess(c, http.StatusCreated, dataTx.GetTxID())
}
