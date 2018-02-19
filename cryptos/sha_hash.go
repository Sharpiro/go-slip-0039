package cryptos

import (
	"crypto/sha256"
)

func GetSha256(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)
	hash := hasher.Sum(nil)
	return hash
}
