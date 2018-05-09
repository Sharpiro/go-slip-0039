package secretsharing

import (
	"bytes"
	cryptoRandom "crypto/rand"
	"strings"
	"testing"
)

func TestGetMnemonicBuffer(tester *testing.T) {
	words := strings.Split("acid inmate pink program cousin crew", " ")
	indexList := getMnemonicIndexes(words)
	actualBuffer := getMnemonicBuffer(indexList, 16+32)
	expectedBuffer := []byte{0, 1, 116, 41, 58, 180, 175, 46}
	if !bytes.Equal(actualBuffer, expectedBuffer) {
		tester.Error("buffers do not match")
	}
}

func TestGetMnemonicBufferjakubtrnka(tester *testing.T) {
	// words := strings.Split("acoustic exclude genius lucky quarter fuel picnic school", " ")
	words := strings.Split("catch lemon often despair resist response hour lemon", " ")
	indexList := getMnemonicIndexes(words)
	getMnemonicBuffer(indexList, 32+32)
}

func TestShareIndexAndThreshold(tester *testing.T) {
	randomLength := 32
	secretBytes := make([]byte, randomLength)
	cryptoRandom.Read(secretBytes)
	wordLists := CreateWordShares(6, 3, secretBytes)

	if wordLists[0][0] != "acoustic" {
		tester.Error()
	}
	if wordLists[1][0] != "angry" {
		tester.Error()
	}
	if wordLists[2][0] != "bean" {
		tester.Error()
	}
	if wordLists[3][0] != "brain" {
		tester.Error()
	}
	if wordLists[4][0] != "catch" {
		tester.Error()
	}
	if wordLists[5][0] != "clump" {
		tester.Error()
	}
}

func TestGetWordList(tester *testing.T) {
	indexList := []uint{102, 20, 175, 1009, 3}
	wordList := getMnemonic(indexList)

	// fmt.Println(wordList)

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
	indexList := getMnemonicIndexes(wordList)

	// fmt.Println(indexList)

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
