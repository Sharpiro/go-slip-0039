package cryptos

import "crypto/rand"

// GetBytes gets random buffer of specified length
func GetBytes(length int) []byte {
	buffer := make([]byte, length)
	rand.Read(buffer)
	return buffer
}

// // GetBytes gets random buffer of specified length
// func GetBytes(length int) []byte {
// 	buffer := make([]byte, length)
// 	for i := 0; i < length; i++ {
// 		buffer[i] = byte(i + 1)
// 	}
// 	return buffer
// }
