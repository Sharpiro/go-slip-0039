package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"

	"go-slip-0039/secretsharing"
)

func main() {
	wordPtr := flag.String("secret", "default", "your secret")

	flag.Parse()

	// secretBytes := []byte(*wordPtr)
	secretBytes, err := hex.DecodeString(*wordPtr)
	if err != nil {
		log.Fatal("an error occurred decoding the hex string to byte")
	}

	bitLength := (len(secretBytes) + 4) << 3
	shares := secretsharing.CreateWordShares(3, 2, secretBytes)
	recoveredSecretBytes := secretsharing.RecoverFromWordShares(shares[1:], bitLength)
	recoveredSecret := hex.EncodeToString(recoveredSecretBytes)

	fmt.Println(*wordPtr)
	fmt.Println(shares)
	fmt.Println(recoveredSecret)
}
