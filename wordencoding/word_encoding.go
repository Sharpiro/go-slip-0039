package wordencoding

import (
	"go-slip-0039/maths/bits"
	"log"
)

func CreateIndexList(checksummedBuffer *bits.SmartBuffer) []uint {
	if checksummedBuffer.Size < 70 {
		log.Fatalf("Expected checksummed buffer size of at least 70, instead was %v", checksummedBuffer.Size)
	}
	if len(checksummedBuffer.Buffer) < 9 {
		log.Fatalf("Expected checksummed buffer length of at least 9 bytes, instead was %v", len(checksummedBuffer.Buffer))
	}

	indexListRaw := bits.HexToPower2(checksummedBuffer.Buffer, 10)

	indexList := indexListRaw[:checksummedBuffer.Size/10]
	// indexList = bits.ResizeWordIndex(indexListRaw, checksummedBuffer.Size)
	return indexList
}

func CreateMnemonicWords(mnemonicIndexes []int) []string {
	words := make([]string, len(mnemonicIndexes))
	for i, v := range mnemonicIndexes {
		if v&1024 != 0 {
			log.Fatal("word index must be less than 1024")
		}
		words[i] = wordList[v]
	}
	return words
}

func RecoverChecksummedBuffer(indexList []uint) *bits.SmartBuffer {
	allBytes := bits.Power2ToHex(indexList, 10)
	// bytesResized := bits.ResizeBytes(allBytes, secretSizeBytes)
	// bufferBitSize := 20 + 5 + 5 + len(indexList) + 30
	bufferBitSize := len(indexList) * 10
	smartBuffer := bits.SmartBufferFromBytes(allBytes, bufferBitSize)
	return smartBuffer
}

func RecoverIndexList(words []string) []int {
	wordMap := make(map[string]int)
	// for i, v := range wordList {

	// }
	indexes := make([]int, len(words))
	for i, v := range words {
		if val, exists := wordMap[v]; exists {
			indexes[i] = val
		} else {
			log.Fatalf("invalid word %v provided while creating index list", val)
		}
	}
	return indexes
}
