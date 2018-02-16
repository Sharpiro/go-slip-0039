package secretsharing

import (
	"../maths"
	"log"
)

type tuple struct {
	item1 int
	item2 int
}

//CreateShares creates shares based off a given secret
func CreateShares(n, k uint, secret []byte) ([]uint, [][]byte) {
	secretLen := len(secret)
	values := make([][]byte, n)
	for i := range values {
		values[i] = make([]byte, secretLen)
	}

	// for each byte in the secret
	for i := 0; i < secretLen; i++ {
		randomPolynomial := maths.CreateRandomPolynomial(k - 1)
		randomPolynomial[0] = secret[i]

		// for each n shares
		for x := uint(1); x <= n; x++ {
			temp := maths.EvaluatePolynomial(randomPolynomial, x)
			values[x-1][i] = byte(temp)
		}
	}

	xValues := make([]uint, n, n)
	yValues := make([][]byte, n, n)
	for i := 0; i < int(n); i++ {
		xValues[i] = uint(i) + 1
		yValues[i] = values[i]
	}
	return xValues, yValues
}

//RecoverSecret recovers the secret provided by k shares
func RecoverSecret(xValues []uint, yValues [][]byte) []byte {
	numberOfShares := len(yValues)
	if numberOfShares < 2 {
		log.Fatal("need at least two shares")
	}
	// sharesBytes = list(map(lambda x: (x[0], binascii.unhexlify(x[1])), shares))
	secretLength := len(yValues[0])
	// secret := bytearray([0] * secretLength)
	secret := make([]byte, secretLength, secretLength)

	// for each byte in the secret
	for i := 0; i < secretLength; i++ {
		// values = [[0 for i in range(2)] for j in range(numberOfShares)]
		values := make([][]uint, numberOfShares, numberOfShares)
		for temp := range values {
			values[temp] = make([]uint, 2, 2)
		}

		// for each k shares
		subXValues := make([]uint, numberOfShares, numberOfShares)
		subYValues := make([]uint, numberOfShares, numberOfShares)
		for j := 0; j < numberOfShares; j++ {
			subXValues[j] = xValues[j]
			subYValues[j] = uint(yValues[j][i])
		}
		interpolation := maths.LagrangeInterpolate(0, subXValues, subYValues)
		secret[i] = byte(interpolation)
	}
	return secret
}
