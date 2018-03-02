package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"go-slip-0039/cryptos"
	"go-slip-0039/secretsharing"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
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
					Value: "fffaac234dac4dacfffaac234dac4dacfffaac234dac4dacfffaac234dac4dac",
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
				cli.BoolFlag{
					Name:  "protected",
					Usage: "hide console input",
				},
				cli.StringFlag{
					Name:  "passphrase",
					Value: "",
					Usage: "an optional passphrase for generating a seed",
				},
				cli.IntFlag{
					Name:  "size",
					Value: 256,
					Usage: "the size in bits of the master secret",
				},
			},
			Action: recover,
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
	protected := context.Bool("protected")
	secretSizeBits := context.Int("size")
	readInput(protected, secretSizeBits)
	return
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

func readInput(protected bool, secretSizeBits int) {
	var data string

	fmt.Println("please enter your first share")
	if protected {
		dataBytes, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal(err)
		}
		data = string(dataBytes)
	} else {
		reader := bufio.NewReader(os.Stdin)
		var err error
		dataBytes, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		data = string(dataBytes)
	}
	split := strings.Split(data, " ")
	index, threshold, _ := secretsharing.AnalyzeShare(split)
	fmt.Println(split)
	fmt.Println("index: ", index)
	fmt.Println("threshold: ", threshold)
	fmt.Println("length: ", secretSizeBits)
}
