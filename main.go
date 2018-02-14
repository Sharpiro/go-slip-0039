package main

import (
	"fmt"

	"./maths"
	"./maths/gfLogMaths"
	"./maths/gfMaths"
)

func main() {
	fmt.Println(maths.Test())
	fmt.Println(gfMaths.TestSubFunc())
	fmt.Println(gfLogMaths.TestSubFunc())
}
