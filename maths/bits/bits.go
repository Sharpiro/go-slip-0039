package bits

import "fmt"

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

// ReverseBitsBigEndian returns a byte array
func ReverseBitsBigEndian(indexes []uint, byteSize int, splitSize int) []byte {
	createdBytes := make([]byte, 0, len(indexes))
	var currentByte uint
	power := byteSize - 1
	for i, x := range indexes {
		for j := splitSize - 1; j >= 0; j-- {
			bit := x & (1 << uint(j)) >> uint(j)
			fmt.Print(bit)
			if bit == 1 {
				currentByte += (1 << uint(power))
			}
			power--
			if power == -1 {
				fmt.Println("\n---------------")
				if i+1 == len(indexes) {
					power = j - 1
				} else {
					createdBytes = append(createdBytes, byte(currentByte))
					power = byteSize - 1
				}
				// power = byteSize - 1
				currentByte = 0
			}
		}
	}
	if power != -1 {
		// currentByte += (1 << uint(power))
		createdBytes = append(createdBytes, byte(currentByte))
	}
	return createdBytes
}

// GetBitBlocksBigEndian gets numbers in little endian format
func GetBitBlocksBigEndian(formattedShare []byte, byteSize int, splitSize int) []uint {
	var createdNumbers []uint
	var currentNumber uint
	power := splitSize - 1
	for i, x := range formattedShare {
		for j := byteSize - 1; j >= 0; j-- {
			bit := x & (1 << uint(j)) >> uint(j)
			fmt.Print(bit)
			if bit == 1 {
				currentNumber += (1 << uint(power))
			}
			power--
			if power == -1 {
				fmt.Println("\n---------------")
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
