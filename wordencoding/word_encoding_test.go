package wordencoding

import (
	"go-slip-0039/maths/bits"
	"reflect"
	"testing"
)

func TestGetMnemonic(tester *testing.T) {
	resizedIndexList := []uint{1, 44, 160, 321, 97, 878, 682, 576}
	expectedMnemonic := []string{"acid", "arena", "clown", "exhaust", "bracket", "system", "problem", "morning"}
	actualMnemonic := CreateMnemonicWords(resizedIndexList)
	if !reflect.DeepEqual(expectedMnemonic, actualMnemonic) {
		tester.Error()
	}
}

func TestBackToIndexList(tester *testing.T) {
	mnemonicWords := []string{"acid", "arena", "clown", "exhaust", "bracket", "system", "problem", "morning"}
	expectedIndexList := []uint{1, 44, 160, 321, 97, 878, 682, 576}
	actualIndexList := RecoverIndexList(mnemonicWords)
	if !reflect.DeepEqual(expectedIndexList, actualIndexList) {
		tester.Error()
	}
}

// skpping resize...

func TestBackToChecksummedShare(tester *testing.T) {
	indexList := []uint{1, 44, 160, 321, 97, 878, 682, 576}
	// entropySizeBytes := 4
	expectedChecksummedShare := bits.SmartBufferFromBytes([]byte{0, 66, 194, 129, 65, 24, 118, 234, 170, 64}, 74)
	actualChecksummedShare := RecoverChecksummedBuffer(indexList)
	if expectedChecksummedShare.Size != actualChecksummedShare.Size {
		tester.Error()
	}
	if !reflect.DeepEqual(expectedChecksummedShare.Buffer, actualChecksummedShare.Buffer) {
		tester.Error()
	}
}

func TestGetWordList(tester *testing.T) {
	indexList := []uint{102, 20, 175, 1009, 3}
	wordList := CreateMnemonicWords(indexList)

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

func TestGetIndexList(tester *testing.T) {
	wordList := []string{"bridge", "alcohol", "cousin", "winter", "actor"}
	indexList := RecoverIndexList(wordList)

	if indexList[0] != 102 {
		tester.Error()
	}
	if indexList[1] != 20 {
		tester.Error()
	}
	if indexList[2] != 175 {
		tester.Error()
	}
	if indexList[3] != 1009 {
		tester.Error()
	}
	if indexList[4] != 3 {
		tester.Error()
	}
}
