package secretsharing

import (
	"bytes"
	"go-slip-0039/maths/bits"
	"go-slip-0039/wordencoding"
	"reflect"
	"strconv"
	"testing"
)

func TestCreateChecksummedShare(tester *testing.T) {
	shamirPart := []byte{255}
	unchecksummedShare := createUnchecksummedShare(shamirPart, 1, 2)
	checksummedShare := unchecksummedShare.GetChecksummedBuffer()
	rebuiltUnchecksummedShare := checksummedShare.GetUnchecksummedBuffer()

	if !bytes.Equal(unchecksummedShare.Buffer, rebuiltUnchecksummedShare.Buffer) {
		tester.Error()
	}
	if checksummedShare.Size != 70 || len(checksummedShare.Buffer) != 9 {
		tester.Error()
	}
}

func TestCreateIndexList(tester *testing.T) {
	checksummedShare := []byte{1, 2, 0, 139, 252, 124, 213, 134, 144}
	smartBuffer := bits.SmartBufferFromBytes(checksummedShare, 70)
	expectedIndexList := []uint{0, 0, 34, 1020, 499, 344, 420}
	actualIndexList := wordencoding.CreateIndexList(smartBuffer)

	if !reflect.DeepEqual(expectedIndexList[2:], actualIndexList[2:]) {
		tester.Error()
	}
}

func TestCreateMnemonicWords(tester *testing.T) {
	indexList := []uint{4, 32, 34, 1020, 499, 344, 420}
	expectedMnemonicWords := []string{"", "", "angry", "write", "laundry", "fatal", "half"}
	actualMnemonicWords := wordencoding.CreateMnemonicWords(indexList)

	if !reflect.DeepEqual(expectedMnemonicWords[2:], actualMnemonicWords[2:]) {
		tester.Error()
	}
}

func TestRecoverIndexes(tester *testing.T) {
	mnemonicWords := []string{"actress", "anchor", "angry", "write", "laundry", "fatal", "half"}
	expectedIndexList := []uint{4, 32, 34, 1020, 499, 344, 420}
	actualIndexList := wordencoding.RecoverIndexList(mnemonicWords)

	if !reflect.DeepEqual(expectedIndexList, actualIndexList) {
		tester.Error()
	}
}

func TestRecoverChecksummedBuffer(tester *testing.T) {
	indexList := []uint{4, 32, 34, 1020, 499, 344, 420}
	checksummedBuffer := wordencoding.RecoverChecksummedBuffer(indexList)
	unchecksummedBuffer := checksummedBuffer.GetUnchecksummedBuffer()

	_ = unchecksummedBuffer
}

func TestRecoverSecret(tester *testing.T) {
	unchecksummedBuffer := bits.SmartBufferFromBytes([]byte{1, 2, 0, 139, 252}, 40)

	unchecksummedBits := unchecksummedBuffer.GetBits()
	indexBits := unchecksummedBits[20:25]
	thresholdBits := unchecksummedBits[25:30]
	shamirBits := unchecksummedBits[30:]
	indexRaw, _ := strconv.ParseUint(indexBits, 2, 64)
	thresholdRaw, _ := strconv.ParseUint(thresholdBits, 2, 64)
	strippedShamirBits := bits.StripPaddingFromNearestTen(shamirBits)
	shamirBytes := bits.GetBytes(strippedShamirBits)

	if indexRaw != 1 {
		tester.Error()
	}
	if thresholdRaw != 2 {
		tester.Error()
	}
	if len(shamirBytes) != 1 || shamirBytes[0] != 255 {
		tester.Error()
	}
}

func TestStripPadding(tester *testing.T) {
	bitsPadded := "1111111100"
	bits := bits.StripPaddingFromNearestTen(bitsPadded)
	if bits != "11111111" {
		tester.Error()
	}
}
