package secretsharing

import (
	"bytes"
	"go-slip-0039/cryptos"
	"go-slip-0039/maths/bits"
	"log"
)

func getIndexesList(smartBuffers []*bits.SmartBuffer, secretSizeBytes int) [][]uint {
	indexesList := make([][]uint, len(smartBuffers))
	for i, v := range smartBuffers {
		indexesList[i] = bits.HexToPower2(v.Buffer, 10)
		indexesList[i] = bits.ResizeWordIndex(indexesList[i], secretSizeBytes)
	}
	return indexesList
}

func getMnemonicList(indexList [][]uint) [][]string {
	wordLists := make([][]string, len(indexList))
	for i, v := range indexList {
		wordLists[i] = getMnemonic(v)
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

func getChecksummedBuffers(indexLists [][]uint, entorpySizeBytes int) []*bits.SmartBuffer {
	mnemonicBuffers := make([]*bits.SmartBuffer, len(indexLists))

	for i, indexList := range indexLists {
		mnemonicBuffer := getChecksummedBuffer(indexList, entorpySizeBytes)
		mnemonicBuffers[i] = mnemonicBuffer
	}
	return mnemonicBuffers
}

func getUnchecksummedBuffers(smartBuffers []*bits.SmartBuffer) []*bits.SmartBuffer {
	unchecksummedBuffers := make([]*bits.SmartBuffer, len(smartBuffers))

	for i, checksummedBuffer := range smartBuffers {
		unchecksummedBuffer := getUnchecksummedBuffer(checksummedBuffer)
		unchecksummedBuffers[i] = unchecksummedBuffer
	}
	return unchecksummedBuffers
}

func getChecksummedBuffer(indexList []uint, entorpySizeBytes int) *bits.SmartBuffer {
	allBytes := bits.Power2ToHex(indexList, 10)
	bytesResized := bits.ResizeBytes(allBytes, entorpySizeBytes)
	bufferBitSize := entorpySizeBytes*8 + 16 + 16 + 10
	smartBuffer := bits.SmartBufferFromBytes(bytesResized, bufferBitSize)
	return smartBuffer
}

func getUnchecksummedBuffer(smartBuffer *bits.SmartBuffer) *bits.SmartBuffer {
	cloneBuffer := smartBuffer.Clone()
	expectedChecksum := cloneBuffer.PopBits(16)
	actualChecksum := cryptos.GetSha256(cloneBuffer.Buffer)[:2]
	if !bytes.Equal(expectedChecksum.Buffer, actualChecksum) {
		log.Fatal("invalid share checksum")
	}
	return cloneBuffer
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

func getMnemonicIndexesList(wordLists [][]string) [][]uint {
	indexesList := make([][]uint, len(wordLists))
	dupeIndexVerifier := make(map[string]bool, len(wordLists))
	for i, wordList := range wordLists {
		if _, exists := dupeIndexVerifier[wordList[0]]; exists {
			log.Fatal("Two shares had identical indexes, each share must have a unique index")
		}
		dupeIndexVerifier[wordList[0]] = true
		indexesList[i] = getMnemonicIndexes(wordList)
	}
	return indexesList
}
