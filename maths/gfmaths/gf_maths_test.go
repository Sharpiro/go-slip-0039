package gfmaths

import (
	"testing"
)

func TestAdd(tester *testing.T) {
	var expected uint = 3
	actual := Add(1, 2)
	if expected != actual {
		tester.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestSubtract(tester *testing.T) {
	var expected uint = 3
	actual := Subtract(1, 2)
	if expected != actual {
		tester.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestMultiplyOne(tester *testing.T) {
	var p uint = 355
	var a uint = 84
	var b uint = 13
	var expected uint = 1
	actual := Multiply(a, b, p)
	if expected != actual {
		tester.Errorf("expected %v, actual %v", expected, actual)
	}
	actual = Multiply(b, a, p)
	if expected != actual {
		tester.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestMultiplyTwo(tester *testing.T) {
	var p uint = 0x11b
	var a uint = 83
	var b uint = 202
	var expected uint = 1
	actual := Multiply(a, b, p)
	if expected != actual {
		tester.Errorf("expected %v, actual %v", expected, actual)
	}
	actual = Multiply(b, a, p)
	if expected != actual {
		tester.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestDividePolynomials(tester *testing.T) {
	quotient, remainder := DividePolynomials(425, 51)
	if 9 != quotient {
		tester.Error()
	}
	if 2 != remainder {
		tester.Error()
	}

	quotient, remainder = DividePolynomials(355, 84)
	if 4 != quotient {
		tester.Error()
	}
	if 51 != remainder {
		tester.Error()
	}

	quotient, remainder = DividePolynomials(84, 51)
	if 3 != quotient {
		tester.Error()
	}
	if 1 != remainder {
		tester.Error()
	}

	quotient, remainder = DividePolynomials(0x11b, 83)
	if 5 != quotient {
		tester.Error()
	}
	if 4 != remainder {
		tester.Error()
	}

	quotient, remainder = DividePolynomials(83, 4)
	if 20 != quotient {
		tester.Error()
	}
	if 3 != remainder {
		tester.Error()
	}

	quotient, remainder = DividePolynomials(4, 3)
	if 3 != quotient {
		tester.Error()
	}
	if 1 != remainder {
		tester.Error()
	}

	quotient, remainder = DividePolynomials(3, 1)
	if 3 != quotient {
		tester.Error()
	}
	if 0 != remainder {
		tester.Error()
	}
}

func TestInverse(tester *testing.T) {
	var b uint = 202
	var a uint = 83
	var p uint = 0x11b
	mult := Multiply(a, b, p)
	if mult != 1 {
		tester.Error()
	}
	inverse := Inverse(a, p)
	if inverse != b {
		tester.Error()
	}
	inverse = Inverse(b, p)
	if inverse != a {
		tester.Error()
	}
}

func TestInverseTwo(tester *testing.T) {
	var b uint = 13
	var a uint = 84
	var p uint = 355
	mult := Multiply(a, b, p)
	if mult != 1 {
		tester.Error()
	}
	inverse := Inverse(a, p)
	if inverse != b {
		tester.Error()
	}
	inverse = Inverse(b, p)
	if inverse != a {
		tester.Error()
	}
}
