package secretsharing

import (
	"bytes"
	"go-slip-0039/cryptos"
	"go-slip-0039/maths"
	"log"
)

// CreateWordShares creates shares based off a given secret
func CreateWordShares(n, k uint, secret []byte) [][]string {
	checksummedSecret := getChecksummedSecret(secret)
	xValues, yValues := createShares(n, k, checksummedSecret)
	formattedShares := createFormattedShares(xValues, yValues, k)
	wordLists := getMnemonicList(formattedShares, len(secret))
	return wordLists
}

// RecoverFromWordShares recovers a secret based off of K supplied word lists
func RecoverFromWordShares(wordLists [][]string, bitLength int) []byte {
	formattedShares := getMnemonicBuffers(wordLists, bitLength)
	xValues, yValues := recoverFromFormattedShare(formattedShares)
	checkSummedSecret := recoverSecret(xValues, yValues)
	secret := getSecret(checkSummedSecret)
	return secret
}

// AnalyzeShare returns useful data about a given share
func AnalyzeShare(share []string, bitLength int) (index, threshold, length int) {
	index, threshold = AnalyzeFirstWord(share[0])
	mnemonicIndexes := getMnemonicIndexes(share)
	_ = getMnemonicBuffer(mnemonicIndexes, bitLength)
	length = len(share[0]) >> 1 // todo: what is this again??
	return index, threshold, length
}

func createShares(n, k uint, secret []byte) ([]uint, [][]byte) {
	if n < k {
		log.Fatalf("n must be greater than k, secret would be unrecoverable")
	}
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

	xValues := make([]uint, n)
	yValues := make([][]byte, n)
	for i := 0; i < int(n); i++ {
		xValues[i] = uint(i) + 1
		yValues[i] = values[i]
	}
	return xValues, yValues
}

func recoverSecret(xValues []uint, yValues [][]byte) []byte {
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
	secret := csSecret
	return secret
}

func recoverFromFormattedShare(shareBlock [][]byte) ([]uint, [][]byte) {
	const encodingOffset byte = 1
	xValues := make([]uint, len(shareBlock))
	yValues := make([][]byte, len(shareBlock))
	for i := 0; i < len(shareBlock); i++ {
		// expectedChecksum := shareBlock[i][len(shareBlock[i])-2:]
		// actualData := shareBlock[i][:len(shareBlock[i])-2]
		// actualChecksum := cryptos.GetSha256(actualData)[:2]
		// if !bytes.Equal(expectedChecksum, actualChecksum) {
		// 	log.Fatal("failed while recovering share, expected checksum does not match actual")
		// }
		index := uint(shareBlock[i][0] + encodingOffset)
		sssPart := shareBlock[i][2 : len(shareBlock[i])-2]
		xValues[i] = index
		yValues[i] = sssPart
	}
	return xValues, yValues
}

func createFormattedShares(xValues []uint, yValues [][]byte, k uint) [][]byte {
	shares := make([][]byte, len(xValues))
	concatLen := len(yValues[0]) + 1 + 1 + 2
	for i := 0; i < len(xValues); i++ {
		index := xValues[i] - 1
		threshold := k - 1
		sssPart := yValues[i]
		concat := make([]byte, 0, concatLen)
		concat = append(concat, byte(index), byte(threshold))
		concat = append(concat, sssPart...)
		checksum := cryptos.GetSha256(concat)[:2]
		concat = append(concat, checksum...)
		shares[i] = concat
	}
	return shares
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
		log.Fatal("actual master secret checksum did not match expected checksum")
	}
	return data
}
