package bits

import (
	"fmt"
	"testing"
)

func TestReverseBitsBigEndian(tester *testing.T) {
	indexes := []uint{102, 20, 175, 1009, 3}
	preBytes := ReverseBitsBigEndian(indexes[:1], 5, 10)
	bytes := ReverseBitsBigEndian(indexes[1:5], 8, 10)
	if len(bytes) > 4 {
		extra := bytes[len(bytes)-1]
		bytes = bytes[:len(bytes)-1]
		bytes[len(bytes)-1] += extra
	}
	// postBytes := ReverseBitsBigEndian(indexes[4:], 10, 10)

	// fmt.Println(preBytes)
	// fmt.Println(bytes)
	// fmt.Println(postBytes)

	if len(preBytes) != 2 {
		tester.Error()
	}
	if preBytes[0] != 3 {
		tester.Error()
	}
	if preBytes[1] != 6 {
		tester.Error()
	}
	if len(bytes) != 4 {
		tester.Error()
	}
	if bytes[0] != 5 {
		tester.Error()
	}
	if bytes[1] != 10 {
		tester.Error()
	}
	if bytes[2] != 255 {
		tester.Error()
	}
	if bytes[3] != 199 {
		tester.Error()
	}
}

func TestReverseBitsBigEndian2(tester *testing.T) {
	indexes := []uint{5, 173, 295, 233, 632}
	// expectedBytes := []byte{0, 5, 43, 82, 115, 166, 120}

	preSlice := indexes[:1]
	preBytes := ReverseBitsBigEndian(preSlice, 5, 10)
	slice := indexes[1:5]
	bytes := ReverseBitsBigEndian(slice, 8, 10)
	// postSlice := indexes[4:]
	// postBytes := ReverseBitsBigEndian(postSlice, 8, 10)

	// fmt.Println(preBytes)
	// fmt.Println()
	// fmt.Println(bytes)
	// fmt.Println(postBytes)

	if len(preBytes) != 2 {
		tester.Error()
	}
	if preBytes[0] != 0 {
		tester.Error()
	}
	if preBytes[1] != 5 {
		tester.Error()
	}
	if len(bytes) != 5 {
		tester.Error()
	}
	if bytes[0] != 43 {
		tester.Error()
	}
	if bytes[1] != 82 {
		tester.Error()
	}
	if bytes[2] != 115 {
		tester.Error()
	}
	if bytes[3] != 166 {
		tester.Error()
	}
	if bytes[4] != 120 {
		tester.Error()
	}
}

func TestReverseBitsBigEndian3(tester *testing.T) {
	indexes := []uint{201, 730, 252, 646, 750, 1}
	// expectedBytes := []byte{50, 109, 163, 242, 134, 187, 129}

	bytes := ReverseBitsBigEndian(indexes, 8, 10)
	if len(bytes) > 7 {
		extra := bytes[len(bytes)-1]
		bytes = bytes[:len(bytes)-1]
		bytes[len(bytes)-1] += extra
	}

	// fmt.Println()
	// fmt.Println(expectedBytes)
	// fmt.Println(bytes)

	if len(bytes) != 7 {
		tester.Error()
	}
	if bytes[0] != 50 {
		tester.Error()
	}
	if bytes[1] != 109 {
		tester.Error()
	}
	if bytes[2] != 163 {
		tester.Error()
	}
	if bytes[3] != 242 {
		tester.Error()
	}
	if bytes[4] != 134 {
		tester.Error()
	}
	if bytes[5] != 187 {
		tester.Error()
	}
	if bytes[6] != 129 {
		tester.Error()
	}
}

func TestReverseBitsBigEndian4(tester *testing.T) {
	indexes := []uint{384, 276, 567, 655, 843, 36}
	expectedBytes := []byte{96, 17, 72, 222, 143, 210, 228}
	_ = expectedBytes

	bytes := ReverseBitsBigEndian(indexes, 8, 10)

	fmt.Println()
	fmt.Println(expectedBytes)
	fmt.Println(bytes)
}

func TestGetBitBlocksBigEndian2(tester *testing.T) {
	list := []byte{96, 17, 72, 222, 143, 210, 228}
	wordIndexes2 := GetBitBlocksBigEndian(list, 8, 10)

	fmt.Println()
	fmt.Println(wordIndexes2)
}

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
