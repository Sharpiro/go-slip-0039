package secretsharing

import (
	"bytes"
	"go-slip-0039/cryptos"
	"go-slip-0039/maths"
	"go-slip-0039/maths/bits"
	"log"
	"strconv"
)

// CreateWordShares creates shares based off a given secret
func CreateWordShares(n, k uint, secret []byte) [][]string {
	checksummedSecret := getChecksummedSecret(secret)
	xValues, yValues := createShares(n, k, checksummedSecret)
	shares := createRawShares(xValues, yValues, k)
	checksummedShares := createChecksummedShares(shares)
	indexesList := getIndexesList(checksummedShares, len(secret))
	mnemoniclists := getMnemonicList(indexesList)
	return mnemoniclists
}

// RecoverFromWordShares recovers a secret based off of K supplied word lists
func RecoverFromWordShares(mnemonicLists [][]string, secretSizeBytes int) []byte {
	indexesList := getMnemonicIndexesList(mnemonicLists)
	checksummedBuffers := getChecksummedBuffers(indexesList, secretSizeBytes)
	unchecksummedBuffers := getUnchecksummedBuffers(checksummedBuffers)
	indexes, shamirParts := recoverFromShare(unchecksummedBuffers)
	checkummedSecret := recoverChecksummedSecret(indexes, shamirParts)
	secret := getUnchecksummedSecret(checkummedSecret)
	return secret
}

// AnalyzeShare returns useful data about a given share
func AnalyzeShare(share []string, secretSizeBytes int) (index, threshold int) {
	mnemonicIndexes := getMnemonicIndexes(share)
	checksummedBuffer := getChecksummedBuffer(mnemonicIndexes, secretSizeBytes)
	unchecksummedBuffer := getUnchecksummedBuffer(checksummedBuffer)
	bits := unchecksummedBuffer.GetBits()
	indexBits := bits[0:5]
	thresholdBits := bits[5:10]
	indexRaw, _ := strconv.ParseInt(indexBits, 2, 64)
	thresholdRaw, _ := strconv.ParseInt(thresholdBits, 2, 64)
	index = int(indexRaw) + 1
	threshold = int(thresholdRaw) + 1

	length := len(share[0]) >> 1 // todo: what is this again??
	_ = length                   // lol still don't remember what this is
	return index, threshold
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

func recoverFromShare(shares []*bits.SmartBuffer) ([]uint, [][]byte) {
	const encodingOffset byte = 1
	xValues := make([]uint, len(shares))
	yValues := make([][]byte, len(shares))
	for i, share := range shares {
		shareBits := share.GetBits()
		indexBits := shareBits[:5]
		thresholdBits := shareBits[5:10]
		shamirPartBits := shareBits[10:]

		rawIndex, _ := strconv.ParseInt(indexBits, 2, 64)
		rawThreshold, _ := strconv.ParseInt(thresholdBits, 2, 64)

		if len(shamirPartBits)%8 != 0 {
			log.Fatal("shamir part bits must be a multiple of 8")
		}
		shamirPart := bits.GetBytes(shamirPartBits)

		threshold := rawThreshold + 1
		_ = threshold

		xValues[i] = uint(rawIndex + 1)
		yValues[i] = shamirPart
	}
	return xValues, yValues
}

func createRawShares(xValues []uint, yValues [][]byte, k uint) []*bits.SmartBuffer {
	shares := make([]*bits.SmartBuffer, 0)
	for i, j := range yValues {
		share := makeShare(j, uint(i+1), k)
		shares = append(shares, share)
	}
	return shares
}

func createChecksummedShares(smartBuffers []*bits.SmartBuffer) []*bits.SmartBuffer {
	checksummedShares := make([]*bits.SmartBuffer, 0)
	for _, j := range smartBuffers {
		checksummedShare := j.GetChecksummedBuffer()
		checksummedShares = append(checksummedShares, checksummedShare)
	}
	return checksummedShares
}

func makeShare(shamirPart []byte, index, threshold uint) *bits.SmartBuffer {
	indexBits := bits.GetBits(byte(index-1), 5)
	thresholdBits := bits.GetBits(byte(threshold-1), 5)
	shamirBits := bits.GetBitsArray(shamirPart, 8)

	allBits := indexBits + thresholdBits + shamirBits
	allBitsPadded := bits.PadBits(allBits)
	bytes := bits.GetBytes(allBitsPadded)
	smartBuffer := bits.SmartBufferFromBytes(bytes, len(allBits))
	return smartBuffer
}

func getChecksummedSecret(secret []byte) []byte {
	checksum := cryptos.GetSha256(secret)[:2]
	checksummedSecret := append(secret, checksum...)
	return checksummedSecret
}

func getUnchecksummedSecret(checksummedSecret []byte) []byte {
	expectedChecksum := checksummedSecret[len(checksummedSecret)-2:]
	data := checksummedSecret[:len(checksummedSecret)-2]
	actualChecksum := cryptos.GetSha256(data)
	if !bytes.Equal(expectedChecksum, actualChecksum[:2]) {
		log.Fatal("actual master secret checksum did not match expected checksum")
	}
	return data
}
