package encryption

import (
	"bytes"
	"testing"
)

func TestEncDec(t *testing.T) {
	key := []byte("secret")
	data := []byte("hello world")

	cipherData, err := Encrypt(data, key)
	if err != nil {
		t.Error(err)
	}

	plainData, err := Decrypt(cipherData, key)
	if err != nil {
		t.Error(err)
	}

	if bytes.Compare(plainData, data) != 0 {
		t.Errorf("Expected %q, got %q", string(data), string(plainData))
	}

	t.Log(string(plainData))
}
