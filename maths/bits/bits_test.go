package bits

import (
	"testing"
)

func TestGetBitBlocksBigEndian(tester *testing.T) {
	pre := []byte{3, 6}
	list := []byte{5, 10, 255, 199} // 0b00000101 0b00001010 0b11111111 0b11000111
	wordIndexes := GetBitBlocksBigEndian(pre, 5, 10)
	wordIndexes2 := GetBitBlocksBigEndian(list, 8, 10)

	// fmt.Println(wordIndexes)
	// fmt.Println(wordIndexes2) // 20, 175, 1009, 3

	if wordIndexes[0] != 102 { //00011 00110
		tester.Error()
	}
	if wordIndexes2[0] != 20 {
		tester.Error()
	}
	if wordIndexes2[1] != 175 {
		tester.Error()
	}
	if wordIndexes2[2] != 1009 {
		tester.Error()
	}
	if wordIndexes2[3] != 3 {
		tester.Error()
	}
}

func TestGetBitBlocksLittleEndian(tester *testing.T) {
	pre := []byte{3, 6}
	list := []byte{5, 10, 255, 199} // 0b00000101 0b00001010 0b11111111 0b11000111
	wordIndexes := GetBitBlocksLittleEndian(pre, 5, 10)
	wordIndexes2 := GetBitBlocksLittleEndian(list, 8, 10)

	// fmt.Println()
	// fmt.Println(wordIndexes)
	// fmt.Println(wordIndexes2) // 641, 271, 1016, 3

	if wordIndexes[0] != 780 { //00011 00110
		tester.Error()
	}
	if wordIndexes2[0] != 641 {
		tester.Error()
	}
	if wordIndexes2[1] != 271 {
		tester.Error()
	}
	if wordIndexes2[2] != 1016 {
		tester.Error()
	}
	if wordIndexes2[3] != 3 {
		tester.Error()
	}
}
