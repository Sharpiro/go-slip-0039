package bits

import (
	"bytes"
	"testing"
)

func TestGetBits(tester *testing.T) {
	var testByte byte = 11
	bits := GetBits(testByte, 8)
	if bits != "00001011" {
		tester.Error()
	}
}

func TestGetBitsArray(tester *testing.T) {
	rawShare := []byte{11, 10, 5, 4, 97, 219}
	bits := GetBitsArray(rawShare, 8)
	if bits != "000010110000101000000101000001000110000111011011" {
		tester.Error()
	}
}

func TestGetBytes(tester *testing.T) {
	bits := "0000000001000010110000101000000101000001000110000111011011000000"
	expectedBytes := []byte{0, 66, 194, 129, 65, 24, 118, 192}
	actualBytes := GetBytes(bits)
	if !bytes.Equal(expectedBytes, actualBytes) {
		tester.Error()
	}
}

func TestGetBytesUneven(tester *testing.T) {
	bits := "0101010101"
	paddedBits := PadBits(bits)
	actualBytes := GetBytes(paddedBits)
	expectedBytes := []byte{85, 64}
	if !bytes.Equal(expectedBytes, actualBytes) {
		tester.Error(0)
	}
}

func TestPadBits(tester *testing.T) {
	bits := "0101010101" // size 10
	paddedBits := PadBits(bits)
	if len(paddedBits) != 16 {
		tester.Error()
	}
}

func TestPadBitsNoChange(tester *testing.T) {
	bits := "0000000001000010" // size 16
	paddedBits := PadBits(bits)
	if len(paddedBits) != 16 {
		tester.Error()
	}
}

func TestPadBitsToNearestTenSize10(tester *testing.T) {
	bits := "0101010101" // size 10
	paddedBits := PadShareToNearestTen(bits)
	if len(paddedBits) != 10 {
		tester.Error()
	}
}

func TestPadBitsToNearestTenSize8(tester *testing.T) {
	bits := "01010111" // size 8
	paddedBits := PadShareToNearestTen(bits)
	if len(paddedBits) != 10 {
		tester.Error()
	}
}

func TestPadBitsToNearestTenSize16(tester *testing.T) {
	bits := "0101011101010111" // size 8
	paddedBits := PadShareToNearestTen(bits)
	if len(paddedBits) != 20 {
		tester.Error()
	}
}

// func TestIntToByteConversion(tester *testing.T) {
// 	indexList := []uint{130, 512, 612, 227, 732, 733, 437, 512}

// 	// 32, 160 9, 144, 227, 183, 45, 214, 214, 0
// 	theTemp := Power2ToHex(indexList, 10)
// 	andBack := ResizeWordIndex(HexToPower2(theTemp, 10), 4)
// 	_ = andBack

// 	if !reflect.DeepEqual(indexList, andBack) {
// 		tester.Error("rebuilt index list did not match original")
// 	}

// 	indexList = []uint{1, 737, 385, 67, 990, 739, 913, 64}
// 	theTemp = Power2ToHex(indexList, 10)
// 	andBack = ResizeWordIndex(HexToPower2(theTemp, 10), 4)
// 	if !reflect.DeepEqual(indexList, andBack) {
// 		tester.Error("rebuilt index list did not match original")
// 	}
// }

// func TestEntropy128BitTest(tester *testing.T) {
// 	// 	acid glance scatter multiply muscle evolve vote hedgehog vanish shoe road sense ugly raise sister scout educate
// 	// anger mansion second exclude grow garden video purchase cost skin crowd surface brush choice machine lock shed
// 	// axis clerk cupboard golden endless minute army lecture fuel soldier peace regret deny extra group execute menu
// 	hexEntropy := "c9d32bb6f9a2024b9e12c2cd4af717c1"
// 	entropySize := len(hexEntropy) / 2 // 128

// 	indexList := []uint{1, 400, 766, 582, 583, 312, 984, 430, 960, 795, 745, 785, 935, 706, 806, 772, 274}
// 	bytes := Power2ToHex(indexList, 10)
// 	rebuiltIndexList := HexToPower2(bytes, 10)
// 	resizedIndexList := ResizeWordIndex(rebuiltIndexList, entropySize)

// 	if !reflect.DeepEqual(indexList, resizedIndexList) {
// 		tester.Error("rebuilt index list did not match original")
// 	}
// }
