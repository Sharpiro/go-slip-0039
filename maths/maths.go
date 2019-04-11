package maths

import (

	// gfArith "go-slip-0039/maths/gfmaths"
	"go-slip-0039/cryptos"
	gfArith "go-slip-0039/maths/gflogmaths"
)

// Point represents a point on a polynomial
type Point struct {
	X uint
	Y []byte
}

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
			numerator := gfArith.Subtract(xInput, xValues[j])       // can probably remove xinput subtraction
			denominator := gfArith.Subtract(xValues[i], xValues[j]) // these 2 values can be probably be swapped w/o issue
			newLi := gfArith.Divide(numerator, denominator)
			li = gfArith.Multiply(li, newLi)
		}
		l := gfArith.Multiply(li, yi)
		y = gfArith.Add(y, l)
	}
	return y
}

// LagrangeInterpolateNew is used to rebuild the original polynomial, and thus the secret
func LagrangeInterpolateNew(xInput uint, points []Point) []byte {
	y := make([]byte, len(points[0].Y))
	for i := 0; i < len(points); i++ {
		var li uint = 1
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			numerator := gfArith.Subtract(xInput, points[j].X)        // can probably remove xinput subtraction
			denominator := gfArith.Subtract(points[i].X, points[j].X) // these 2 values can be probably be swapped w/o issue
			newLi := gfArith.Divide(numerator, denominator)
			li = gfArith.Multiply(li, newLi)
		}
		l := gfArith.MultiplyX(li, points[i].Y)
		y = gfArith.AddX(y, l)
	}
	return y
}

// CreateRandomPolynomial creates a random polynomial of the given degree
func CreateRandomPolynomial(degree uint) []byte {
	randomBytes := cryptos.GetRandomBytes(int(degree) + 1)
	// randomBytes := make([]byte, degree+1, degree+1)
	for getPolynomialDegree(randomBytes) != degree {
		randomBytes = cryptos.GetRandomBytes(int(degree) + 1)
		// if _, err := rand.Read(randomBytes); err != nil {
		// 	log.Fatal("an error occurred generating random bytes")
		// }
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
