package secretsharing

import (
	"bytes"
	cryptoRandom "crypto/rand"
	"encoding/hex"
	"fmt"
	mathRandom "math/rand"
	"testing"
	"time"
)

var _tester *testing.T

func TestSecretSharing(tester *testing.T) {
	_tester = tester
	mathRandom.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 100; i++ {
		randomN := uint(mathRandom.Intn(31) + 2)             // 2 <= randomN <= 32
		randomK := uint(mathRandom.Intn(int(randomN-1)) + 2) // 2 <= randomK <= randomN
		randomLength := mathRandom.Intn(64) + 1
		secretBytes := make([]byte, randomLength, randomLength)
		cryptoRandom.Read(secretBytes)
		xValues, yValues := CreateShares(randomN, randomK, secretBytes)
		assertEqual(secretBytes, RecoverSecret(xValues, yValues))
		for j := 0; j < 10; j++ {
			randXValues, randYValues := getRandomSlices(xValues, yValues, randomK)
			assertEqual(secretBytes, RecoverSecret(randXValues, randYValues))
		}
	}
}

// func TestgetRandomSlices(tester *testing.T) {
// 	mathRandom.Seed(time.Now().UTC().UnixNano())
// 	data := []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	randomSlice := getRandomSlices(data, 10)
// 	fmt.Println(randomSlice)
// }

func getRandomSlices(xValues []uint, yValues [][]byte, k uint) ([]uint, [][]byte) {
	tracker := make(map[int]bool, k)
	randXSlice := make([]uint, k, k)
	randYSlice := make([][]byte, k, k)
	var i uint
	for i < k {
		rand := mathRandom.Intn(len(xValues))
		if _, exists := tracker[rand]; !exists {
			randXSlice[i] = xValues[rand]
			randYSlice[i] = yValues[rand]
			tracker[rand] = true
			i++
		}
	}
	return randXSlice, randYSlice
}

func assertEqual(secretBytes, recoveredBytes []byte) {
	if !bytes.Equal(secretBytes, recoveredBytes) {
		recoveredString := hex.EncodeToString(recoveredBytes)
		secretString := hex.EncodeToString(secretBytes)
		fmt.Println(secretString)
		fmt.Println(recoveredString)
		_tester.Error("error")
	}
}
