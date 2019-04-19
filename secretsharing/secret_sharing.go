package secretsharing

import (
	"bytes"
	"go-slip-0039/cryptos"
	"go-slip-0039/maths"
	"go-slip-0039/maths/bits"
	"go-slip-0039/wordencoding"
	"log"
	"math"
	"strconv"
	"strings"
)

// CreateMnemonicWordsList creates shares based off a given secret
func CreateMnemonicWordsList(shareCount, threshold byte, secret []byte, passphrase string) [][]string {
	shares := splitSecret(threshold, shareCount, secret)
	_ = shares
	id := cryptos.GetRandomBelow((int(math.Pow(2, 15))) - 1)
	var mnemonicWordsList [][]string
	for _, share := range shares {
		indexList := createIndexList(id, 0, 0, 1, 1, share.X, threshold, share.Y)
		mnemonicWords := wordencoding.CreateMnemonicWords(indexList)
		mnemonicWordsList = append(mnemonicWordsList, mnemonicWords)
	}
	return mnemonicWordsList
}

func createIndexList(id int, iterationExponent, groupIndex, groupThreshold, groupCount, memberIndex,
	memberThreshold byte, shareValue []byte) []int {
	idBits := bits.GetBits(byte(id), 15)
	iterationExponentBits := bits.GetBits(byte(iterationExponent), 5)
	groupIndexBits := bits.GetBits(byte(groupIndex), 4)
	groupThresholdBits := bits.GetBits(byte(groupThreshold-1), 4)
	groupCountBits := bits.GetBits(byte(groupCount-1), 4)
	memberIndexBits := bits.GetBits(byte(memberIndex), 4)
	memberThresholdBits := bits.GetBits(byte(memberThreshold-1), 4)
	shareValueBits := bits.GetBitsArray(shareValue, 8)
	paddedShareSize := len(shareValueBits) + (10 - (len(shareValueBits) % 10))
	paddingSize := paddedShareSize - len(shareValueBits)
	padding := strings.Repeat("0", paddingSize)
	paddedShareValueBits := padding + shareValueBits
	shareBits := idBits + iterationExponentBits + groupIndexBits + groupThresholdBits + groupCountBits +
		memberIndexBits + memberThresholdBits + paddedShareValueBits

	// var groups []uint
	indexes := make([]int, 0, 17)
	for i := 0; i < len(shareBits); i += 10 {
		indexBits := shareBits[i : i+10]
		index, err := strconv.ParseInt(indexBits, 2, 15)
		if err != nil {
			log.Fatal("failed converting bits to bytes")
		}
		indexes = append(indexes, int(index))
	}

	checksum := cryptos.RS1024CreateChecksum(indexes)
	indexes = append(indexes, checksum...)
	return indexes
}

// RecoverSecretFromMnemonicShares recovers a secret based off of K supplied word lists
func RecoverSecretFromMnemonicShares(mnemonicWordsList [][]string) []byte {
	var xValues []uint
	var yValues [][]byte

	_, index, secretThreshold, shamirBuffer := RecoverShare(mnemonicWordsList[0])
	xValues = append(xValues, index)
	yValues = append(yValues, shamirBuffer)

	for i := 1; i < len(mnemonicWordsList); i++ {
		mnemonicWords := mnemonicWordsList[i]
		_, index, shareThreshold, shamirBuffer := RecoverShare(mnemonicWords)
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
		indexMap[index] = true
	}

	secret := recoverSecret(xValues, yValues)
	return secret
}

// RecoverShare returns full information from a share
func RecoverShare(share []string) (nonce, index, threshold uint, shamirBytes []byte) {
	if len(share) < 7 {
		log.Fatalf("invalid share, minimum share size is 7 words, actual was %v", len(share))
	}

	// mnemonicIndexes := wordencoding.RecoverIndexList(share)
	mnemonicIndexes := make([]uint, 0)
	checksummedBuffer := wordencoding.RecoverChecksummedBuffer(mnemonicIndexes)
	unchecksummedBuffer := checksummedBuffer.GetUnchecksummedBuffer()
	unchecksummedBits := unchecksummedBuffer.GetBits()
	nonceBits := unchecksummedBits[0:20]
	indexBits := unchecksummedBits[20:25]
	thresholdBits := unchecksummedBits[25:30]
	shamirBits := unchecksummedBits[30:]
	nonceRaw, _ := strconv.ParseUint(nonceBits, 2, 64)
	indexRaw, _ := strconv.ParseUint(indexBits, 2, 64)
	thresholdRaw, _ := strconv.ParseUint(thresholdBits, 2, 64)
	if indexRaw < 1 || indexRaw > 31 {
		log.Fatal("invalid index, must be 1 <= index <= 31")
	}
	if thresholdRaw < 1 || thresholdRaw > 31 {
		log.Fatal("invalid threshold, must be 1 <= threshold <= 31")
	}
	strippedShamirBits := bits.StripPaddingFromNearestTen(shamirBits)
	shamirBytes = bits.GetBytes(strippedShamirBits)
	return uint(nonceRaw), uint(indexRaw), uint(thresholdRaw), shamirBytes
}

func splitSecret(threshold, shareCount byte, secret []byte) []maths.Point {
	if shareCount < threshold {
		log.Fatalf("n must be greater than k, secret would be unrecoverable")
	}

	digest := createDigest(secret)
	basePoints := make([]maths.Point, threshold)
	basePoints[0] = maths.Point{X: 254, Y: digest}
	basePoints[1] = maths.Point{X: 255, Y: secret}
	shares := make([]maths.Point, shareCount)

	for i := range shares {
		shares[i] = maths.Point{X: byte(i), Y: maths.LagrangeInterpolate(byte(i), basePoints)}
	}

	return shares
}

func createDigest(secret []byte) []byte {
	randomKey := cryptos.GetRandomBytes(len(secret) - 4)
	hmac := cryptos.GetHmacSha256(randomKey, secret)[:4]
	digest := append(hmac, randomKey...)
	return digest
}

func recoverSecret(xValues []uint, yValues [][]byte) []byte {
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
		interpolation := maths.LagrangeInterpolateOld(0, subXValues, subYValues)
		checksummedSecret[i] = byte(interpolation)
	}
	return checksummedSecret
}

// func createUnchecksummedShare(threshold byte, share maths.Point) *bits.SmartBuffer {
// 	identifierBits := bits.GetBitsArray(cryptos.GetRandomBytes(3), 8)[:20]
// 	indexBits := bits.GetBits(byte(share.X), 5)
// 	thresholdBits := bits.GetBits(byte(threshold), 5)
// 	shareValueBits := bits.GetBitsArray(shamirPart, 8)
// 	paddedShareValueBits := bits.PadShareToNearestTen(shareValueBits)

// 	allBits := identifierBits + indexBits + thresholdBits + paddedShareValueBits
// 	allBitsPadded := bits.PadBits(allBits)
// 	bytes := bits.GetBytes(allBitsPadded)
// 	smartBuffer := bits.SmartBufferFromBytes(bytes, len(allBits))
// 	return smartBuffer
// }

func recoverUnchecksummedSecret(checksummedSecret []byte) []byte {
	expectedChecksum := checksummedSecret[len(checksummedSecret)-2:]
	data := checksummedSecret[:len(checksummedSecret)-2]
	actualChecksum := cryptos.GetSha256(data)
	if !bytes.Equal(expectedChecksum, actualChecksum[:2]) {
		log.Fatal("actual master secret checksum did not match expected checksum")
	}
	return data
}
