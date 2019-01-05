package cryptos

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

// CreatePbkdf2Hash creates a hash based off a given secret and an optional passphrase
func CreatePbkdf2Hash(password string, identifier string, threshold byte) []byte {
	const iterations int = 20000
	const byteLength int = 32

	idPartOne := identifier[:10]
	idPartTwo := identifier[10:20]
	idPartThree := identifier[20:]
	_ = idPartOne
	_ = idPartTwo
	_ = idPartThree
	// identifierBytes := bits.GetBits(byte(index), 5)
	identifierBytes := []byte{1, 2, 3, 4, 5, 6}
	salt := []byte("slip0039")
	salt = append(salt, identifierBytes...)
	salt = append(salt, threshold)
	hash := pbkdf2.Key([]byte(password), salt, iterations, byteLength, sha256.New)
	return hash
}
