package cryptos

import (
	"bytes"
	"testing"
)

func TestGCMEncryption(tester *testing.T) {
	key := [32]byte{}
	expectedPlainBytes := []byte("plainText")

	actualCryptoBytes, err := EncryptGCM(expectedPlainBytes, &key)
	actualPlainBytes, err := DecryptGCM(actualCryptoBytes, &key)

	if err != nil {
		tester.Error("encryption error")
	}

	if !bytes.Equal(expectedPlainBytes, actualPlainBytes) {
		tester.Error("hash mismatch")
	}
}
