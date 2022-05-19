package server

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"net/http"

	"github.com/bitcoinsv/bsvd/bsvec"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (s Server) Register(c echo.Context) error {
	apiKey := c.Param("apikey")

	ctx := context.Background()
	_, err := s.repository.GetKey(ctx, apiKey)
	if err == nil {
		return s.sendError(c, http.StatusBadRequest, errRegisterApiKeyHasAlreadyBeenRegistered, errors.Wrap(err, "apikey has already been registered"))
	}
	if err != sql.ErrNoRows {
		return s.sendError(c, http.StatusInternalServerError, errRegisterFailedToCheckApiKey, errors.Wrap(err, "failed to check apiKey existence"))
	}

	privateKey, err := bsvec.NewPrivateKey(bsvec.S256())
	if err != nil {
		return s.sendError(c, http.StatusInternalServerError, errRegisterFailedToGenerateNewPrivateKey, errors.Wrap(err, "failed to generate new private key"))
	}

	h := sha256.Sum256([]byte(apiKey))

	signature, err := privateKey.Sign(h[:])
	if err != nil {
		return s.sendError(c, http.StatusInternalServerError, errRegisterFailedToSignPrivateKeywithApiKey, errors.Wrap(err, "failed to sign private key with api key"))
	}

	signatureStr := hex.EncodeToString(signature.Serialize())
	publicKeyStr := hex.EncodeToString(privateKey.PubKey().SerializeCompressed())

	err = s.taal.Register(signatureStr, publicKeyStr, apiKey)
	if err != nil {
		return s.sendError(c, http.StatusInternalServerError, errRegisterFailedToRegisterApiKey, errors.Wrap(err, "failed to register api key"))
	}

	key, err := GetKeyFromPrivateKey(apiKey, privateKey)
	if err != nil {
		return s.sendError(c, http.StatusInternalServerError, errRegisterFailedToGetKeysFromPrivateKey, errors.Wrap(err, "failed to create keys from private key"))
	}

	err = s.repository.InsertKey(ctx, key)
	if err != nil {

		return s.sendError(c, http.StatusInternalServerError, errRegisterFailedToStoreKeys, errors.Wrap(err, "failed to write private key to DB"))
	}

	return c.NoContent(http.StatusOK)
}
