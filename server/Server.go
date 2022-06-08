package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"taal-client/client"
	"taal-client/console"
	"taal-client/settings"
	"taal-client/utils"

	"github.com/bitcoinsv/bsvd/bsvec"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/libsv/go-bt"
	"github.com/pkg/errors"
)

//go:generate moq -pkg service_test -out repository_mock_test.go . Repository
type Repository interface {
	InsertKey(ctx context.Context, key Key) error
	GetKey(ctx context.Context, apiKey string) (Key, error)
	GetAllKeys(ctx context.Context) ([]Key, error)
	InsertTransaction(ctx context.Context, tx Transaction) error
	GetAllTransactions(ctx context.Context, hoursBack int) ([]Transaction, error)
	Health(ctx context.Context) error
}

type Server struct {
	server     *echo.Echo
	address    string
	taal       *client.Client
	repository Repository
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

func New(address string, taal *client.Client, repo Repository) Server {
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
			"X-Mode",
			"X-Key",
			"Filename",
		},
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if settings.GetBool("debugServer") {
				req := c.Request()
				log.Printf("Server [%s]: %s %s\n", utils.GetSourceIP(req), req.Method, req.URL)
			}
			return next(c)
		}
	})

	s := Server{
		server:     e,
		address:    address,
		taal:       taal,
		repository: repo,
	}

	group := e.Group("/api/v1")

	// Deprecated: Endpoint paths don't adhere to REST design principles because they are verbs
	group.POST("/write", s.write)
	group.GET("/read/:txid", s.read)
	group.POST("/register/:apikey", s.register)

	group.DELETE("/apikeys/:apikey", s.revoke)
	group.POST("/apikeys/:apikey", s.register)
	group.GET("/apikeys", s.getApiKeys)

	group.GET("/settings", s.getSettings)
	group.PUT("/settings", s.putSetting)

	group.POST("/transactions", s.write)
	group.GET("/transactions/:txid", s.read)
	group.GET("/transactions/info", s.getTransactions)

	group.GET("/health", func(c echo.Context) error {
		err := s.repository.Health(context.Background())

		if err != nil {
			return s.sendError(c, http.StatusInternalServerError, errHealthDB, errors.Wrap(err, "failed to call database"))
		}

		return c.String(http.StatusOK, "server is running")
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

	log.Printf("INFO: Service starting on %s", s.address)
	log.Printf("INFO: Requests will be proxied to %q", s.taal.Url)

	// Try to parse the s.address.  If an error occurs, we will insert "localhost"
	_, err := url.Parse(s.address)
	if err != nil {
		log.Printf("INFO: Console available at http://localhost%s", s.address)
	} else {
		log.Printf("INFO: Console available at http://%s", s.address)
	}

	if err := s.server.Start(s.address); err != nil && err != http.ErrServerClosed {
		log.Printf("ERROR: HTTP server failed: %v", err)
		return err
	}

	return nil
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
