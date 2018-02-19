package maths

import (
	"testing"

	gfArith "./gfmaths"
)

func TestGetPolynomialDegree(tester *testing.T) {
	polynomial := []byte{1, 0, 0, 0, 1, 1, 0, 1, 1}

	if getPolynomialDegree(polynomial) != 8 {
		tester.Errorf("Degree was invalid")
	}
}

func TestCreateRandomPolynomial(tester *testing.T) {
	polynomial := CreateRandomPolynomial(5)

	if getPolynomialDegree(polynomial) != 5 {
		tester.Errorf("Degree was invalid")
	}
}

func TestEvaluatePolynomial(tester *testing.T) {
	bytePolynomial := []byte{123, 166, 94}
	polynomial := []uint{123, 166, 94}
	var prime uint = 0x11b

	for x := uint(0); x < 7; x++ {
		temp1 := gfArith.Multiply(x, x, prime)
		temp2 := gfArith.Multiply(temp1, polynomial[2], prime)
		temp3 := gfArith.Multiply(polynomial[1], x, prime)
		temp4 := gfArith.Add(temp2, temp3)
		expected := gfArith.Add(temp4, polynomial[0])

		actual := EvaluatePolynomial(bytePolynomial, x)
		if expected != actual {
			tester.Errorf("expected %v but was %v", expected, actual)
		}
	}
}

func TestLagrangeInterpolate(tester *testing.T) {
	xValues := []uint{1, 2, 3, 4, 5, 6}
	yValues := []uint{31, 68, 230, 250, 88, 3}
	var expected uint = 189

	actual := LagrangeInterpolate(0, xValues, yValues)
	if expected != actual {
		tester.Errorf("expected %v but was %v", expected, actual)
	}
}
