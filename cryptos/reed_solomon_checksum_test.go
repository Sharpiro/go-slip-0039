package cryptos

import (
	"testing"
)

func TestReedSolomon(tester *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	checksum := RS1024CreateChecksum(data)
	validChecksum := RS1024VerifyChecksum(append(data, checksum...))

	if !validChecksum {
		tester.Error("checksums did not match")
	}
}
