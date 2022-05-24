package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type JsonStruct struct {
	ApiKey     string `json:"apiKey"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

const keyFolder = "./keys"

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
