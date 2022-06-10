package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

func pad(buf []byte, size int) ([]byte, error) {
	bufLen := len(buf)
	padLen := size - bufLen%size
	padded := make([]byte, bufLen+padLen)
	copy(padded, buf)
	for i := 0; i < padLen; i++ {
		padded[bufLen+i] = byte(padLen)
	}
	return padded, nil
}

func unPad(padded []byte, size int) ([]byte, error) {
	if len(padded)%size != 0 {
		return nil, errors.New("pkcs7: Padded value wasn't in correct size.")
	}

	bufLen := len(padded) - int(padded[len(padded)-1])
	buf := make([]byte, bufLen)
	copy(buf, padded[:bufLen])
	return buf, nil
}

func Encrypt(plainData, key []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(key)
	keyHash := h.Sum(nil)

	plainData, err := pad(plainData, aes.BlockSize)
	if err != nil {
		return nil, fmt.Errorf(`plainText: "%s" has error`, plainData)
	}

	if len(plainData)%aes.BlockSize != 0 {
		err := fmt.Errorf(`plainText: "%s" has the wrong block size`, plainData)
		return nil, err
	}

	block, err := aes.NewCipher(keyHash)
	if err != nil {
		return nil, err
	}

	cipherData := make([]byte, aes.BlockSize+len(plainData))
	iv := cipherData[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherData[aes.BlockSize:], plainData)

	return cipherData, nil
}

func Decrypt(cipherData, key []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(key)
	keyHash := h.Sum(nil)

	block, err := aes.NewCipher(keyHash)
	if err != nil {
		return nil, err
	}

	if len(cipherData) < aes.BlockSize {
		return nil, errors.New("cipherData too short")
	}

	iv := cipherData[:aes.BlockSize]
	cipherData = cipherData[aes.BlockSize:]

	if len(cipherData)%aes.BlockSize != 0 {
		return nil, errors.New("cipherData is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherData, cipherData)

	return unPad(cipherData, aes.BlockSize)
}
