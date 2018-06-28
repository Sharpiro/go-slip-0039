package secretsharing

import (
	"go-slip-0039/cryptos"
	"strings"
	"testing"
)

func TestGetMnemonicBuffer(tester *testing.T) {
	words := strings.Split("academic alpha crystal ocean rapid pave", " ")
	indexList := getMnemonicIndexes(words)
	entropySizeBytes := 2
	getMnemonicBuffer(indexList, entropySizeBytes)
}

// func TestGetMnemonicBufferjakubtrnka(tester *testing.T) {
// 	// 	acid glance scatter multiply muscle evolve vote hedgehog vanish shoe road sense ugly raise sister scout educate
// 	// anger mansion second exclude grow garden video purchase cost skin crowd surface brush choice machine lock shed
// 	// axis clerk cupboard golden endless minute army lecture fuel soldier peace regret deny extra group execute menu
// 	hexEntropy := "c9d32bb6f9a2024b9e12c2cd4af717c1"
// 	entropySizeBytes := len(hexEntropy) / 2 // 16 bytes | 128 bits

// 	indexList := []uint{1, 400, 766, 582, 583, 312, 984, 430, 960, 795, 745, 785, 935, 706, 806, 772, 274}
// 	getMnemonicBuffer(indexList, entropySizeBytes)
// }

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
