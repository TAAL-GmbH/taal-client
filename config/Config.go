package config

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/bitcoinsv/bsvd/bsvec"
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

func GetPrivateKey(apiKey string) (*bsvec.PrivateKey, error) {
	f, err := os.Open(fmt.Sprintf("%s/%s.json", keyFolder, apiKey))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var s JsonStruct
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, err
	}

	privateKeyDecoded, err := hex.DecodeString(s.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to decode private key: %w", err)
	}

	privateKey, _ := bsvec.PrivKeyFromBytes(bsvec.S256(), privateKeyDecoded)

	return privateKey, nil
}

func StorePrivateKey(apiKey string, pk *bsvec.PrivateKey) error {
	// Make sure path exists
	if err := os.Mkdir(keyFolder, 0755); err != nil {
		if !os.IsExist(err) {
			return err
		}
	}

	f, err := os.OpenFile(fmt.Sprintf("%s/%s.json", keyFolder, apiKey), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")

	j := &JsonStruct{
		ApiKey:     apiKey,
		PrivateKey: hex.EncodeToString(pk.Serialize()),
		PublicKey:  hex.EncodeToString(pk.PubKey().SerializeCompressed()),
	}

	if err := enc.Encode(j); err != nil {
		return err
	}

	return nil
}

func MoveConfigKeysToArchive() error {
	return os.Rename("keys", "keys_archive")
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
