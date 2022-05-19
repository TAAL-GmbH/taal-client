package server

import (
	"encoding/hex"

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

	pubKeyAddress, err := bsvutil.NewAddressPubKey(pk.PubKey().SerializeCompressed(), &chaincfg.MainNetParams)
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

func GetPrivateKeyCurve(privateKey string) (*bsvec.PrivateKey, error) {
	privateKeyDecoded, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode private key")
	}

	privateKeyCurve, _ := bsvec.PrivKeyFromBytes(bsvec.S256(), privateKeyDecoded)

	return privateKeyCurve, nil

}