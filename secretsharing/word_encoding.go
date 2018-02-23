package secretsharing

import (
	"../maths/bits"
	"log"
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

func getIndexLists(wordLists [][]string) [][]byte {
	indexLists := make([][]byte, len(wordLists))
	for _, wordList := range wordLists {
		indexList := getIndexList(wordList)
		_=indexList
		// indexLists[i] = indexList
	}
	return indexLists
}

func getIndexList(words []string) []uint {
	indexes := make([]uint, len(words))
	for i, v := range words {
		if val, exists := wordMap[v]; exists {
			indexes[i] = val
		} else {
			log.Fatal("invalid word found while creating index list")
		}
	}
	return indexes
}
