package cryptos

import (
	"encoding/hex"
	"testing"
)

func TestCreatePbkdf2Seed(tester *testing.T) {
	expectedHashHex := "e3c8ee431ab8f206718b2e6a1bee7fa6f9668d638b23e76b57110c912f8ca9e4"
	secret := []byte("defaultSecret")
	passPhrase := "defaultPassPhrase"
	hash := CreatePbkdf2Seed(secret, passPhrase)
	actualHashHex := hex.EncodeToString(hash)

	if expectedHashHex != actualHashHex {
		tester.Error("hash mismatch")
	}
}

func TestCreatePbkdf2NoPw(tester *testing.T) {
	expectedHashHex := "bf0e1a10bae97a80757501d762bd14a0981dc2be555f7b8ca0658ad31b556415"
	secret := []byte("defaultSecret2")
	var passPhrase string
	hash := CreatePbkdf2Seed(secret, passPhrase)
	actualHashHex := hex.EncodeToString(hash)

	if expectedHashHex != actualHashHex {
		tester.Error("hash mismatch")
	}
}

func TestCreatePbkdf2Bytes(tester *testing.T) {
	expectedHashHex := "c0945bb29cbb261874e77e00cec20567d22e0e5db32c20d11578aee05aae5c56"
	secret, _ := hex.DecodeString("fffaac234dac")
	var passPhrase string
	hash := CreatePbkdf2Seed(secret, passPhrase)
	actualHashHex := hex.EncodeToString(hash)

	if expectedHashHex != actualHashHex {
		tester.Error("hash mismatch")
	}
}

func TestAes(tester *testing.T) {
	EncryptAes(make([]byte, 16), []byte{})
}
