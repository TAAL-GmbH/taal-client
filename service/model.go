package service

import (
	"encoding/hex"
	"taal-client/config"

	"github.com/bitcoinsv/bsvd/bsvec"
	"github.com/bitcoinsv/bsvd/chaincfg"
	"github.com/bitcoinsv/bsvutil"
	"github.com/pkg/errors"
)

type Key struct {
	ApiKey     string `db:"api_key"`
	PublicKey  string `db:"public_key"`
	PrivateKey string `db:"private_key"`
	Address    string `db:"address"`
}

func GetKeyFromPrivateKey(apiKey string, pk *bsvec.PrivateKey) (Key, error) {
	privateKey := hex.EncodeToString(pk.Serialize())
	publicKey := hex.EncodeToString(pk.PubKey().SerializeCompressed())

	pubKeyAddress, err := bsvutil.NewAddressPubKey([]byte(publicKey), &chaincfg.MainNetParams)
	if err != nil {
		return Key{}, errors.Wrap(err, "unable to create address")
	}

	key := Key{
		ApiKey:     apiKey,
		PublicKey:  publicKey,
		PrivateKey: privateKey,
		Address:    pubKeyAddress.EncodeAddress(),
	}

	return key, nil
}

func GetKeyFromConfigKey(configKey config.JsonStruct) (Key, error) {
	privateKey, err := GetPrivateKey(configKey.PrivateKey)
	if err != nil {
		return Key{}, nil
	}

	pubKeyAddress, err := bsvutil.NewAddressPubKey(privateKey.PubKey().SerializeCompressed(), &chaincfg.MainNetParams)
	if err != nil {
		return Key{}, errors.Wrap(err, "unable to create address")
	}

	key := Key{
		ApiKey:     configKey.ApiKey,
		PublicKey:  configKey.PublicKey,
		PrivateKey: configKey.PrivateKey,
		Address:    pubKeyAddress.EncodeAddress(),
	}
	return key, nil
}

func GetPrivateKey(pk string) (*bsvec.PrivateKey, error) {
	privateKeyDecoded, err := hex.DecodeString(pk)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode private key")
	}

	privateKey, _ := bsvec.PrivKeyFromBytes(bsvec.S256(), privateKeyDecoded)

	return privateKey, nil

}
