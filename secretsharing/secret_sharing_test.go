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

func TestShareIndexAndThreshold(tester *testing.T) {
	randomLength := 32
	secretBytes := make([]byte, randomLength)
	cryptoRandom.Read(secretBytes)
	wordLists := CreateSharesX(6, 3, secretBytes)

	if wordLists[0][0] != wordList[5] {
		tester.Error()
	}
	if wordLists[1][0] != wordList[37] {
		tester.Error()
	}
	if wordLists[2][0] != wordList[69] {
		tester.Error()
	}
	if wordLists[3][0] != wordList[101] {
		tester.Error()
	}
	if wordLists[4][0] != wordList[133] {
		tester.Error()
	}
	if wordLists[5][0] != wordList[165] {
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

func TestGetWordList(tester *testing.T) {
	combined := []uint{102, 20, 175, 1009, 3}
	wordList := getWordList(combined)

	fmt.Println(wordList)

	if wordList[0] != "bridge" {
		tester.Error()
	}
	if wordList[1] != "alcohol" {
		tester.Error()
	}
	if wordList[2] != "cousin" {
		tester.Error()
	}
	if wordList[3] != "winter" {
		tester.Error()
	}
	if wordList[4] != "actor" {
		tester.Error()
	}
}

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
