package server

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"taal-client/client"
	"taal-client/config"
	"taal-client/console"

	"github.com/bitcoinsv/bsvd/bsvec"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/libsv/go-bt"
)

type Server struct {
	server  *echo.Echo
	address string
	taal    *client.Client
}

type successResponse struct {
	Status int32  `json:"status"`
	Txid   string `json:"txid"`
}

type errorResponse struct {
	Status int32  `json:"status"`
	Code   int32  `json:"code"`
	Err    string `json:"error"`
}

func New(address string, taal *client.Client) Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	// e.Debug = true

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderXRequestedWith,
			"X-Tag",
		},
	}))

	s := Server{
		server:  e,
		address: address,
		taal:    taal,
	}

	group := e.Group("/api/v1")
	group.POST("/write", s.write)
	group.GET("/read/:txid", s.read)

	group.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is from the client")
	})

	e.GET("*", func(c echo.Context) error {
		return console.AppHandler(c)
	})

	return s
}

func (s Server) Start(stopServer chan bool) error {
	go func() {
		<-stopServer
		log.Println("INFO: server is shutting down")
		if err := s.server.Shutdown(context.Background()); err != nil {
			log.Printf("WARN: failed to shutdown server: %v", err)
		}
	}()

	log.Printf("INFO: starting on %s", s.address)
	log.Printf("INFO: Requests will be proxied to %q", s.taal.Url)
	log.Printf("INFO: Console available at http://localhost:9500")
	log.Printf("INFO: Example interface available at http://localhost:9500/example")

	if err := s.server.Start(s.address); err != nil && err != http.ErrServerClosed {
		log.Printf("ERROR: HTTP server failed: %v", err)
		return err
	}

	return nil
}

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

func (s Server) read(c echo.Context) error {
	authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
	if authHeader == "" {
		return s.sendError(c, http.StatusUnauthorized, 1, errors.New("missing authorization"))
	}

	apiKey := strings.Replace(authHeader, "Bearer ", "", 1)

	readerCloser, contentType, err := s.taal.ReadTransaction(apiKey, c.Param("txid"))
	if err != nil {
		return s.sendError(c, http.StatusBadRequest, 13, fmt.Errorf("failed to get transaction: %w", err))
	}

	return c.Stream(http.StatusOK, contentType, readerCloser)
}

func signTx(pk *bsvec.PrivateKey, tx *bt.Tx) error {
	signer := &bt.InternalSigner{
		PrivateKey:  pk,
		SigHashFlag: 0,
	}

	if err := tx.Sign(0, signer); err != nil {
		return fmt.Errorf("failed to sign OP return tx: %w", err)
	}

	return nil
}

func buildOpReturnOutput(appTag string, mimeType string, reqBody []byte) (*bt.Output, error) {
	var data [][]byte
	if appTag != "" {
		data = append(data, []byte(appTag))
	}

	data = append(data, []byte(mimeType))
	data = append(data, reqBody)

	return bt.NewOpReturnPartsOutput(data)
}

func (s Server) sendError(c echo.Context, status int, code int32, err error) error {
	log.Printf("WARN: %s", err.Error())

	return c.JSON(status, &errorResponse{
		Status: int32(status),
		Code:   code,
		Err:    err.Error(),
	})
}

func (s Server) sendSuccess(c echo.Context, status int, txid string) error {
	return c.JSON(status, &successResponse{
		Status: int32(status),
		Txid:   txid,
	})
}
