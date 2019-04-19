package maths

import (
	"bytes"
	"testing"

	gfLogMaths "go-slip-0039/maths/gflogmaths"
	gfArith "go-slip-0039/maths/gfmaths"
)

func TestLagrangeInterpolateNew(tester *testing.T) {
	var xInput byte = 1
	points := []Point{
		Point{X: 0, Y: []byte{230, 163, 161, 72, 7, 200, 177, 202, 45, 3, 226, 89, 184, 109, 151, 15}},
		Point{X: 254, Y: []byte{113, 103, 248, 79, 228, 213, 132, 41, 200, 0, 220, 174, 220, 67, 30, 226}},
		Point{X: 255, Y: []byte{248, 138, 218, 228, 36, 53, 166, 246, 55, 69, 55, 115, 40, 139, 2, 202}},
	}
	expected := []byte{111, 78, 131, 227, 199, 40, 147, 21, 210, 70, 9, 132, 76, 165, 139, 39}

	result := LagrangeInterpolate(xInput, points)
	if !bytes.Equal(expected, result) {
		tester.Error(result)
	}
}

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

	actual := LagrangeInterpolateOld(0, xValues, yValues)
	if expected != actual {
		tester.Errorf("expected %v but was %v", expected, actual)
	}
}
