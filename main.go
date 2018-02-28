package main

import (
	"encoding/hex"
	"fmt"
	"go-slip-0039/cryptos"
	"go-slip-0039/secretsharing"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-slip-0039"
	app.Usage = "create and recover slip-0039 secrets"
	app.Commands = []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "create secret shares",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "secret",
					Value: "fffaac234dac",
					Usage: "the master secret",
				},
			},
			Action: create,
		},
		{
			Name:    "recover",
			Aliases: []string{"r"},
			Usage:   "recover secret shares",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "shares",
					Value: "a list of shares",
					Usage: "the master secret",
				},
				cli.StringFlag{
					Name:  "passphrase",
					Value: "",
					Usage: "an optional passphrase for generating a seed",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("recovering is harder....")
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func create(context *cli.Context) {
	// secretBytes := []byte(*secretPtr)
	secret := context.String("secret")
	secretBytes, err := hex.DecodeString(secret)
	if err != nil {
		log.Fatal("an error occurred decoding the hex string to bytes")
	}
	shares := secretsharing.CreateWordShares(3, 2, secretBytes)

	fmt.Printf("secret: %v\n", secret)
	fmt.Println(shares)
}

func recover(context *cli.Context) {
	var shares [][]string
	var secretBytes []byte
	passPhrase := context.String("passphrase")

	bitLength := (len(secretBytes) + 4) << 3
	recoveredSecretBytes := secretsharing.RecoverFromWordShares(shares[1:], bitLength)
	recoveredSecret := hex.EncodeToString(recoveredSecretBytes)
	generatedSeed := cryptos.CreatePbkdf2Seed(recoveredSecretBytes, passPhrase)
	generatedSeedHex := hex.EncodeToString(generatedSeed)

	fmt.Printf("passphrase: %v\n", passPhrase)
	fmt.Printf("recovered secret: %v\n", recoveredSecret)
	fmt.Printf("seed: %v\n", generatedSeedHex)
}
