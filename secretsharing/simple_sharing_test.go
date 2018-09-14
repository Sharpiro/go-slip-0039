package secretsharing

import (
	"go-slip-0039/maths/bits"
	"go-slip-0039/wordencoding"
	"testing"
)

func TestCreateChecksummedShare(tester *testing.T) {
	shamirPart := []byte{255}
	// expectedShare := []byte{8, 130, 194, 129, 65, 24, 118, 192}
	unchecksummedShare := createUnchecksummedShare(shamirPart, 1, 2)
	checksummedShare := unchecksummedShare.GetChecksummedBuffer()
	temp := checksummedShare.GetUnchecksummedBuffer()

	_ = temp

	// if !bytes.Equal(expectedShare, actualShare.Buffer) {
	// 	tester.Error()
	// }
	// if actualShare.Size != 58 {
	// 	tester.Error()
	// }
}

func TestCreateIndexList(tester *testing.T) {
	checksummedShare := []byte{1, 2, 0, 139, 252, 124, 213, 134, 144}
	smartBuffer := bits.SmartBufferFromBytes(checksummedShare, 70)
	// expectedShare := []byte{8, 130, 194, 129, 65, 24, 118, 192}
	indexList := wordencoding.CreateIndexList(smartBuffer)

	_ = indexList
}

func TestCreateMnemonicWords(tester *testing.T) {
	indexList := []uint{1233}
	smartBuffer := wordencoding.CreateMnemonicWords(indexList)
	// expectedShare := []byte{8, 130, 194, 129, 65, 24, 118, 192}

	_ = smartBuffer
}
