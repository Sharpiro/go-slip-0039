package bits

import (
	"go-slip-0039/cryptos"
)

// SmartBuffer holds a buffer as well as the number of actual bits
type SmartBuffer struct {
	Buffer []byte
	Size   int
}

func (smartBuffer *SmartBuffer) GetChecksum() []byte {
	checksum := cryptos.GetSha256(smartBuffer.Buffer)[:2]
	return checksum
}

func (smartBuffer *SmartBuffer) Append(buffer []byte) *SmartBuffer {
	smartBufferBits := GetBitsArray(smartBuffer.Buffer, 8)[:smartBuffer.Size]
	bufferBits := GetBitsArray(buffer, 8)
	combinedBits := smartBufferBits + bufferBits
	paddedBits := PadBits(combinedBits)
	newSmartBuffer := NewSmartBuffer(GetBytes(paddedBits), len(combinedBits))
	return newSmartBuffer
}

func (smartBuffer *SmartBuffer) GetChecksummedBuffer() *SmartBuffer {
	checksum := smartBuffer.GetChecksum()
	newSmartBuffer := smartBuffer.Append(checksum)
	return newSmartBuffer
}

func NewSmartBuffer(buffer []byte, size int) *SmartBuffer {
	return &SmartBuffer{buffer, size}
}
