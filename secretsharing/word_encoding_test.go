package secretsharing

import (
	"go-slip-0039/cryptos"
	"testing"
)

func TestShareIndexAndThresholdSimple(tester *testing.T) {
	secretBytes := []byte{9, 8, 7, 6}
	wordLists := CreateWordShares(3, 2, secretBytes)

	if wordLists[0][0] != "acid" {
		tester.Error()
	}
	if wordLists[1][0] != "anger" {
		tester.Error()
	}
	if wordLists[2][0] != "axis" {
		tester.Error()
	}
}

func TestShareIndexAndThreshold(tester *testing.T) {
	randomLength := 32
	secretBytes := cryptos.GetBytes(randomLength)
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
