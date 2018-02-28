package cryptos

import (
	"crypto/sha256"
)

// GetSha256 creates a sha 256 hash based off the given data
func GetSha256(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)
	hash := hasher.Sum(nil)
	return hash
}
