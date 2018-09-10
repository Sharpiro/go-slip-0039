package wordencoding

import (
	"go-slip-0039/maths/bits"
	"log"
)

func CreateIndexList(checksummedBuffer *bits.SmartBuffer, secretSizeBytes int) []uint {
	indexList := bits.HexToPower2(checksummedBuffer.Buffer, 10)
	indexList = bits.ResizeWordIndex(indexList, secretSizeBytes)
	return indexList
}

func CreateMnemonicWords(mnemonicIndexes []uint) []string {
	words := make([]string, len(mnemonicIndexes))
	for i, v := range mnemonicIndexes {
		if v&1024 != 0 {
			log.Fatal("word index must be less than 1024")
		}
		words[i] = wordList[v]
	}
	return words
}

func RecoverChecksummedBuffer(indexList []uint, paddedSecretSizeBits int) *bits.SmartBuffer {
	allBytes := bits.Power2ToHex(indexList, 10)
	// bytesResized := bits.ResizeBytes(allBytes, secretSizeBytes)
	bufferBitSize := 20 + 5 + 5 + paddedSecretSizeBits + 30
	smartBuffer := bits.SmartBufferFromBytes(allBytes, bufferBitSize)
	return smartBuffer
}

func RecoverIndexes(words []string) []uint {
	indexes := make([]uint, len(words))
	for i, v := range words {
		if val, exists := wordMap[v]; exists {
			indexes[i] = val
		} else {
			log.Fatalf("invalid word %v provided while creating index list", val)
		}
	}
	return indexes
}
