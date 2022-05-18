package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"taal-client/client"
	"taal-client/config"
	"taal-client/repository"
	"taal-client/service"
)

func usage() {
	fmt.Println(`
Usage
-----
taal-client register <api-key>
  Creates a private key which is held locally, and sends the public key to Taal to be linked to an existing API key.

taal-client start
  Starts listening for requests on :9500.  This value can be changed with the LISTEN environment variable.

  All requests will be sent to https://tapi.taal.com by default unless overridden with the TAAL_URL environment variable.

	DEBUG = 1 will log all transactions to console

Environment variables
---------------------
  LISTEN
  TAAL_URL
  TAAL_TIMEOUT
  DEBUG

Example
-------
  DEBUG=1 LISTEN=localhost:8080 TAAL_URL=http://localhost:4000 TAAL_TIMEOUT=1m ./taal_client start
  
	`)

	os.Exit(1)
}

func main() {
	conf := config.Load()

	db, err := repository.NewDB("file:repository.db")
	if err != nil {
		log.Fatalf("app terminated with error: %v", err)
		return
	}

	err = repository.RunMigrations(db)
	if err != nil {
		log.Fatalf("app terminated with error: %v", err)
		return
	}

	defer db.Close()

	service := service.Service{
		Taal:       client.New(conf.TaalUrl, conf.TaalTimeOut),
		Config:     conf,
		Repository: repository.NewRepository(*db),
	}

	ctx := context.Background()

	err = service.MoveKeysToDB(ctx)
	if err != nil {
		log.Fatalf("app terminated with error: %v", err)
	}

	if len(os.Args) == 2 && os.Args[1] == "start" {
		service.Start()
	} else if len(os.Args) == 3 && os.Args[1] == "register" {
		// taal-client register <api-key>
		apiKey := os.Args[2]
		service.Register(apiKey)
		os.Exit(0)
	} else {
		usage()
	}
}
