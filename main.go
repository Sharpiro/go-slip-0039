package main

import (
	"fmt"
)

func main() {
	// start := time.Now()
	// elapsed := time.Since(start)
	// print(elapsed)
	list := []byte{3, 6, 12, 254, 188}
	const byteSize uint = 8
	tenBitNumbers := make([]uint, 0, 30)
	var tenBitNumber uint
	// fmt.Println("Bits, least to most significant:")
	var power uint
	for i := 2; i < len(list); i++ {
		x := list[i]
		//fmt.Printf("%b\n", x)
		// fmt.Println("---------------")
		// fmt.Println(x)
		// fmt.Println("---------------")
		for j := uint(0); j < byteSize; j++ {
			bit := x & (1 << j) >> j
			// fmt.Println(bit)
			// if bit == 1 {
			// 	tenBitNumber += (1 << power)
			// 	power++
			// } else {
			// 	power++
			// 	continue
			// }
			if bit == 0 {
				power++
				continue
			} else {
				tenBitNumber += (1 << power)
				power++
			}
			if power == 10 {
				// fmt.Println(tenBitNumber)
				tenBitNumbers = append(tenBitNumbers, tenBitNumber)
				power = 0
				tenBitNumber = 0
			}
		}
	}
	tenBitNumbers = append(tenBitNumbers, tenBitNumber)
	fmt.Println(tenBitNumbers)
}
