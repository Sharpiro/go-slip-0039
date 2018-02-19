package main

import (
	"fmt"
)

func main() {
	// start := time.Now()
	// elapsed := time.Since(start)
	// print(elapsed)
	list := []byte{3, 6, 12, 254, 188}
	var number uint
	fmt.Println("Bits, least to most significant:")
	var k uint
	for j := 2; j < len(list); j++ {
		x := list[j]
		//fmt.Printf("%b\n", x)
		fmt.Println("---------------")
		fmt.Println(x)
		fmt.Println("---------------")
		for i := uint(0); i < 8; i++ {
			bit := x & (1 << i) >> i
			fmt.Println(bit)
			if bit == 0 {
				k++
				continue
			} else {
				number += (1 << k)
				k++
			}
			if k == 10 {
				fmt.Println(number)
				k = 0
				number = 0
			}
		}
	}
	fmt.Println(number)
}
