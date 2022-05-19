package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"taal-client/client"
	"taal-client/console"

	"github.com/bitcoinsv/bsvd/bsvec"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/libsv/go-bt"
)

//go:generate moq -pkg service_test -out repository_mock_test.go . Repository
type Repository interface {
	InsertKey(ctx context.Context, key Key) error
	GetKey(ctx context.Context, apiKey string) (Key, error)
	GetAllKeys(ctx context.Context) ([]Key, error)
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
		},
	}))

	s := Server{
		server:     e,
		address:    address,
		taal:       taal,
		repository: repo,
	}

	group := e.Group("/api/v1")
	group.POST("/write", s.write)
	group.GET("/read/:txid", s.read)
	group.POST("/register/:apikey", s.Register)
	group.GET("/apikeys", s.GetApiKeys)

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
