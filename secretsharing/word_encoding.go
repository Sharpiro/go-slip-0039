package secretsharing

import (
	"bytes"
	"go-slip-0039/cryptos"
	"go-slip-0039/maths/bits"
	"log"
)

func getMnemonicList(formattedShares [][]byte, secretByteSize int) [][]string {
	wordLists := make([][]string, len(formattedShares))
	for i := range wordLists {
		// first := bits.GetBitBlocksBigEndian(formattedShares[i][:2], 5, 10)
		// second := bits.GetBitBlocksBigEndian(formattedShares[i][2:], 8, 10)
		// combined := append(first, second...)
		combined := bits.HexToPower2(formattedShares[i], 10)
		resized := bits.ResizeWordIndex(combined, secretByteSize)
		andBack := bits.Power2ToHex(combined, 10)
		andBackResized := bits.Power2ToHex(resized, 10)
		_ = andBack
		_ = andBackResized
		wordLists[i] = getMnemonic(resized)
	}
	return wordLists
}

func getMnemonic(mnemonicIndexes []uint) []string {
	words := make([]string, len(mnemonicIndexes))
	for i, v := range mnemonicIndexes {
		if v&1024 != 0 {
			log.Fatal("word index must be less than 1024")
		}
		words[i] = wordList[v]
	}
	return words
}

// AnalyzeFirstWord analyzes the first word of a share to provide data about the share
func AnalyzeFirstWord(firstWord string) (index int, threshold int) {
	indexList := getMnemonicIndexes([]string{firstWord})
	preBytes := bits.ReverseBitsBigEndian(indexList, 5, 10, 16)
	if len(preBytes) != 2 {
		log.Fatalf("Failed analyzing first word, expected 2 bytes, but was %v", len(preBytes))
	}
	return int(preBytes[0] + 1), int(preBytes[1] + 1)
}

func getMnemonicBuffers(wordLists [][]string, entorpySizeBytes int) [][]byte {
	mnemonicBuffers := make([][]byte, len(wordLists))
	dupeIndexVerifier := make(map[string]bool, len(wordLists))

	for i, wordList := range wordLists {
		if _, exists := dupeIndexVerifier[wordList[0]]; exists {
			log.Fatal("Two shares had identical indexes, each share must have a unique index")
		}
		dupeIndexVerifier[wordList[0]] = true
		mnemonicIndexes := getMnemonicIndexes(wordList)
		mnemonicBuffer := getMnemonicBuffer(mnemonicIndexes, entorpySizeBytes)
		mnemonicBuffers[i] = mnemonicBuffer
	}
	return mnemonicBuffers
}

func getMnemonicBuffer(indexList []uint, entorpySizeBytes int) []byte {
	// tempRebuild2 := bits.HexToPower2(allBytes, 10)
	// tempBuild2Resized := bits.ResizeWordIndex(tempRebuild2, entorpySizeBytes)

	// _ = tempRebuild2
	// _ = tempBuild2Resized
	// preBytes := bits.ReverseBitsBigEndian(indexList[:1], 5, 10, 16)
	// dataWithChecksum := bits.ReverseBitsBigEndian(indexList[1:], 8, 10, entorpySizeBytes*8)
	allBytes := bits.Power2ToHex(indexList, 10)
	expectedChecksum := allBytes[len(allBytes)-2:]
	data := allBytes[:len(allBytes)-2]
	actualChecksum := cryptos.GetSha256(data)[:2]
	if !bytes.Equal(expectedChecksum, actualChecksum) {
		log.Fatal("invalid share checksum")
	}
	final := append(data, expectedChecksum...)
	return final
}

func getMnemonicIndexes(words []string) []uint {
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
