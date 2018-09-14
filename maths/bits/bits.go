package bits

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func GetBits(x byte, padding int) string {
	binString := strconv.FormatInt(int64(x), 2)
	paddedBinary := strings.Repeat("0", padding-len(binString))
	paddedString := paddedBinary + binString
	return paddedString
}

func GetBitsArray(buffer []byte, padding int) string {
	var bits string
	for _, j := range buffer {
		bits = bits + GetBits(j, padding)
	}
	return bits
}

func GetBytes(bits string) []byte {
	bytes := make([]byte, 0)
	for i := 0; i < len(bits); i += 8 {
		data := bits[i : i+8]
		parsed, err := strconv.ParseInt(data, 2, 64)
		if err != nil {
			log.Fatal("failed converting bits to bytes")
		}
		bytes = append(bytes, byte(parsed))
	}
	return bytes
}

func PadShareToNearestTen(share string) string {
	remainder := len(share) % 10
	if remainder == 0 {
		return share
	}
	paddingLength := 10 - remainder
	padding := GetBits(0, paddingLength)
	paddedShare := share + padding
	return paddedShare
}

func PadBits(bits string) string {
	finalPaddedBitsSize := int(math.Ceil(float64(len(bits))/8)) * 8
	remainingPaddingSize := finalPaddedBitsSize - len(bits)
	var paddingBits string
	if remainingPaddingSize > 0 {
		paddingBits = GetBits(0, remainingPaddingSize)
	}

	bits = bits + paddingBits
	return bits
}

// Power2ToHex  Converts vector of integers representing number base 2^p to a byte-vector
// with complexity O( vector.size() )
// power of 2 in a base must be 9 >= x <= 24
func Power2ToHex(indexes []uint, power uint) []byte {
	if power < 9 || power > 24 {
		log.Fatal("base 2-power must be 9 >= x <= 24")
	}
	output := make([]byte, 0)
	var appended uint
	for _, v := range indexes {
		lastleft := (8 - appended%8) % 8
		toappend := power
		if lastleft != 0 {
			output[len(output)-1] |= byte(v >> (power - lastleft))
			toappend -= lastleft
			appended += lastleft
		}
		for toappend >= 8 {
			output = append(output, byte(v>>(toappend-8)))
			toappend -= 8
			appended += 8
		}
		if toappend != 0 {
			output = append(output, byte(v<<(8-toappend)))
			appended += toappend
		}
	}
	return output
}

// HexToPower2  Converts vector of bytes into array of integers representing number base 2^p
// with complexity O( vector.size() )
// power of 2 in a base must be 9 >= x <= 24
func HexToPower2(data []byte, p uint) []uint {
	if p < 9 || p > 24 {
		log.Fatal("base 2-power must be 9 >= x <= 24")
	}
	output := make([]uint, 0)
	var bitholder uint
	var bitsread uint
	for _, x := range data {
		var willread uint
		if willread = 8; (p - bitsread) <= 8 {
			willread = p - bitsread
		}
		bitholder <<= willread
		bitholder |= uint((x >> (8 - willread)))
		bitsread += willread
		if bitsread == p {
			output = append(output, bitholder)
			bitholder = uint(x & (0xff >> (willread)))
			bitsread = 8 - willread
		}
	}
	last := bitholder << (p - bitsread)
	output = append(output, last)
	return output
}

// ResizeWordIndex resizes a base 1024 array based upon the original entropy size
func ResizeWordIndex(data []uint, shareSizeBits int) []uint {
	// lineBits := (shareSizeBits*8 + 42)
	var xyz int
	if xyz = 0; shareSizeBits%10 != 0 {
		xyz = 1
	}
	newSize := shareSizeBits/10 + xyz
	data = data[:newSize]
	return data
}

// ResizeBytes resizes bytes
func ResizeBytes(data []byte, entropySizeBytes int) []byte {
	lineBits := float64(entropySizeBytes*8 + 42)
	fullByteSize := int(math.Ceil(lineBits / 8))
	return data[:fullByteSize]
}

// ReverseBitsBigEndian returns a byte array
func ReverseBitsBigEndian(indexes []uint, byteSize, splitSize, totalBits int) []byte {
	createdBytes := make([]byte, 0, len(indexes))
	var currentByte uint
	power := byteSize - 1
	for i, x := range indexes {
		for j := getJ(i, len(indexes), splitSize, byteSize, totalBits); j >= 0; j-- {
			bit := x & (1 << uint(j)) >> uint(j)
			// fmt.Print(bit)
			if bit == 1 {
				currentByte += (1 << uint(power))
			}
			totalBits--
			power--
			if power == -1 {
				// fmt.Println("\n---------------")
				createdBytes = append(createdBytes, byte(currentByte))
				power = byteSize - 1
				currentByte = 0
			}
		}
	}
	return createdBytes
}

func getJ(i, len, splitSize, byteSize, bitsRemaining int) int {
	if len > 1 && i+1 == len && bitsRemaining < byteSize {
		return bitsRemaining - 1
	}
	return splitSize - 1
}

// GetBitBlocksBigEndian gets numbers in little endian format
func GetBitBlocksBigEndian(formattedShare []byte, byteSize int, splitSize int) []uint {
	var createdNumbers []uint
	var currentNumber uint
	power := splitSize - 1
	for i, x := range formattedShare {
		for j := byteSize - 1; j >= 0; j-- {
			bit := x & (1 << uint(j)) >> uint(j)
			// fmt.Print(bit)
			if bit == 1 {
				currentNumber += (1 << uint(power))
			}
			power--
			if power == -1 {
				// fmt.Println("\n---------------")
				createdNumbers = append(createdNumbers, currentNumber)
				if i+1 == len(formattedShare) {
					power = j - 1
				} else {
					power = splitSize - 1
				}
				currentNumber = 0
			}
		}
	}
	if power != -1 {
		createdNumbers = append(createdNumbers, currentNumber)
	}
	return createdNumbers
}

// GetBitBlocksLittleEndian gets numbers in big endian format
func GetBitBlocksLittleEndian(formattedShare []byte, byteSize uint, splitSize int) []uint {
	createdNumbers := make([]uint, 0, 30)
	var currentNumber uint
	power := splitSize - 1
	for i := 0; i < len(formattedShare); i++ {
		x := formattedShare[i]
		for j := uint(0); j < byteSize; j++ {
			bit := x & (1 << j) >> j
			// fmt.Print(bit)
			if bit == 1 {
				currentNumber += (1 << uint(power))
			}
			power--
			if power == -1 {
				// fmt.Println("\n---------------")
				createdNumbers = append(createdNumbers, currentNumber)
				if i+1 == len(formattedShare) {
					power = int(byteSize - j - 2)
				} else {
					power = splitSize - 1
				}
				currentNumber = 0
			}
		}
	}
	if power != -1 {
		createdNumbers = append(createdNumbers, currentNumber)
	}
	return createdNumbers
}
