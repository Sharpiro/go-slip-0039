package bits

import (
	"bytes"
	"testing"
)

func TestSmartBufferCopy(tester *testing.T) {
	smartBuffer := *NewSmartBuffer([]byte{1, 2, 3, 4, 5}, 40)
	pointerToSmartBuffer := &smartBuffer
	copyOfSmartBuffer := smartBuffer
	pointerToSmartBuffer.Size = 12

	if smartBuffer.Size != 12 || pointerToSmartBuffer.Size != 12 {
		tester.Error()
	}

	if copyOfSmartBuffer.Size != 40 {
		tester.Error()
	}
}

func TestChecksum(tester *testing.T) {
	shareSize := 58
	share := NewSmartBuffer([]byte{0, 66, 194, 129, 65, 24, 118, 192}, shareSize)
	expectedChecksumShare := []byte{0, 66, 194, 129, 65, 24, 118, 234, 170, 64}

	checksum := share.GetChecksum()
	actualChecksummedShare := share.Append(checksum)
	actualChecksummedShare2 := share.GetChecksummedBuffer()

	if share.Size != shareSize {
		tester.Error()
	}
	if actualChecksummedShare.Size != shareSize+16 {
		tester.Error()
	}
	if !bytes.Equal(expectedChecksumShare, actualChecksummedShare.Buffer) {
		tester.Error()
	}

	if !bytes.Equal(actualChecksummedShare2.Buffer, actualChecksummedShare.Buffer) {
		tester.Error()
	}

	if actualChecksummedShare2.Size != actualChecksummedShare.Size {
		tester.Error()
	}
}
