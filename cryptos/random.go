package cryptos

import "crypto/rand"

// GetRandomBytes gets random buffer of specified length
func GetRandomBytes(length int) []byte {
	buffer := make([]byte, length)
	rand.Read(buffer)
	return buffer
}

// // GetRandomBytes gets fake buffer sequence of specified length
// func GetRandomBytes(length int) []byte {
// 	buffer := make([]byte, length)
// 	for i := 0; i < length; i++ {
// 		buffer[i] = byte(i + 1)
// 	}
// 	return buffer
// }
