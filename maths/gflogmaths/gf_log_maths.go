package gflogmaths

import "go-slip-0039/maths/gfmaths"

// var exp, log []uint = newField(3)
var exp, log []uint = ExponentTable, LogTable

// Add 2 items
func Add(a, b uint) uint {
	return a ^ b
}

// Subtract 2 items
func Subtract(a, b uint) uint {
	return a ^ b
}

// Multiply 2 numbers reduced by a polynomial
func Multiply(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}
	return exp[log[a]+log[b]]
}

// Inverse gets the inverse of a number given a polynomial
func Inverse(x uint) uint {
	if x == 0 {
		return 0
	}
	return exp[255-log[x]]
}

// Divide performs a * 1/b
func Divide(a, b uint) uint {
	inverseB := Inverse(b)
	return Multiply(a, inverseB)
}

func newField(a uint) (exp, log []uint) {
	exp = make([]uint, 512)
	log = make([]uint, 256)
	var x uint = 1
	for i := uint(0); i < 255; i++ {
		exp[i] = x
		exp[i+255] = x
		log[x] = i
		x = gfmaths.Multiply(x, a)
	}
	log[0] = 255
	exp[510] = 0
	exp[511] = 0
	return exp, log
}
