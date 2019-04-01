package secretsharing

import (
	"go-slip-0039/cryptos"
	"go-slip-0039/maths/bits"
	"log"
)

func getEncryptionKey(passphrase string, sharedIdentifier string, threshold uint) []byte {
	passwordBytes := []byte(passphrase)
	identifierNumberOneBits := sharedIdentifier[:10]
	identifierNumberOneBytes := bits.GetBytesLittleEndian(identifierNumberOneBits, 2)
	identifierNumberTwoBits := sharedIdentifier[10:20]
	identifierNumberTwoBytes := bits.GetBytesLittleEndian(identifierNumberTwoBits, 2)
	identifierNumberThreeBits := sharedIdentifier[20:30]
	identifierNumberThreeBytes := bits.GetBytesLittleEndian(identifierNumberThreeBits, 2)
	thresholdBytes := []byte{byte(threshold)}
	salt := append([]byte("slip0039"), identifierNumberOneBytes...)
	salt = append(salt, identifierNumberTwoBytes...)
	salt = append(salt, identifierNumberThreeBytes...)
	salt = append(salt, thresholdBytes...)

	key := cryptos.CreatePbkdf2Hash(passwordBytes, salt)
	return key
}

func getMasterSecret(preMasterSecret []byte, key *[32]byte) []byte {
	secret, err := cryptos.EncryptGCM(preMasterSecret, key)
	if err != nil {
		log.Fatal(err)
	}
	return secret
}

func getPreMasterSecret(masterSecret []byte, key *[32]byte) []byte {
	preMasterSecret, err := cryptos.DecryptGCM(masterSecret, key)
	if err != nil {
		log.Fatal(err)
	}
	return preMasterSecret
}
