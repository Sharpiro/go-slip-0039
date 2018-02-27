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

func getIndexLists(wordLists [][]string, bitLength int) [][]byte {
	indexLists := make([][]byte, len(wordLists))
	for i, wordList := range wordLists {
		indexList := getIndexList(wordList)
		preBytes := bits.ReverseBitsBigEndian(indexList[:1], 5, 10, 16)
		// bytes := bits.ReverseBitsBigEndian(indexList[1:len(indexList)-1], 8, 10)
		bytes := bits.ReverseBitsBigEndian(indexList[1:], 8, 10, bitLength)
		// if len(bytes) > 7 {
		// 	extra := bytes[len(bytes)-1]
		// 	bytes = bytes[:len(bytes)-1]
		// 	bytes[len(bytes)-1] += extra
		// }
		// postBytes := bits.ReverseBitsBigEndian(indexList[len(indexList)-1:], 8, 10)
		combined := append(preBytes, bytes...)
		// combined[len(combined)-1] += postBytes[0]
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
			log.Fatal("invalid word found while creating index list")
		}
	}
	return indexes
}
