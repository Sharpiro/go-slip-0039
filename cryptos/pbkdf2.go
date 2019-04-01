package cryptos

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

// CreatePbkdf2Hash creates a hash based off a given secret and an optional passphrase
func CreatePbkdf2Hash(password, salt []byte) []byte {
	const iterations int = 20000
	const byteLength int = 32

	hash := pbkdf2.Key(password, salt, iterations, byteLength, sha256.New)
	return hash
}
