package bits

import (
	"testing"
)

func TestSmartBufferCopy(tester *testing.T) {
	smartBuffer := SmartBuffer{buffer: []byte{1, 2, 3, 4, 5}, size: 40}
	pointerToSmartBuffer := &smartBuffer
	copyOfSmartBuffer := smartBuffer
	pointerToSmartBuffer.size = 12

	if smartBuffer.size != 12 || pointerToSmartBuffer.size != 12 {
		tester.Error()
	}

	if copyOfSmartBuffer.size != 40 {
		tester.Error()
	}
}
