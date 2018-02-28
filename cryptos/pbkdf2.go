package cryptos

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

// CreatePbkdf2Seed creates a hash based off a given secret and an optional passphrase
func CreatePbkdf2Seed(secret []byte, passPhrase string) []byte {
	const iterations int = 20000
	const byteLength int = 32

	salt := []byte("SLIP0039" + passPhrase)
	dk := pbkdf2.Key(secret, salt, iterations, byteLength, sha256.New)
	return dk
}
