package server

import (
	"encoding/hex"

	"github.com/bitcoinsv/bsvd/bsvec"
	"github.com/bitcoinsv/bsvd/chaincfg"
	"github.com/bitcoinsv/bsvutil"
	"github.com/pkg/errors"
)

type Key struct {
	ApiKey     string  `db:"api_key" json:"api_key"`
	PublicKey  string  `db:"public_key" json:"public_key"`
	PrivateKey string  `db:"private_key" json:"-"`
	Address    string  `db:"address" json:"address"`
	CreatedAt  string  `db:"created_at" json:"createdAt"`
	RevokedAt  *string `db:"revoked_at" json:"revokedAt"`
}

type Keys struct {
	Keys []Key `json:"keys"`
}

type Transaction struct {
	ID        string `db:"id" json:"id"`
	ApiKey    string `db:"api_key" json:"api_key"`
	DataBytes int    `db:"data_bytes" json:"data_bytes"`
	CreatedAt string `db:"created_at" json:"created_at"`
	Filename  string `db:"filename" json:"filename"`
	Secret    string `db:"secret" json:"secret"`
}

type Transactions struct {
	Transactions []Transaction `json:"transactions"`
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

func GetPrivateKey(privateKey string) (*bsvec.PrivateKey, error) {
	privateKeyDecoded, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode private key")
	}

	privateKeyCurve, _ := bsvec.PrivKeyFromBytes(bsvec.S256(), privateKeyDecoded)

	return privateKeyCurve, nil

}
