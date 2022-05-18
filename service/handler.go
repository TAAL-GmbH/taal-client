package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bitcoinsv/bsvd/bsvec"
	"github.com/pkg/errors"

	"taal-client/client"
	"taal-client/config"
	"taal-client/server"
)

type Repository interface {
	InsertKey(ctx context.Context, key Key) error
	GetKey(ctx context.Context, apiKey string) (Key, error)
}

type Service struct {
	Taal       *client.Client
	Config     *config.Config
	Repository Repository
}

func (s Service) Start() {
	// taal-client start

	stopServer := make(chan bool, 1)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-stop
		stopServer <- true
	}()

	serv := server.New(s.Config.ListenAddress, s.Taal)

	if err := serv.Start(stopServer); err != nil {
		log.Fatalf("app terminated with error: %v", err)
	}
}

func (s Service) Register(apiKey string) {
	privateKey, err := config.GetPrivateKey(apiKey)
	if err != nil {
		if os.IsNotExist(err) {
			// The file is not there, so we can continue registration
		} else {
			fmt.Printf("failed to check apiKey existence: %s", err)
			return
		}
	}

	if privateKey != nil {
		fmt.Println("apikey has already been registered")
		return
	}

	privateKey, err = bsvec.NewPrivateKey(bsvec.S256())
	if err != nil {
		fmt.Printf("failed to generate new private key: %s", err)
		return
	}

	h := sha256.Sum256([]byte(apiKey))

	signature, err := privateKey.Sign(h[:])
	if err != nil {
		fmt.Printf("failed to sign private key with api key: %s", err)
		return
	}

	signatureStr := hex.EncodeToString(signature.Serialize())
	publicKeyStr := hex.EncodeToString(privateKey.PubKey().SerializeCompressed())

	err = s.Taal.Register(signatureStr, publicKeyStr, apiKey)
	if err != nil {
		fmt.Printf("failed to register: %s", err)
		return
	}

	err = config.StorePrivateKey(apiKey, privateKey)
	if err != nil {
		fmt.Printf("failed to write private key: %s", err)
		return
	}

	fmt.Println("Registration successful")
}

func (s Service) StorePrivateKey(ctx context.Context, apiKey string, pk *bsvec.PrivateKey) error {
	key, err := GetKeyFromPrivateKey(apiKey, pk)
	if err != nil {
		errors.Wrap(err, "unable to create public key and address from private key")
	}

	err = s.Repository.InsertKey(ctx, key)
	if err != nil {
		return errors.Wrap(err, "failed to store keys")
	}

	return nil
}

func (s Service) MoveKeysToDB(ctx context.Context) error {
	// Check if there are configKeys in the configKeys folder. If yes, then store them in the DB
	configKeys, err := config.GetKeysFromJson()
	if err != nil {
		return errors.Wrap(err, "failed to get keys from json files")
	}
	if len(configKeys) > 0 {
		for _, configKey := range configKeys {

			key, err := GetKeyFromConfigKey(configKey)
			if err != nil {
				return errors.Wrap(err, "failed to get key from config key")
			}

			err = s.Repository.InsertKey(ctx, key)
			if err != nil {
				return errors.Wrap(err, "failed to insert keys from json files into DB")
			}
		}
	}

	err = config.MoveConfigKeysToArchive()
	if err != nil {
		return errors.Wrap(err, "failed to move config keys to archive folder")
	}

	return nil
}
