package secretsharing

import (
	"bytes"
	"go-slip-0039/cryptos"
	"go-slip-0039/maths"
	"go-slip-0039/maths/bits"
	"go-slip-0039/wordencoding"
	"log"
	"strconv"
)

// CreateMnemonicWordsList creates shares based off a given secret
func CreateMnemonicWordsList(n, k uint, secret []byte) [][]string {
	checksummedSecret := createChecksummedSecret(secret)
	xValues, yValues := createShamirData(n, k, checksummedSecret)

	var mnemonicWordsList [][]string
	for i := 0; i < len(xValues); i++ {
		unchecksummedShare := createUnchecksummedShare(yValues[i], xValues[i], k)
		checksummedShare := unchecksummedShare.GetChecksummedBuffer()
		indexList := wordencoding.CreateIndexList(checksummedShare, len(secret))
		mnemonicWords := wordencoding.CreateMnemonicWords(indexList)
		mnemonicWordsList = append(mnemonicWordsList, mnemonicWords)
	}
	return mnemonicWordsList
}

// RecoverSecretFromMnemonicShares recovers a secret based off of K supplied word lists
func RecoverSecretFromMnemonicShares(mnemonicWordsList [][]string, secretSizeBytes int) []byte {
	var xValues []uint
	var yValues [][]byte

	index, secretThreshold, shamirBuffer := RecoverShare(mnemonicWordsList[0], secretSizeBytes)
	xValues = append(xValues, index)
	yValues = append(yValues, shamirBuffer)
	for i := 1; i < len(mnemonicWordsList); i++ {
		mnemonicWords := mnemonicWordsList[i]
		index, shareThreshold, shamirBuffer := RecoverShare(mnemonicWords, secretSizeBytes)
		if shareThreshold != secretThreshold {
			log.Fatalf("the share's threshold '%v', did not match the first share's threshold '%v'", shareThreshold, secretThreshold)
		}
		xValues = append(xValues, index)
		yValues = append(yValues, shamirBuffer)
	}
	secret := RecoverSecretFromShamirData(xValues, yValues)
	return secret
}

// RecoverSecretFromShamirData recovers a secret based off of the raw shamir data
func RecoverSecretFromShamirData(xValues []uint, yValues [][]byte) []byte {
	indexMap := make(map[uint]bool, len(xValues))
	for i := 0; i < len(xValues); i++ {
		index := xValues[i]
		if _, exists := indexMap[index]; exists {
			log.Fatalf("share with index '%v' was already entered.  Each share must have a unique index", index)
		}
	}

	checkummedSecret := recoverChecksummedSecret(xValues, yValues)
	secret := recoverUnchecksummedSecret(checkummedSecret)
	return secret
}

// RecoverShare returns full information from a share
func RecoverShare(share []string, secretSizeBytes int) (index, threshold uint, shamirBytes []byte) {
	mnemonicIndexes := wordencoding.RecoverIndexes(share)
	checksummedBuffer := wordencoding.RecoverChecksummedBuffer(mnemonicIndexes, secretSizeBytes)
	unchecksummedBuffer := checksummedBuffer.GetUnchecksummedBuffer(2)
	unchecksummedBits := unchecksummedBuffer.GetBits()
	indexBits := unchecksummedBits[0:5]
	thresholdBits := unchecksummedBits[5:10]
	shamirBits := unchecksummedBits[10:]
	indexRaw, _ := strconv.ParseUint(indexBits, 2, 64)
	thresholdRaw, _ := strconv.ParseUint(thresholdBits, 2, 64)
	index = uint(indexRaw)
	if index < 1 || index > 31 {
		log.Fatal("invalid index, must be 1 <= index <= 31")
	}
	threshold = uint(thresholdRaw)
	if threshold < 1 || threshold > 31 {
		log.Fatal("invalid threshold, must be 1 <= threshold <= 31")
	}
	shamirBytes = bits.GetBytes(shamirBits)
	return index, threshold, shamirBytes
}

func createShamirData(n, k uint, secret []byte) ([]uint, [][]byte) {
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

func recoverChecksummedSecret(xValues []uint, yValues [][]byte) []byte {
	numberOfShares := len(yValues)
	if numberOfShares < 2 {
		log.Fatal("need at least two shares to recover a secret")
	}
	checksummedSecretLength := len(yValues[0])
	checksummedSecret := make([]byte, checksummedSecretLength)

	// for each byte in the secret
	for i := 0; i < checksummedSecretLength; i++ {
		subXValues := make([]uint, numberOfShares)
		subYValues := make([]uint, numberOfShares)
		// for each k shares
		for j := 0; j < numberOfShares; j++ {
			subXValues[j] = xValues[j]
			subYValues[j] = uint(yValues[j][i])
		}
		interpolation := maths.LagrangeInterpolate(0, subXValues, subYValues)
		checksummedSecret[i] = byte(interpolation)
	}
	return checksummedSecret
}

func createUnchecksummedShare(shamirPart []byte, index, threshold uint) *bits.SmartBuffer {
	indexBits := bits.GetBits(byte(index), 5)
	thresholdBits := bits.GetBits(byte(threshold), 5)
	shamirBits := bits.GetBitsArray(shamirPart, 8)

	allBits := indexBits + thresholdBits + shamirBits
	allBitsPadded := bits.PadBits(allBits)
	bytes := bits.GetBytes(allBitsPadded)
	smartBuffer := bits.SmartBufferFromBytes(bytes, len(allBits))
	return smartBuffer
}

func createChecksummedSecret(unchecksummedSecret []byte) []byte {
	checksum := cryptos.GetSha256(unchecksummedSecret)[:2]
	checksummedSecret := append(unchecksummedSecret, checksum...)
	return checksummedSecret
}

func recoverUnchecksummedSecret(checksummedSecret []byte) []byte {
	expectedChecksum := checksummedSecret[len(checksummedSecret)-2:]
	data := checksummedSecret[:len(checksummedSecret)-2]
	actualChecksum := cryptos.GetSha256(data)
	if !bytes.Equal(expectedChecksum, actualChecksum[:2]) {
		log.Fatal("actual master secret checksum did not match expected checksum")
	}
	return data
}
