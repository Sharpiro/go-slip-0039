package secretsharing

import (
	"bytes"
	"go-slip-0039/cryptos"
	"go-slip-0039/maths"
	"go-slip-0039/maths/bits"
	"go-slip-0039/wordencoding"
	"log"
	"strconv"
	"strings"
)

// CreateMnemonicWordsList creates shares based off a given secret
func CreateMnemonicWordsList(shareCount, threshold byte, secret []byte, passphrase string) [][]string {
	// sharedIdentifierBits := bits.GetBitsArray(cryptos.GetRandomBytes(4), 8)[:30]
	// encryptionKey := getEncryptionKey(passphrase, sharedIdentifierBits, threshold)
	shares := splitSecret(threshold, shareCount, secret)
	_ = shares

	var mnemonicWordsList [][]string
	for i := 0; i < len(shares); i++ {
		// unchecksummedShare := createUnchecksummedShare(threshold, shares[i])
		// checksummedShare := unchecksummedShare.GetChecksummedBuffer()
		indexList := createIndexList(99, 0, 0, 1, 1, 0, 3, shares[i].Y)
		_ = indexList
		// mnemonicWords := wordencoding.CreateMnemonicWords(indexList)
		// mnemonicWordsList = append(mnemonicWordsList, mnemonicWords)
	}
	return mnemonicWordsList
}

func createIndexList(id, iterationExponent, groupIndex, groupThreshold, groupCount, memberIndex, memberThreshold int,
	share []byte) []uint {
	temp1 := bits.GetBits(byte(id), 15)
	temp2 := bits.GetBits(byte(iterationExponent), 5)
	temp3 := bits.GetBits(byte(groupIndex), 4)
	temp4 := bits.GetBits(byte(groupThreshold-1), 4)
	temp5 := bits.GetBits(byte(groupCount-1), 4)
	temp6 := bits.GetBits(byte(memberIndex), 4)
	temp7 := bits.GetBits(byte(memberThreshold-1), 4)
	temp8 := bits.GetBitsArray(share, 8)
	paddedShareSize := len(temp8) + (10 - (len(temp8) % 10))
	paddingSize := paddedShareSize - len(temp8)
	padding := strings.Repeat("0", paddingSize)
	temp9 := padding + temp8
	checksum := strings.Repeat("0", 30)
	concat := temp1 + temp2 + temp3 + temp4 + temp5 + temp6 + temp7 + temp9 + checksum
	_ = concat
	return []uint{}
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

	mnemonicIndexes := wordencoding.RecoverIndexList(share)
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

	checksum := cryptos.GetRandomBytes(len(secret))
	basePoints := make([]maths.Point, threshold)
	basePoints[0] = maths.Point{X: 254, Y: checksum}
	basePoints[1] = maths.Point{X: 255, Y: secret}
	shares := make([]maths.Point, shareCount)

	for i := range shares {
		shares[i] = maths.Point{X: byte(i), Y: maths.LagrangeInterpolate(byte(i), basePoints)}
	}

	return shares
}

func uRange(maxExclusive uint) []uint {
	data := make([]uint, maxExclusive)
	for i := range data {
		data[i] = uint(i)
	}
	return data
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
