package gfmaths

import (
	"math"
)

type ecdTableEntry struct {
	quotient  uint
	auxillary uint
}

const prime uint = 0x11b

func Add(a, b uint) uint {
	return a ^ b
}

func Subtract(a, b uint) uint {
	return a ^ b
}

// func Multiply(a, b, prime uint) uint {
// 	var p uint = 0
// 	for a > 0 && b > 0 {
// 		if b&1 == 1 {
// 			p = p ^ a
// 		}
// 		if a&0x80 >= 128 {
// 			a = (a << 1) ^ prime
// 		} else {
// 			a <<= 1
// 		}
// 		b >>= 1
// 	}
// 	return p
// }

func Multiply(a, b, poly uint) uint {
	var z uint = 0
	for a > 0 {
		if a&1 == 1 {
			z ^= b
		}
		a >>= 1
		b <<= 1
		if b&0x100 != 0 { // if b >= 256
			b ^= poly
		}
	}
	return z
}

func DividePolynomials(dividend, divisor uint) (uint, uint) {
	var quotient uint = 0
	var remainder uint = 0
	var dividendIndex uint = 0
	minDivisorPower := math.Floor(math.Log2(float64(divisor)))
	minDivisorValue := uint(math.Pow(2, minDivisorPower))
	maxDividendPower := uint(math.Ceil(math.Log2(float64(dividend + 1))))

	for dividendIndex < maxDividendPower {
		for remainder < minDivisorValue && dividendIndex < maxDividendPower {
			bit := getBitAtPosition(dividend, maxDividendPower-dividendIndex-1)
			if bit == 1 {
				remainder = (remainder << 1) + 1
			} else {
				remainder = remainder << 1
			}
			quotient <<= 1
			dividendIndex += 1
			if remainder >= minDivisorValue {
				remainder = remainder ^ divisor
				quotient += 1
			}
		}
	}
	return quotient, remainder
}

func Inverse(a, p uint) uint {
	n := 2
	quotientAuxillary := []ecdTableEntry{ecdTableEntry{0, 0}, ecdTableEntry{0, 1}}
	remainder := a
	dividend := p
	divisor := a
	var newAux uint = 1
	var quotient uint = 0

	for remainder != 1 {
		quotient, remainder = DividePolynomials(dividend, divisor)
		twoOldAux := quotientAuxillary[n-2].auxillary
		oneOldAux := quotientAuxillary[n-1].auxillary
		newAux = Add(twoOldAux, Multiply(oneOldAux, quotient, p))

		quotientAuxillary = append(quotientAuxillary, ecdTableEntry{quotient, newAux})
		dividend = divisor
		divisor = remainder
		n += 1
	}
	return newAux
}

func getBitAtPosition(number, position uint) uint {
	return (number >> position) & 1
}
