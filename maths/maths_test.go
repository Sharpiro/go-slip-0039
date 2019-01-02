package maths

import (
	"testing"

	gfLogMaths "go-slip-0039/maths/gflogmaths"
	gfArith "go-slip-0039/maths/gfmaths"
)

func TestGfMultiplyAll(tester *testing.T) {
	for i := uint(0x00); i <= 0xff; i++ {
		for j := uint(0x00); j <= 0xff; j++ {
			regResult := gfArith.Multiply(i, j)
			tableResult := gfLogMaths.Multiply(i, j)

			if regResult != tableResult {
				tester.Errorf("regResult %v, tableResult %v", regResult, tableResult)
			}
		}
	}
}

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

	for x := uint(0); x < 7; x++ {
		temp1 := gfArith.Multiply(x, x)
		temp2 := gfArith.Multiply(temp1, polynomial[2])
		temp3 := gfArith.Multiply(polynomial[1], x)
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
