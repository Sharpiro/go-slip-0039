package secretsharing

import (
	"testing"
)

func TestRecoverMasterSecret(tester *testing.T) {
	// secretBytes := []byte{9, 8, 7, 6}
	// in decimal = 999475216
	getEncryptionKey("", "111011100100101100100100010000", 3)
}
