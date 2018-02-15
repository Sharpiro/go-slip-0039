package main

import (
	"fmt"
	// maths "./maths/gflogmaths"
	// maths "./maths/gfmaths"
)

func main() {
	// fmt.Println(maths.Subtract(2, 1))
	// fmt.Println(maths.Add(2, 1))

	for i := 512; i < 10000; i++ {
		if i&256 != 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("Done")
}
