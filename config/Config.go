package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type Config struct {
	ListenAddress string
	TaalUrl       string
	TaalTimeOut   time.Duration
	ConnectToDB   bool
	Host          string
	DBName        string
	Username      string
	Password      string
	Port          int
}

type JsonStruct struct {
	ApiKey     string `json:"apiKey"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

const keyFolder = "./keys"

func Load() (*Config, error) {
	config := &Config{
		ListenAddress: ":9500",
		TaalUrl:       "https://tapi.taal.com",
		TaalTimeOut:   5 * time.Second,
	}

	l := os.Getenv("LISTEN")
	if l != "" {
		config.ListenAddress = l
	}

	url := os.Getenv("TAAL_URL")
	if url != "" {
		config.TaalUrl = url
	}

	timeout := os.Getenv("TAAL_TIMEOUT")
	if timeout != "" {
		d, err := time.ParseDuration(timeout)
		if err != nil {
			log.Printf("%s is not a valid duration [%v]", timeout, err)
		} else {
			config.TaalTimeOut = d
		}
	}

	connectToDB := os.Getenv("CONNECT_TO_DB")
	config.ConnectToDB = connectToDB == "1"
	if !config.ConnectToDB {
		return config, nil
	}

	host := os.Getenv("HOST")
	if host != "" {
		config.Host = host
	}

	dbName := os.Getenv("DBNAME")
	if dbName != "" {
		config.DBName = dbName
	}

	username := os.Getenv("USERNAME")
	if username != "" {
		config.Username = username
	}

	password := os.Getenv("PASSWORD")
	if password != "" {
		config.Password = password
	}

	portstr := os.Getenv("PORT")
	if portstr != "" {
		port, err := strconv.Atoi(portstr)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to parse port %s to integer", portstr)
		}
		config.Port = port
	}

	return config, nil
}

func GetKeysFromJson() ([]JsonStruct, error) {
	keys := make([]JsonStruct, 0)

	items, _ := ioutil.ReadDir(keyFolder)
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		filepath := fmt.Sprintf("%s/%s", keyFolder, item.Name())
		f, err := os.Open(filepath)
		if err != nil {
			return nil, err
		}

		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}

		var key JsonStruct
		if err := json.Unmarshal(b, &key); err != nil {
			return nil, err
		}

		keys = append(keys, key)
	}

	return keys, nil
}
