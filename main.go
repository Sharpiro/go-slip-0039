package main

import (
	"bytes"
	cryptoRandom "crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	mathRandom "math/rand"
	"time"

	"./secretsharing"
)

func main() {
	// xValues, yValues := secretsharing.CreateShares(6, 3, []byte{1, 2, 3})
	// fmt.Println(xValues)
	// fmt.Println(yValues)
	start := time.Now()
	// tests
	for i := 0; i < 10000; i++ {
		randomLength := mathRandom.Intn(64) + 1
		fmt.Println(randomLength)
		secretBytes := make([]byte, randomLength, randomLength)
		cryptoRandom.Read(secretBytes)
		xValues, yValues := secretsharing.CreateShares(6, 3, secretBytes)
		recoveredBytes := secretsharing.RecoverSecret(xValues, yValues)
		if !bytes.Equal(secretBytes, recoveredBytes) {
			secretString := hex.EncodeToString(secretBytes)
			recoveredString := hex.EncodeToString(recoveredBytes)
			fmt.Println(secretString)
			fmt.Println(recoveredString)
			log.Fatal("fatal")
		}
		// fmt.Println(recoveredString)
		//  assert secretBytes == secretsharing.RecoverSecret(shares[:3])
		//  assert secretBytes == secretsharing.RecoverSecret(shares[-3:])
		//  assert secretBytes == secretsharing.RecoverSecret([shares[1], shares[3], shares[4]])
	}
	elapsed := time.Since(start)
	fmt.Println("finished successfully")
	fmt.Println(elapsed)
}
