package secretsharing

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"go-slip-0039/cryptos"
	"go-slip-0039/maths/bits"
	mathRandom "math/rand"
	"reflect"
	"strings"
	"testing"
	"time"
)

var _tester *testing.T

func TestMakeShare(tester *testing.T) {
	shamirPart := []byte{11, 10, 5, 4, 97, 219}
	expectedShare := []byte{0, 66, 194, 129, 65, 24, 118, 192}
	actualShare := createUnchecksummedShare(shamirPart, 1, 2)
	if !bytes.Equal(expectedShare, actualShare.Buffer) {
		tester.Error()
	}
	if actualShare.Size != 58 {
		tester.Error()
	}
}

func TestGetChecksummedShare(tester *testing.T) {
	share := bits.SmartBufferFromBytes([]byte{0, 66, 194, 129, 65, 24, 118, 192}, 58)
	expectedChecksummedShare := bits.SmartBufferFromBytes([]byte{0, 66, 194, 129, 65, 24, 118, 234, 170, 64}, 74)
	actualChecksummedShare := share.GetChecksummedBuffer()
	if !bytes.Equal(expectedChecksummedShare.Buffer, actualChecksummedShare.Buffer) {
		tester.Error()
	}
	if expectedChecksummedShare.Size != actualChecksummedShare.Size {
		tester.Error()
	}
}

func TestToIndexList(tester *testing.T) {
	checksummedShare := bits.SmartBufferFromBytes([]byte{0, 66, 194, 129, 65, 24, 118, 234, 170, 64}, 74)
	expectedIndexList := []uint{1, 44, 160, 321, 97, 878, 682, 576, 0}
	actualIndexList := bits.HexToPower2(checksummedShare.Buffer, 10)
	if !reflect.DeepEqual(expectedIndexList, actualIndexList) {
		tester.Error()
	}
}

func TestToIndexListResized(tester *testing.T) {
	secretSizeBytes := 4
	indexList := []uint{1, 44, 160, 321, 97, 878, 682, 576, 0}
	expectedResizedIndexList := []uint{1, 44, 160, 321, 97, 878, 682, 576}
	actualResizedIndexList := bits.ResizeWordIndex(indexList, secretSizeBytes)
	if !reflect.DeepEqual(expectedResizedIndexList, actualResizedIndexList) {
		tester.Error()
	}
}

func TestGetMnemonic(tester *testing.T) {
	resizedIndexList := []uint{1, 44, 160, 321, 97, 878, 682, 576}
	expectedMnemonic := []string{"acid", "arena", "clown", "exhaust", "bracket", "system", "problem", "morning"}
	actualMnemonic := createMnemonicWords(resizedIndexList)
	if !reflect.DeepEqual(expectedMnemonic, actualMnemonic) {
		tester.Error()
	}
}

func TestBackToIndexList(tester *testing.T) {
	mnemonicWords := []string{"acid", "arena", "clown", "exhaust", "bracket", "system", "problem", "morning"}
	expectedIndexList := []uint{1, 44, 160, 321, 97, 878, 682, 576}
	actualIndexList := recoverIndexes(mnemonicWords)
	if !reflect.DeepEqual(expectedIndexList, actualIndexList) {
		tester.Error()
	}
}

// skpping resize...

func TestBackToChecksummedShare(tester *testing.T) {
	indexList := []uint{1, 44, 160, 321, 97, 878, 682, 576}
	entropySizeBytes := 4
	expectedChecksummedShare := bits.SmartBufferFromBytes([]byte{0, 66, 194, 129, 65, 24, 118, 234, 170, 64}, 74)
	actualChecksummedShare := recoverChecksummedBuffer(indexList, entropySizeBytes)
	if expectedChecksummedShare.Size != actualChecksummedShare.Size {
		tester.Error()
	}
	if !reflect.DeepEqual(expectedChecksummedShare.Buffer, actualChecksummedShare.Buffer) {
		tester.Error()
	}
}

func TestBackToUnchecksummedShare(tester *testing.T) {
	checksummedShare := bits.SmartBufferFromBytes([]byte{0, 66, 194, 129, 65, 24, 118, 234, 170, 64}, 74)
	expectedUnchecksummedShare := bits.SmartBufferFromBytes([]byte{0, 66, 194, 129, 65, 24, 118, 192}, 58)
	actualUnchecksummedShare := checksummedShare.GetUnchecksummedBuffer(2)

	if expectedUnchecksummedShare.Size != actualUnchecksummedShare.Size {
		tester.Error()
	}
	if !reflect.DeepEqual(expectedUnchecksummedShare.Buffer, actualUnchecksummedShare.Buffer) {
		tester.Error()
	}
}

func TestSecretSharing(tester *testing.T) {
	_tester = tester
	mathRandom.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 100; i++ {
		randomN := uint(mathRandom.Intn(31) + 2)             // 2 <= randomN <= 32
		randomK := uint(mathRandom.Intn(int(randomN-1)) + 2) // 2 <= randomK <= randomN
		randomLength := mathRandom.Intn(64) + 1
		secretBytes := cryptos.GetBytes(randomLength)
		xValues, yValues := createShamirData(randomN, randomK, secretBytes)
		assertEqual(secretBytes, recoverChecksummedSecret(xValues, yValues))
		for j := 0; j < 10; j++ {
			randXValues, randYValues := getRandomSlices(xValues, yValues, randomK)
			assertEqual(secretBytes, recoverChecksummedSecret(randXValues, randYValues))
		}
	}
}

func TestRecoverFromWordShares1Byte(tester *testing.T) {
	var shares = [][]string{
		// strings.Split("acid world predict country obey", " "),
		strings.Split("anger width radio engage cement", " "),
		strings.Split("axis weather reward furnace library", " "),
	}
	expectedSecret := []byte{0xff}
	actualSecret := RecoverSecretFromMnemonicShares(shares, len(expectedSecret))
	if !bytes.Equal(expectedSecret, actualSecret) {
		tester.Error()
	}
}

func TestRecoverFromWordShares32Bytes(tester *testing.T) {
	n := 4
	k := 3
	_ = n
	_ = k
	var shares = [][]string{
		strings.Split("acoustic benefit smoke cricket primary image runway priority search symptom unique hundred coach pelican organ wealth under recall universe click grass group pave staff delay actor divert endorse shock elder", " "),
		strings.Split("angry bulb type sausage under juice october destroy lemon spray siege wrestle heavy predict sauce hand primary rough silent resemble city hurdle lock earth insane anxiety brand surface music picture", " "),
		// strings.Split("bean brown upset question program jewel pact coach science stadium slow usual corn primary robot jungle twist robust slush please glance idea lemon eyebrow debris animal busy similar stadium window", " "),
		strings.Split("brain cousin code salt trouble enforce find devote mercy token animal world group ocean leaf hazard pistol menu angry repair club fresh elbow drift join device suit tackle purchase glue", " "),
	}
	// "13f253e7a4712e2b9a08da7a07e1a5a067ae92adb3fa13649966690c39d901ce"
	expectedSecret := []byte{19, 242, 83, 231, 164, 113, 46, 43, 154, 8, 218, 122, 7, 225, 165, 160, 103, 174, 146, 173, 179, 250, 19, 100, 153, 102, 105, 12, 57, 217, 1, 206}
	actualSecret := RecoverSecretFromMnemonicShares(shares, len(expectedSecret))
	if !bytes.Equal(expectedSecret, actualSecret) {
		tester.Error()
	}
}

func TestRecoverFromWordShares(tester *testing.T) {
	var shares = [][]string{
		[]string{"acid", "arena", "clown", "exhaust", "bracket", "system", "problem", "morning"},
		[]string{"axis", "awake", "desert", "awkward", "bread", "thunder", "rude", "timber"},
	}
	expectedSecret := []byte{9, 8, 7, 6}
	actualSecret := RecoverSecretFromMnemonicShares(shares, len(expectedSecret))
	if !bytes.Equal(expectedSecret, actualSecret) {
		tester.Error()
	}
}

func TestRecoverFromWordShares2(tester *testing.T) {
	var shares = [][]string{
		[]string{"acoustic", "answer", "bowl", "imitate", "adapt", "adult", "army", "agent", "early", "nice", "lock"},
		[]string{"bean", "actress", "desert", "velvet", "again", "anything", "cover", "lizard", "drum", "manage", "image"},
		[]string{"clump", "desert", "taxi", "gentle", "eternal", "damage", "similar", "bean", "avoid", "earth", "obtain"},
	}
	expectedSecret := []byte{8, 7, 6, 5, 4, 3, 2, 1}
	actualSecret := RecoverSecretFromMnemonicShares(shares, len(expectedSecret))
	if !bytes.Equal(expectedSecret, actualSecret) {
		tester.Error()
	}
}

func TestSecretSharingWords(tester *testing.T) {
	for j := 0; j < 10; j++ {
		for i := 5; i < 64; i++ {
			byteLength := i + 1
			actualSecret := cryptos.GetBytes(byteLength)
			wordShares := CreateMnemonicWordsList(6, 3, actualSecret)
			expectedSecret := RecoverSecretFromMnemonicShares(wordShares, byteLength)

			if !bytes.Equal(actualSecret, expectedSecret) {
				tester.Error("secrets do not match")
			}
		}
	}
}

func TestGetChecksummedSecret(tester *testing.T) {
	data := []byte("data")
	hash := cryptos.GetSha256(data)[:2]
	css := createChecksummedSecret(data)
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
