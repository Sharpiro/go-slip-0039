package secretsharing

import (
	"bytes"
	"fmt"
	"log"

	"../cryptos"
	"../maths"
)

//CreateShares creates shares based off a given secret
func CreateShares(n, k uint, secret []byte) ([]uint, [][]byte) {
	if n < k {
		log.Fatalf("n must be greater than k, secret would be unrecoverable")
	}
	checksummedSecret := getChecksummedSecret(secret)
	secretLen := len(checksummedSecret)
	values := make([][]byte, n)
	for i := range values {
		values[i] = make([]byte, secretLen)
	}

	// for each byte in the secret
	for i := 0; i < secretLen; i++ {
		randomPolynomial := maths.CreateRandomPolynomial(k - 1)
		randomPolynomial[0] = checksummedSecret[i]

		// for each n shares
		for x := uint(1); x <= n; x++ {
			temp := maths.EvaluatePolynomial(randomPolynomial, x)
			values[x-1][i] = byte(temp)
		}
	}

	xValues := make([]uint, n)
	yValues := make([][]byte, n)
	for i := 0; i < int(n); i++ {
		xValues[i] = uint(i) + 1
		yValues[i] = values[i]
	}
	formattedShares := createFormattedShare(xValues, yValues)
	_ = getWordLists(formattedShares)
	return xValues, yValues
}

//RecoverSecret recovers the secret provided by k shares
func RecoverSecret(xValues []uint, yValues [][]byte) []byte {
	numberOfShares := len(yValues)
	if numberOfShares < 2 {
		log.Fatal("need at least two shares to recover a secret")
	}
	secretLength := len(yValues[0])
	csSecret := make([]byte, secretLength)

	// for each byte in the secret
	for i := 0; i < secretLength; i++ {
		subXValues := make([]uint, numberOfShares)
		subYValues := make([]uint, numberOfShares)
		// for each k shares
		for j := 0; j < numberOfShares; j++ {
			subXValues[j] = xValues[j]
			subYValues[j] = uint(yValues[j][i])
		}
		interpolation := maths.LagrangeInterpolate(0, subXValues, subYValues)
		csSecret[i] = byte(interpolation)
	}
	secret := getSecret(csSecret)
	return secret
}

func createFormattedShare(xValues []uint, yValues [][]byte) [][]byte {
	shares := make([][]byte, len(xValues))
	for i := 0; i < len(xValues); i++ {
		index := xValues[i]
		threshold := len(xValues)
		sssPart := yValues[i]
		concatLen := len(sssPart) + 1 + 1 + 2
		concat := make([]byte, 0, concatLen)
		concat = append(concat, byte(index), byte(threshold))
		concat = append(concat, sssPart...)
		checksum := cryptos.GetSha256(concat)[:2]
		concat = append(concat, checksum...)
		shares[i] = concat
	}
	return shares
}

func getWordLists(formattedShares [][]byte) [][]uint {
	wordLists := make([][]uint, len(formattedShares))
	for i := range wordLists {
		wordLists[i] = getWordIndexes(formattedShares[i])
	}
	return wordLists
}

func getWordIndexes(formattedShare []byte) []uint {
	var blockSize uint = 5
	tenBitNumbers := make([]uint, 0, 30)
	// tenBitNumbers[0] = uint(formattedShare[0])
	// tenBitNumbers[1] = uint(formattedShare[1])
	var tenBitNumber uint
	// fmt.Println("Bits, least to most significant:")
	var power uint
	// for i := 2; i < len(formattedShare); i++ {
	for i := 0; i < 2; i++ {
		x := formattedShare[i]
		// fmt.Println("---------------")
		for j := uint(0); j < blockSize; j++ {
			bit := x & (1 << j) >> j
			fmt.Println(bit)
			if bit == 1 {
				tenBitNumber += (1 << power)
			}
			power++
			if power == 10 {
				fmt.Println("---------------")
				tenBitNumbers = append(tenBitNumbers, tenBitNumber)
				power = 0
				tenBitNumber = 0
			}
		}
	}
	if power != 0 {
		tenBitNumbers = append(tenBitNumbers, tenBitNumber)
	}
	return tenBitNumbers
}

func getChecksummedSecret(secret []byte) []byte {
	checksum := cryptos.GetSha256(secret)[:2]
	checksummedSecret := append(secret, checksum...)
	return checksummedSecret
}

func getSecret(csSecret []byte) []byte {
	expectedChecksum := csSecret[len(csSecret)-2:]
	data := csSecret[:len(csSecret)-2]
	actualChecksum := cryptos.GetSha256(data)
	if !bytes.Equal(expectedChecksum, actualChecksum[:2]) {
		log.Fatal("checksums do not match")
	}
	secret := csSecret[:len(csSecret)-2]
	return secret
}
