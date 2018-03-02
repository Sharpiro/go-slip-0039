package secretsharing

import (
	"bytes"
	cryptoRandom "crypto/rand"
	"encoding/hex"
	"fmt"
	mathRandom "math/rand"
	"testing"
	"time"

	"go-slip-0039/cryptos"
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
		xValues, yValues := createShares(randomN, randomK, secretBytes)
		assertEqual(secretBytes, recoverSecret(xValues, yValues))
		for j := 0; j < 10; j++ {
			randXValues, randYValues := getRandomSlices(xValues, yValues, randomK)
			assertEqual(secretBytes, recoverSecret(randXValues, randYValues))
		}
	}
}

func TestSecretSharingWords(tester *testing.T) {
	// actualSecret := []byte("doggg")
	// actualSecret := []byte{1}
	// actualSecret := []byte{1, 1, 1}
	// formattedShares := [][]byte{[]byte{0, 5, 43, 82, 115, 166, 120}}
	// wordLists := getWordLists(formattedShares)
	// indexLists := getIndexLists(wordLists)

	for j := 0; j < 10; j++ {
		for i := 0; i < 64; i++ {
			byteLength := i + 1
			bitLength := (byteLength + 2 + 2) << 3
			actualSecret := make([]byte, byteLength)
			cryptoRandom.Read(actualSecret)
			wordShares := CreateWordShares(6, 3, actualSecret)
			expectedSecret := RecoverFromWordShares(wordShares, bitLength)

			// fmt.Println(wordShares)
			// fmt.Println(formattedShares[0])
			// fmt.Println(indexLists[0])
			// fmt.Println(len(actualSecret))
			// fmt.Println(len(expectedSecret))
			// fmt.Println(actualSecret)
			// fmt.Println(expectedSecret)
			if !bytes.Equal(actualSecret, expectedSecret) {
				tester.Error("secrets do not match")
			}
		}
	}
}

func TestShareFormatting(tester *testing.T) {
	secretBytes := make([]byte, 32)
	mathRandom.Read(secretBytes)
	xValues, yValues := createShares(6, 3, secretBytes)
	formattedShares := createFormattedShares(xValues, yValues, 3)
	recoveredXValues, recoveredYValues := recoverFromFormattedShare(formattedShares)
	for i, v := range formattedShares {
		if len(v) != 36 {
			tester.Error("expected formatted share to be 36 bytes")
		}
		if !bytes.Equal(yValues[i], recoveredYValues[i]) {
			tester.Error("recovered y values did not match expected")
		}
		if xValues[i] != recoveredXValues[i] {
			tester.Error("recovered x values did not match expected")
		}
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
