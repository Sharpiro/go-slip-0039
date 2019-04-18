package cryptos

import (
	"crypto/rand"
	"log"
	"math/big"
)

// GetRandomBytes gets random buffer of specified length
func GetRandomBytes(length int) []byte {
	buffer := make([]byte, length)
	rand.Read(buffer)
	return buffer
}

// GetRandomBelow returns a random number below 'number'
func GetRandomBelow(number int) int {
	temp, err := rand.Int(rand.Reader, big.NewInt(int64(number)))
	if err != nil {
		log.Fatal(err)
	}
	return int(temp.Int64())
}

// // GetRandomBytes gets fake buffer sequence of specified length
// func GetRandomBytes(length int) []byte {
// 	buffer := make([]byte, length)
// 	for i := 0; i < length; i++ {
// 		buffer[i] = byte(i + 1)
// 	}
// 	return buffer
// }
