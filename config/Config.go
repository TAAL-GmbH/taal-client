package config

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/bitcoinsv/bsvd/bsvec"
)

type Config struct {
	ListenAddress string
	TaalUrl       string
	TaalTimeOut   time.Duration
}

type jsonStruct struct {
	ApiKey     string `json:"apiKey"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

const keyFolder = "./keys"

func Load() *Config {
	config := &Config{
		ListenAddress: ":9500",
		TaalUrl:       "https://mapi.taal.com",
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

	return config
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

	var s jsonStruct
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

	j := &jsonStruct{
		ApiKey:     apiKey,
		PrivateKey: hex.EncodeToString(pk.Serialize()),
		PublicKey:  hex.EncodeToString(pk.PubKey().SerializeCompressed()),
	}

	if err := enc.Encode(j); err != nil {
		return err
	}

	return nil
}
