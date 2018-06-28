package bits

// SmartBuffer holds a buffer as well as the number of actual bits
type SmartBuffer struct {
	buffer []byte
	size   int
}
