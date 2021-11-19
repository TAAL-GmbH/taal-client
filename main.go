package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"taal-client/client"
	"taal-client/config"
	"taal-client/server"

	"github.com/bitcoinsv/bsvd/bsvec"
)

func usage() {
	fmt.Println(`
Usage
-----
taal-client register <api-key>
  Creates a private key which is held locally, and sends the public key to Taal to be linked to an existing API key.

taal-client start
  Starts listening for requests on :9500.  This value can be changed with the LISTEN environment variable.

  All requests will be sent to https://mapi.taal.com by default unless overridden with the TAAL_URL environment variable.

Environment variables
---------------------
  LISTEN
  TAAL_URL
  TAAL_TIMEOUT

Example
-------
  LISTEN=localhost:8080 TAAL_URL=http://localhost:4000 TAAL_TIMEOUT=1m ./taal_client start
  
	`)

	os.Exit(1)
}

func start(conf *config.Config, taal *client.Client) {
	// taal-client start

	stopServer := make(chan bool, 1)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-stop
		stopServer <- true
	}()

	s := server.New(conf.ListenAddress, taal)

	if err := s.Start(stopServer); err != nil {
		log.Fatalf("app terminated with error: %v", err)
	}
}

func register(conf *config.Config, taal *client.Client) {
	// taal-client register <api-key>
	apiKey := os.Args[2]

	privateKey, err := config.GetPrivateKey(apiKey)
	if err != nil {
		if os.IsNotExist(err) {
			// The file is not there, so we can continue registration
		} else {
			fmt.Printf("failed to check apiKey existence: %s", err)
			os.Exit(1)
		}
	}

	if privateKey != nil {
		fmt.Println("apikey has already been registered")
		os.Exit(1)
	}

	privateKey, err = bsvec.NewPrivateKey(bsvec.S256())
	if err != nil {
		fmt.Printf("failed to generate new private key: %s", err)
		os.Exit(1)
	}

	h := sha256.Sum256([]byte(apiKey))

	signature, err := privateKey.Sign(h[:])
	if err != nil {
		fmt.Printf("failed to sign private key with api key: %s", err)
		os.Exit(1)
	}

	signatureStr := hex.EncodeToString(signature.Serialize())
	publicKeyStr := hex.EncodeToString(privateKey.PubKey().SerializeCompressed())

	err = taal.Register(signatureStr, publicKeyStr, apiKey)
	if err != nil {
		fmt.Printf("failed to register: %s", err)
		os.Exit(1)
	}

	err = config.StorePrivateKey(apiKey, privateKey)
	if err != nil {
		fmt.Printf("failed to write private key: %s", err)
		os.Exit(1)
	}

	fmt.Println("Registration successful")
	os.Exit(0)
}

func main() {
	conf := config.Load()

	taal := client.New(conf.TaalUrl, conf.TaalTimeOut)

	if len(os.Args) == 2 && os.Args[1] == "start" {
		start(conf, taal)
	} else if len(os.Args) == 3 && os.Args[1] == "register" {
		register(conf, taal)
	} else {
		usage()
	}
}
