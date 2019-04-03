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
)

var _tester *testing.T

func TestShareIndexAndThresholdSimple(tester *testing.T) {
	// secretBytes := []byte{9, 8, 7, 6}
	secretBytes := []byte{9}
	wordLists := CreateMnemonicWordsList(5, 3, secretBytes, "")

	if wordLists[0][0] != "angry" {
		tester.Error()
	}
	if wordLists[1][0] != "bean" {
		tester.Error()
	}
	if wordLists[2][0] != "brain" {
		tester.Error()
	}
}

func TestShareIndexAndThreshold(tester *testing.T) {
	randomLength := 32
	secretBytes := cryptos.GetRandomBytes(randomLength)
	wordLists := CreateMnemonicWordsList(6, 3, secretBytes, "")

	if wordLists[0][0] != "animal" {
		tester.Error()
	}
	if wordLists[1][0] != "beauty" {
		tester.Error()
	}
	if wordLists[2][0] != "brand" {
		tester.Error()
	}
	if wordLists[3][0] != "category" {
		tester.Error()
	}
	if wordLists[4][0] != "cluster" {
		tester.Error()
	}
	if wordLists[5][0] != "crunch" {
		tester.Error()
	}
}

func TestMakeShare(tester *testing.T) {
	shamirPart := []byte{11, 10, 5, 4, 97, 219}
	expectedShare := []byte{8, 130, 194, 129, 65, 24, 118, 192}
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

// func TestToIndexList(tester *testing.T) {
// 	checksummedShare := bits.SmartBufferFromBytes([]byte{0, 66, 194, 129, 65, 24, 118, 234, 170, 64}, 74)
// 	expectedIndexList := []uint{1, 44, 160, 321, 97, 878, 682, 576, 0}
// 	actualIndexList := bits.HexToPower2(checksummedShare.Buffer, 10)
// 	if !reflect.DeepEqual(expectedIndexList, actualIndexList) {
// 		tester.Error()
// 	}
// }

// func TestToIndexListResized(tester *testing.T) {
// 	secretSizeBytes := 4
// 	indexList := []uint{1, 44, 160, 321, 97, 878, 682, 576, 0}
// 	expectedResizedIndexList := []uint{1, 44, 160, 321, 97, 878, 682, 576}
// 	actualResizedIndexList := bits.ResizeWordIndex(indexList, secretSizeBytes)
// 	if !reflect.DeepEqual(expectedResizedIndexList, actualResizedIndexList) {
// 		tester.Error()
// 	}
// }

func TestBackToUnchecksummedShare(tester *testing.T) {
	checksummedShare := bits.SmartBufferFromBytes([]byte{0, 66, 194, 129, 65, 24, 118, 234, 170, 64}, 74)
	expectedUnchecksummedShare := bits.SmartBufferFromBytes([]byte{0, 66, 194, 129, 65, 24, 118, 192}, 58)
	actualUnchecksummedShare := checksummedShare.GetUnchecksummedBuffer()

	if expectedUnchecksummedShare.Size != actualUnchecksummedShare.Size {
		tester.Error()
	}
	if !reflect.DeepEqual(expectedUnchecksummedShare.Buffer, actualUnchecksummedShare.Buffer) {
		tester.Error()
	}
}

func TestRecoverFromWordShares1Byte(tester *testing.T) {
	var shares = [][]string{
		strings.Split("angry myth faith desk small", " "),
		// strings.Split("bean clog city human catch", " "),
		strings.Split("brain erase special firm grid", " "),
	}
	expectedSecret := []byte{0xff}
	actualSecret := RecoverSecretFromMnemonicShares(shares)
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
		strings.Split("animal jaguar planet elevator travel actress number rural cube betray mixture shoe exact blossom coral play apart distance today living size inmate frame entry ranch faculty eyebrow picnic layer rifle", " "),
		strings.Split("beauty inmate theater music sadness tower chaos jewel join number theater chuckle number tray panic exclude anything brain exile pulse sugar hybrid chair rare pulse steak erase scene trick fiscal", " "),
		// strings.Split("brand bean recall unaware raven napkin army lawn skull fuel rich other razor eyebrow code actor profit obscure fish bowl design lawn tackle design harvest make catalog edit advice rapid", " "),
		strings.Split("category filter rough piano vivid moment master city parade knife share radar else token moment elbow december impact misery artefact hybrid scorpion pluck torch swift harbor lift token cross edge", " "),
	}
	// "13f253e7a4712e2b9a08da7a07e1a5a067ae92adb3fa13649966690c39d901ce"
	expectedSecret := []byte{19, 242, 83, 231, 164, 113, 46, 43, 154, 8, 218, 122, 7, 225, 165, 160, 103, 174, 146, 173, 179, 250, 19, 100, 153, 102, 105, 12, 57, 217, 1, 206}
	actualSecret := RecoverSecretFromMnemonicShares(shares)
	if !bytes.Equal(expectedSecret, actualSecret) {
		tester.Error()
	}
}

func TestRecoverFromWordShares(tester *testing.T) {
	var shares = [][]string{
		strings.Split("angry axis cycle analyst line morning measure exercise", " "),
		strings.Split("bean direct adapt cross public erase level drift", " "),
		strings.Split("brain cheap beyond firm repeat aerobic prison academic", " "),
	}
	expectedSecret := []byte{9, 8, 7, 6}
	actualSecret := RecoverSecretFromMnemonicShares(shares)
	if !bytes.Equal(expectedSecret, actualSecret) {
		tester.Error()
	}
}

func TestRecoverFromWordShares2(tester *testing.T) {
	var shares = [][]string{
		strings.Split("animal head debris negative artefact slush alpha frequent wash aspect warning", " "),
		// strings.Split("beauty drink recipe december fiction again manual network source source force", " "),
		strings.Split("brand chat other science broken slow luxury sunset decorate rare burden", " "),
		// strings.Split("category paper recall hurt carbon ceiling thunder someone twice myself morning", " "),
		strings.Split("cluster wire orient ordinary express used spread line believe stadium lock", " "),
	}
	expectedSecret := []byte{8, 7, 6, 5, 4, 3, 2, 1}
	actualSecret := RecoverSecretFromMnemonicShares(shares)
	if !bytes.Equal(expectedSecret, actualSecret) {
		tester.Error()
	}
}

func TestSecretSharingWords(tester *testing.T) {
	for j := 0; j < 10; j++ {
		for i := 5; i < 64; i++ {
			byteLength := i + 1
			actualSecret := cryptos.GetRandomBytes(byteLength)
			wordShares := CreateMnemonicWordsList(6, 3, actualSecret, "")
			expectedSecret := RecoverSecretFromMnemonicShares(wordShares)

			if !bytes.Equal(actualSecret, expectedSecret) {
				tester.Error("secrets do not match")
			}
		}
	}
}

func TestRecoverShare(tester *testing.T) {
	share := strings.Split("actress anchor angry wolf league target circle", " ")
	RecoverShare(share)
}

func TestRecoverSecretFromMnemonicShares(tester *testing.T) {
	shares := [][]string{
		strings.Split("actress anchor angry wolf league target circle", " "),
		strings.Split("actress anchor brain wealth jaguar stage music", " "),
	}
	secret := RecoverSecretFromMnemonicShares(shares)
	if len(secret) != 1 || secret[0] != 255 {
		tester.Error()
	}
}

// func TestGetChecksummedSecret(tester *testing.T) {
// 	data := []byte("data")
// 	hash := cryptos.GetSha256(data)[:2]
// 	css := createChecksummedSecret(data)
// 	if !bytes.Equal(data, css[:4]) {
// 		tester.Error()
// 	}
// 	if !bytes.Equal(hash, css[4:]) {
// 		tester.Error()
// 	}
// }

// func TestSecretSharing(tester *testing.T) {
// 	_tester = tester
// 	mathRandom.Seed(time.Now().UTC().UnixNano())
// 	for i := 0; i < 100; i++ {
// 		randomN := uint(mathRandom.Intn(31) + 2)             // 2 <= randomN <= 32
// 		randomK := uint(mathRandom.Intn(int(randomN-1)) + 2) // 2 <= randomK <= randomN
// 		randomLength := mathRandom.Intn(64) + 1
// 		secretBytes := cryptos.GetBytes(randomLength)
// 		xValues, yValues := createShamirData(randomN, randomK, secretBytes)
// 		assertEqual(secretBytes, recoverChecksummedSecret(xValues, yValues))
// 		for j := 0; j < 10; j++ {
// 			randXValues, randYValues := getRandomSlices(xValues, yValues, randomK)
// 			assertEqual(secretBytes, recoverChecksummedSecret(randXValues, randYValues))
// 		}
// 	}
// }

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
