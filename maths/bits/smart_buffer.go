package bits

import (
	"go-slip-0039/cryptos"
	"log"
)

// SmartBuffer holds a buffer as well as the number of actual bits
type SmartBuffer struct {
	Buffer []byte
	Size   int
}

func (smartBuffer *SmartBuffer) GetChecksum() *SmartBuffer {
	checksumBytes := cryptos.GetSha256(smartBuffer.Buffer)[:4]
	checksumSmartBuffer := SmartBufferFromBytes(checksumBytes, 30)
	return checksumSmartBuffer
}

func (thisBuffer *SmartBuffer) Append(otherBuffer *SmartBuffer) *SmartBuffer {
	thisBufferBits := thisBuffer.GetBits()
	otherBufferBits := otherBuffer.GetBits()
	combinedBits := thisBufferBits + otherBufferBits
	newSmartBuffer := SmartBufferFromBits(combinedBits)
	return newSmartBuffer
}

func (smartBuffer *SmartBuffer) PopBits(size int) *SmartBuffer {
	smartBufferBits := GetBitsArray(smartBuffer.Buffer, 8)[:smartBuffer.Size]
	poppedBits := smartBufferBits[len(smartBufferBits)-size:]
	remainingBits := smartBufferBits[:len(smartBufferBits)-size]
	if len(poppedBits) != size {
		log.Fatal("'PopBits returned incorrectly sized slice")
	}
	if len(remainingBits) != smartBuffer.Size-size {
		log.Fatal("'PopBits returned incorrectly sized slice")
	}

	remainingSmartBuffer := SmartBufferFromBits(remainingBits)
	smartBuffer.Buffer = remainingSmartBuffer.Buffer
	smartBuffer.Size = remainingSmartBuffer.Size
	poppedSmartBuffer := SmartBufferFromBits(poppedBits)
	return poppedSmartBuffer
}

func (smartBuffer *SmartBuffer) GetBits() string {
	return GetBitsArray(smartBuffer.Buffer, 8)[:smartBuffer.Size]
}

func (smartBuffer *SmartBuffer) GetChecksummedBuffer() *SmartBuffer {
	checksum := smartBuffer.GetChecksum()
	newSmartBuffer := smartBuffer.Append(checksum)
	return newSmartBuffer
}

func (smartBuffer *SmartBuffer) GetUnchecksummedBuffer() *SmartBuffer {
	const checksumSize int = 30
	cloneBuffer := smartBuffer.Clone()
	expectedChecksum := cloneBuffer.PopBits(checksumSize)
	actualChecksum := cloneBuffer.GetChecksum()
	// actualChecksum := cryptos.GetSha256(cloneBuffer.Buffer)[:4]
	if !expectedChecksum.Equals(actualChecksum) {
		log.Fatal("invalid share checksum")
	}
	return cloneBuffer
}

func (smartBuffer *SmartBuffer) Clone() *SmartBuffer {
	clonedBuffer := make([]byte, len(smartBuffer.Buffer))
	copy(clonedBuffer, smartBuffer.Buffer)
	return SmartBufferFromBytes(clonedBuffer, smartBuffer.Size)
}

func (thisBuffer *SmartBuffer) Equals(otherBuffer *SmartBuffer) bool {
	thisBufferBits := thisBuffer.GetBits()
	otherBufferBits := otherBuffer.GetBits()
	areEqual := thisBufferBits == otherBufferBits
	return areEqual
}

func SmartBufferFromBytes(buffer []byte, sizeBits int) *SmartBuffer {
	return &SmartBuffer{buffer, sizeBits}
}

func SmartBufferFromBits(bits string) *SmartBuffer {
	paddedBits := PadBits(bits)
	buffer := GetBytes(paddedBits)
	return &SmartBuffer{buffer, len(bits)}
}
