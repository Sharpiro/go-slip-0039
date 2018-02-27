package secretsharing

import (
	cryptoRandom "crypto/rand"
	"testing"
)

func TestShareIndexAndThreshold(tester *testing.T) {
	randomLength := 32
	secretBytes := make([]byte, randomLength)
	cryptoRandom.Read(secretBytes)
	wordLists := CreateWordShares(6, 3, secretBytes)

	if wordLists[0][0] != "adapt" {
		tester.Error()
	}
	if wordLists[1][0] != "antenna" {
		tester.Error()
	}
	if wordLists[2][0] != "become" {
		tester.Error()
	}
	if wordLists[3][0] != "bread" {
		tester.Error()
	}
	if wordLists[4][0] != "ceiling" {
		tester.Error()
	}
	if wordLists[5][0] != "coconut" {
		tester.Error()
	}
}

func TestGetWordList(tester *testing.T) {
	indexList := []uint{102, 20, 175, 1009, 3}
	wordList := getWordList(indexList)

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
	indexList := getIndexList(wordList)

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
