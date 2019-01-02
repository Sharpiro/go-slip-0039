package cryptos

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"

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

// EncryptAes creates a hash based off a given secret and an optional passphrase
func EncryptAes(key, data []byte) []byte {
	aesCipher, err := aes.NewCipher(key)
	_ = aesCipher
	_ = err
	temp, err := cipher.NewGCM(aesCipher)
	_ = temp
	cryptoBytes := make([]byte, 32)
	aesCipher.Encrypt(cryptoBytes, data)
	return cryptoBytes
}

func Encrypt(plaintext []byte, key *[32]byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}
