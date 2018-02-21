package secretsharing

import (
	"bytes"
	cryptoRandom "crypto/rand"
	"encoding/hex"
	"fmt"
	mathRandom "math/rand"
	"testing"
	"time"

	"../cryptos"
)

var _tester *testing.T

func TestSecretSharing(tester *testing.T) {
	_tester = tester
	mathRandom.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 100; i++ {
		randomN := uint(mathRandom.Intn(31) + 2)             // 2 <= randomN <= 32
		randomK := uint(mathRandom.Intn(int(randomN-1)) + 2) // 2 <= randomK <= randomN
		randomLength := mathRandom.Intn(64) + 1
		secretBytes := make([]byte, randomLength)
		cryptoRandom.Read(secretBytes)
		xValues, yValues := CreateShares(randomN, randomK, secretBytes)
		assertEqual(secretBytes, RecoverSecret(xValues, yValues))
		for j := 0; j < 10; j++ {
			randXValues, randYValues := getRandomSlices(xValues, yValues, randomK)
			assertEqual(secretBytes, RecoverSecret(randXValues, randYValues))
		}
	}
}

func TestGetWordIndexes(tester *testing.T) {
	pre := []byte{3, 6}
	list := []byte{5, 10, 255, 199} // 0b00000101 0b00001010 0b11111111 0b11000111
	wordIndexes := getWordIndexes(pre, 5, 10)
	wordIndexes2 := getWordIndexes(list, 8, 10)

	// fmt.Println(wordIndexes)
	// fmt.Println(wordIndexes2)

	if wordIndexes[0] != 102 { //00011 00110
		tester.Error()
	}
	if wordIndexes2[0] != 20 {
		tester.Error()
	}
	if wordIndexes2[1] != 175 {
		tester.Error()
	}
	if wordIndexes2[2] != 1009 {
		tester.Error()
	}
	if wordIndexes2[3] != 3 {
		tester.Error()
	}
}

// func Test256BitShare(tester *testing.T) {
// 	secretBytes := make([]byte, 32)
// 	mathRandom.Read(secretBytes)
// 	xValues, yValues := CreateShares(6, 3, secretBytes)
// 	fXValues := createFormattedShare(xValues, yValues)
// 	fmt.Println(fXValues[0])
// 	fmt.Println(len(fXValues[0]))
// }

func TestGetChecksummedSecret(tester *testing.T) {
	data := []byte("data")
	hash := cryptos.GetSha256(data)[:2]
	css := getChecksummedSecret(data)
	if !bytes.Equal(data, css[:4]) {
		tester.Error()
	}
	if !bytes.Equal(hash, css[4:]) {
		tester.Error()
	}
}

func getRandomSlices(xValues []uint, yValues [][]byte, k uint) ([]uint, [][]byte) {
	tracker := make(map[int]bool, k)
	randXSlice := make([]uint, k)
	randYSlice := make([][]byte, k)
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
