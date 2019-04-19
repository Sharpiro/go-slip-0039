package cryptos

import "crypto/sha256"
import "crypto/hmac"

// GetHmacSha256 creates an hmac sha 256 hash
func GetHmacSha256(key, message []byte) []byte {

	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	hash := mac.Sum(nil)
	return hash
}
