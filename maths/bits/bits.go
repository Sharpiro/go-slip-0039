package bits

// func getWordLists(formattedShares [][]byte) [][]uint {
// 	wordLists := make([][]uint, len(formattedShares))
// 	for i := range wordLists {
// 		wordLists[i] = GetBitBlocks(formattedShares[i], 8, 10)
// 	}
// 	return wordLists
// }

func GetBitBlocksBigEndian(formattedShare []byte, byteSize int, splitSize int) []uint {
	var createdNumbers []uint
	var currentNumber uint
	power := splitSize - 1
	for i := 0; i < len(formattedShare); i++ {
		x := formattedShare[i]
		for j := int(byteSize - 1); j >= 0; j-- {
			bit := x & (1 << uint(j)) >> uint(j)
			// fmt.Println(bit)
			if bit == 1 {
				currentNumber += (1 << uint(power))
			}
			power--
			if power == -1 {
				// fmt.Println("---------------")
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
