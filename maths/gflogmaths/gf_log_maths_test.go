package gflogmaths

import (
	"testing"
)

func TestMultiplyTwo(tester *testing.T) {
	var a uint = 83
	var b uint = 202
	var expected uint = 1
	actual := Multiply(a, b)
	if expected != actual {
		tester.Errorf("expected %v, actual %v", expected, actual)
	}
	actual = Multiply(b, a)
	if expected != actual {
		tester.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestMultiplyThree(tester *testing.T) {
	var a uint = 0xb6
	var b uint = 0x53
	var expected uint = 0x36
	actual := Multiply(a, b)
	if expected != actual {
		tester.Errorf("expected %v, actual %v", expected, actual)
	}
	actual = Multiply(b, a)
	if expected != actual {
		tester.Errorf("expected %v, actual %v", expected, actual)
	}
}

func TestMultiplyFour(tester *testing.T) {
	var a uint = 0xfd
	var b uint = 0xfd
	var expected uint = 0x17
	actual := Multiply(a, b)
	if expected != actual {
		tester.Errorf("expected %v, actual %v", expected, actual)
	}
	actual = Multiply(b, a)
	if expected != actual {
		tester.Errorf("expected %v, actual %v", expected, actual)
	}
}
