package secretsharing

import (
	"log"

	"go-slip-0039/maths/bits"
)

func getWordLists(formattedShares [][]byte) [][]string {
	wordLists := make([][]string, len(formattedShares))
	for i := range wordLists {
		first := bits.GetBitBlocksBigEndian(formattedShares[i][:2], 5, 10)
		second := bits.GetBitBlocksBigEndian(formattedShares[i][2:], 8, 10)
		combined := append(first, second...)
		wordLists[i] = getWordList(combined)
	}
	return wordLists
}

func getWordList(combined []uint) []string {
	words := make([]string, len(combined))
	for i, v := range combined {
		if v&1024 != 0 {
			log.Fatal("word index must be less than 1024")
		}
		words[i] = wordList[v]
	}
	return words
}

// AnalyzeFirstWord analyzes the first word of a share to provide data about the share
func AnalyzeFirstWord(firstWord string) (index byte, threshold byte) {
	indexList := getIndexList([]string{firstWord})
	preBytes := bits.ReverseBitsBigEndian(indexList, 5, 10, 16)
	if len(preBytes) != 2 {
		log.Fatalf("Failed analyzing first word, expected 2 bytes, but was %v", len(preBytes))
	}
	return preBytes[0] + 1, preBytes[1] + 1
}

func getIndexLists(wordLists [][]string, bitLength int) [][]byte {
	indexLists := make([][]byte, len(wordLists))
	dupeIndexVerifier := make(map[string]bool, len(wordLists))

	for i, wordList := range wordLists {
		if _, exists := dupeIndexVerifier[wordList[0]]; exists {
			log.Fatal("Two shares had identical indexes, each share must have a unique index")
		}
		dupeIndexVerifier[wordList[0]] = true
		indexList := getIndexList(wordList)
		preBytes := bits.ReverseBitsBigEndian(indexList[:1], 5, 10, 16)
		bytes := bits.ReverseBitsBigEndian(indexList[1:], 8, 10, bitLength)
		combined := append(preBytes, bytes...)
		indexLists[i] = combined
	}
	return indexLists
}

func getIndexList(words []string) []uint {
	indexes := make([]uint, len(words))
	for i, v := range words {
		if val, exists := wordMap[v]; exists {
			indexes[i] = val
		} else {
			log.Fatal("invalid word provided while creating index list")
		}
	}
	return indexes
}
