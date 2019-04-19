package cryptos

import (
	"encoding/hex"
	"testing"
)

func TestGetHmacSha256(tester *testing.T) {
	expectedHashHex := "5031fe3d989c6d1537a013fa6e739da23463fdaec3b70137d828e36ace221bd0"
	key := []byte("key")
	message := []byte("data")
	hash := GetHmacSha256(key, message)
	actualHashHex := hex.EncodeToString(hash)
	if expectedHashHex != actualHashHex {
		tester.Error("hashes did not match")
	}
}
