package main

import (
	"./secretsharing"
	"encoding/hex"
	"math"
	// maths "./maths/gflogmaths"
	"fmt"
	// maths "./maths/gfmaths"
	cryptoRandom "crypto/rand"
	"log"
	mathRandom "math/rand"
)

func main() {
	// xValues, yValues := secretsharing.CreateShares(6, 3, []byte{1, 2, 3})
	// fmt.Println(xValues)
	// fmt.Println(yValues)

	// tests
	for i := 0; i < 1000; i++ {
		randomLength := mathRandom.Intn(int(math.Pow(2, 5))) + 1
		secretBytes := make([]byte, randomLength, randomLength)
		cryptoRandom.Read(secretBytes)
		secretString := hex.EncodeToString(secretBytes)
		xValues, yValues := secretsharing.CreateShares(6, 3, secretBytes)
		recovered := secretsharing.RecoverSecret(xValues, yValues)
		recoveredString := hex.EncodeToString(recovered)
		if recoveredString != secretString {
			fmt.Println(secretString)
			fmt.Println(recoveredString)
			fmt.Println(secretBytes)
			fmt.Println(recovered)
			log.Fatal("fatal")
		}
		fmt.Println(recoveredString)
		//  assert secretBytes == secretsharing.RecoverSecret(shares[:3])
		//  assert secretBytes == secretsharing.RecoverSecret(shares[-3:])
		//  assert secretBytes == secretsharing.RecoverSecret([shares[1], shares[3], shares[4]])
	}
	fmt.Println("finished successfully")
}
