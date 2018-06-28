package bits

import (
	"reflect"
	"testing"
)

func TestGetBits(tester *testing.T) {
	indexList := []byte{11, 10, 5, 4, 97, 219}
	bitsx := GetBits(indexList[0], 8)
	_ = bitsx
}

func TestGetBytes(tester *testing.T) {
	bits := "0000000001000010110000101000000101000001000110000111011011000000"
	size := 58
	bytes := GetBytes(bits, size)
	_ = bytes
}

func TestIntToByteConversion(tester *testing.T) {
	indexList := []uint{130, 512, 612, 227, 732, 733, 437, 512}

	// 32, 160 9, 144, 227, 183, 45, 214, 214, 0
	theTemp := Power2ToHex(indexList, 10)
	andBack := ResizeWordIndex(HexToPower2(theTemp, 10), 4)
	_ = andBack

	if !reflect.DeepEqual(indexList, andBack) {
		tester.Error("rebuilt index list did not match original")
	}

	indexList = []uint{1, 737, 385, 67, 990, 739, 913, 64}
	theTemp = Power2ToHex(indexList, 10)
	andBack = ResizeWordIndex(HexToPower2(theTemp, 10), 4)
	if !reflect.DeepEqual(indexList, andBack) {
		tester.Error("rebuilt index list did not match original")
	}
}

func TestEntropy128BitTest(tester *testing.T) {
	// 	acid glance scatter multiply muscle evolve vote hedgehog vanish shoe road sense ugly raise sister scout educate
	// anger mansion second exclude grow garden video purchase cost skin crowd surface brush choice machine lock shed
	// axis clerk cupboard golden endless minute army lecture fuel soldier peace regret deny extra group execute menu
	hexEntropy := "c9d32bb6f9a2024b9e12c2cd4af717c1"
	entropySize := len(hexEntropy) / 2 // 128

	indexList := []uint{1, 400, 766, 582, 583, 312, 984, 430, 960, 795, 745, 785, 935, 706, 806, 772, 274}
	bytes := Power2ToHex(indexList, 10)
	rebuiltIndexList := HexToPower2(bytes, 10)
	resizedIndexList := ResizeWordIndex(rebuiltIndexList, entropySize)

	if !reflect.DeepEqual(indexList, resizedIndexList) {
		tester.Error("rebuilt index list did not match original")
	}
}

func TestReverseBitsBigEndian(tester *testing.T) {
	indexes := []uint{102, 20, 175, 1009, 3}
	preBytes := ReverseBitsBigEndian(indexes[:1], 5, 10, 16)
	bytes := ReverseBitsBigEndian(indexes[1:5], 8, 10, 32)

	// fmt.Println(preBytes)
	// fmt.Println(bytes)

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
	expectedBytes := []byte{0, 5, 43, 82, 115, 166, 120}
	_ = expectedBytes

	preBytes := ReverseBitsBigEndian(indexes[:1], 5, 10, 16)
	bytes := ReverseBitsBigEndian(indexes[1:5], 8, 10, 40)

	// fmt.Println()
	// fmt.Println(preBytes)
	// fmt.Println(bytes)

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

func TestGetBitBlocksBigEndian3(tester *testing.T) {
	expectedIndexes := []uint{201, 730, 252, 646, 750, 1}
	_ = expectedIndexes

	bytes := []byte{50, 109, 163, 242, 134, 187, 129}
	wordIndexes := GetBitBlocksBigEndian(bytes, 8, 10)

	// fmt.Println()
	// fmt.Println(wordIndexes)

	if len(wordIndexes) != 6 {
		tester.Error()
	}
	if wordIndexes[0] != 201 {
		tester.Error()
	}
	if wordIndexes[1] != 730 {
		tester.Error()
	}
	if wordIndexes[2] != 252 {
		tester.Error()
	}
	if wordIndexes[3] != 646 {
		tester.Error()
	}
	if wordIndexes[4] != 750 {
		tester.Error()
	}
	if wordIndexes[5] != 1 {
		tester.Error()
	}
}

func TestReverseBitsBigEndian3(tester *testing.T) {
	indexes := []uint{201, 730, 252, 646, 750, 1}
	expectedBytes := []byte{50, 109, 163, 242, 134, 187, 129}
	_ = expectedBytes

	bytes := ReverseBitsBigEndian(indexes, 8, 10, 56)

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

	bytes := ReverseBitsBigEndian(indexes, 8, 10, 56)

	// fmt.Println()
	// fmt.Println(expectedBytes)
	// fmt.Println(bytes)

	if len(bytes) != 7 {
		tester.Error()
	}
	if bytes[0] != 96 {
		tester.Error()
	}
	if bytes[1] != 17 {
		tester.Error()
	}
	if bytes[2] != 72 {
		tester.Error()
	}
	if bytes[3] != 222 {
		tester.Error()
	}
	if bytes[4] != 143 {
		tester.Error()
	}
	if bytes[5] != 210 {
		tester.Error()
	}
	if bytes[6] != 228 {
		tester.Error()
	}
}

func TestGetBitBlocksBigEndian2(tester *testing.T) {
	expectedIndexes := []uint{384, 276, 567, 655, 843, 36}
	_ = expectedIndexes

	list := []byte{96, 17, 72, 222, 143, 210, 228}
	wordIndexes := GetBitBlocksBigEndian(list, 8, 10)

	// fmt.Println()
	// fmt.Println(wordIndexes)

	if len(wordIndexes) != 6 {
		tester.Error()
	}
	if wordIndexes[0] != 384 {
		tester.Error()
	}
	if wordIndexes[1] != 276 {
		tester.Error()
	}
	if wordIndexes[2] != 567 {
		tester.Error()
	}
	if wordIndexes[3] != 655 {
		tester.Error()
	}
	if wordIndexes[4] != 843 {
		tester.Error()
	}
	if wordIndexes[5] != 36 {
		tester.Error()
	}
}

func TestGetBitBlocksBigEndian5(tester *testing.T) {
	expectedIndexes := []uint{892, 412, 913, 843, 864}
	_ = expectedIndexes

	bytes := []byte{223, 25, 206, 71, 75, 216}
	wordIndexes := GetBitBlocksBigEndian(bytes, 8, 10)

	// fmt.Println()
	// fmt.Println(expectedIndexes)
	// fmt.Println(wordIndexes)

	if len(wordIndexes) != 5 {
		tester.Error()
	}
	if wordIndexes[0] != 892 {
		tester.Error()
	}
	if wordIndexes[1] != 412 {
		tester.Error()
	}
	if wordIndexes[2] != 913 {
		tester.Error()
	}
	if wordIndexes[3] != 843 {
		tester.Error()
	}
	if wordIndexes[4] != 864 {
		tester.Error()
	}
}

func TestReverseBitsBigEndian5(tester *testing.T) {
	indexes := []uint{892, 412, 913, 843, 864}
	expectedBytes := []byte{223, 25, 206, 71, 75, 216}
	_ = expectedBytes

	bytes := ReverseBitsBigEndian(indexes, 8, 10, 48)

	// fmt.Println()
	// fmt.Println(expectedBytes)
	// fmt.Println(bytes)

	if len(bytes) != 6 {
		tester.Error()
	}
	if bytes[0] != 223 {
		tester.Error()
	}
	if bytes[1] != 25 {
		tester.Error()
	}
	if bytes[2] != 206 {
		tester.Error()
	}
	if bytes[3] != 71 {
		tester.Error()
	}
	if bytes[4] != 75 {
		tester.Error()
	}
	if bytes[5] != 216 {
		tester.Error()
	}
}

func TestGetBitBlocksBigEndian(tester *testing.T) {
	preBytes := []byte{3, 6}
	bytes := []byte{5, 10, 255, 199} // 0b00000101 0b00001010 0b11111111 0b11000111
	preIndexes := GetBitBlocksBigEndian(preBytes, 5, 10)
	indexes := GetBitBlocksBigEndian(bytes, 8, 10)

	// fmt.Println(wordIndexes)
	// fmt.Println(wordIndexes2) // 20, 175, 1009, 3

	if len(preIndexes) != 1 {
		tester.Error()
	}
	if preIndexes[0] != 102 { //00011 00110
		tester.Error()
	}
	if len(indexes) != 4 {
		tester.Error()
	}
	if indexes[0] != 20 {
		tester.Error()
	}
	if indexes[1] != 175 {
		tester.Error()
	}
	if indexes[2] != 1009 {
		tester.Error()
	}
	if indexes[3] != 3 {
		tester.Error()
	}
}

func TestGetBitBlocksLittleEndian(tester *testing.T) {
	preBytes := []byte{3, 6}
	bytes := []byte{5, 10, 255, 199} // 0b00000101 0b00001010 0b11111111 0b11000111
	preIndexes := GetBitBlocksLittleEndian(preBytes, 5, 10)
	indexes := GetBitBlocksLittleEndian(bytes, 8, 10)

	// fmt.Println()
	// fmt.Println(wordIndexes)
	// fmt.Println(wordIndexes2) // 641, 271, 1016, 3

	if len(preIndexes) != 1 {
		tester.Error()
	}
	if preIndexes[0] != 780 { //00011 00110
		tester.Error()
	}
	if len(indexes) != 4 {
		tester.Error()
	}
	if indexes[0] != 641 {
		tester.Error()
	}
	if indexes[1] != 271 {
		tester.Error()
	}
	if indexes[2] != 1016 {
		tester.Error()
	}
	if indexes[3] != 3 {
		tester.Error()
	}
}
