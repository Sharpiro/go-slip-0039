package cryptos

const customizationString = "shamir"

func rs1024Polymod(values []int) int {
	GEN := []int{0xe0e040, 0x1c1c080, 0x3838100, 0x7070200, 0xe0e0009, 0x1c0c2412,
		0x38086c24, 0x3090fc48, 0x21b1f890, 0x3f3f120}
	chk := 1
	for _, v := range values {
		b := (chk >> 20)
		chk = (chk&0xfffff)<<10 ^ v
		for i := uint(0); i < 10; i++ {
			if (b>>i)&1 > 0 {
				chk ^= GEN[i]
			}
		}
	}
	return int(chk)
}

// RS1024CreateChecksum creates a reed-solomon checksum
func RS1024CreateChecksum(data []int) []int {
	csData := make([]int, len(customizationString))
	for i := 0; i < len(customizationString); i++ {
		csData[i] = int(customizationString[i])
	}
	values := append(csData, data...)
	values = append(values, []int{0, 0, 0}...)
	polymod := rs1024Polymod(values) ^ 1
	_ = polymod

	var checksum []int
	for i := 3 - 1; i >= 0; i-- {
		checksum = append(checksum, (polymod>>uint(10*i))&1023)
	}
	return checksum
}

// RS1024VerifyChecksum verifies a reed-solomon checksum
func RS1024VerifyChecksum(data []int) bool {
	csData := make([]int, len(customizationString))
	for i := 0; i < len(customizationString); i++ {
		csData[i] = int(customizationString[i])
	}
	values := append(csData, data...)
	return rs1024Polymod(values) == 1
}
