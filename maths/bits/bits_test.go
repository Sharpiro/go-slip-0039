package bits

import (
	"testing"
)

func TestGetWordIndexes(tester *testing.T) {
	pre := []byte{3, 6}
	list := []byte{5, 10, 255, 199} // 0b00000101 0b00001010 0b11111111 0b11000111
	wordIndexes := GetBitBlocksBigEndian(pre, 5, 10)
	wordIndexes2 := GetBitBlocksBigEndian(list, 8, 10)

	// fmt.Println(wordIndexes)
	// fmt.Println(wordIndexes2)

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
