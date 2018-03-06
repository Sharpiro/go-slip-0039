package maths

import (
	"crypto/rand"
	"log"

	// gfArith "go-slip-0039/maths/gfmaths"
	gfArith "go-slip-0039/maths/gflogmaths"
)

// LagrangeInterpolate is used to rebuild the original polynomial, and thus the secret
func LagrangeInterpolate(xInput uint, xValues []uint, yValues []uint) uint {
	var y uint
	for i := 0; i < len(xValues); i++ {
		var li uint = 1
		yi := yValues[i]
		for j := 0; j < len(xValues); j++ {
			if i == j {
				continue
			}
			numerator := gfArith.Subtract(xInput, xValues[j])
			denominator := gfArith.Subtract(xValues[i], xValues[j])
			newLi := gfArith.Divide(numerator, denominator)
			li = gfArith.Multiply(li, newLi)
		}
		l := gfArith.Multiply(li, yi)
		y = gfArith.Add(y, l)
	}
	return y
}

// CreateRandomPolynomial creates a random polynomial of the given degree
func CreateRandomPolynomial(degree uint) []byte {
	randomBytes := make([]byte, degree+1, degree+1)
	for getPolynomialDegree(randomBytes) != degree {
		if _, err := rand.Read(randomBytes); err != nil {
			log.Fatal("an error occurred generating random bytes")
		}
	}
	return randomBytes
}

// EvaluatePolynomial evaluates a polynomial for a given x value
func EvaluatePolynomial(polynomial []byte, x uint) uint {
	var result uint
	for i := len(polynomial) - 1; i >= 0; i-- {
		product := gfArith.Multiply(result, x)
		result = gfArith.Add(product, uint(polynomial[i]))
	}
	return result
}

func getPolynomialDegree(poly []byte) uint {
	for i := len(poly) - 1; i >= 0; i-- {
		if poly[i] != 0 {
			return uint(i)
		}
	}
	return 0
}
