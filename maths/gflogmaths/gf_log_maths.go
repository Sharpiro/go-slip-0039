package gflogmaths

import "go-slip-0039/maths/gfmaths"

// var exp, log []uint = newField(3)
var exp, log []uint = exponentTable, logTable

// Add 2 items
func Add(a, b uint) uint {
	return a ^ b
}

// AddX 2 items
func AddX(a, b []byte) []byte {
	temp := make([]byte, len(b))
	for i := 0; i < len(b); i++ {
		temp[i] = a[i] ^ b[i]
	}
	return temp
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
	logA := log[a]
	logB := log[b]
	exp := exp[logA+logB]
	return exp
}

// MultiplyX 2 numbers reduced by a polynomial
func MultiplyX(a uint, b []byte) []byte {
	if a == 0 {
		return []byte{}
	}

	temp := make([]byte, len(b))
	for i := 0; i < len(b); i++ {
		temp[i] = byte(Multiply(a, uint(b[i])))
	}
	// logA := log[a]
	// logB := log[b]
	// exp := exp[logA+logB]
	return temp
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
