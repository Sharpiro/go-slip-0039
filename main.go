package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"go-slip-0039/secretsharing"
	"log"
	"os"
	"strconv"
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
					Name:  "secrethex, s",
					Usage: "the master secret in valid hex format",
				},
				cli.StringFlag{
					Name:  "n",
					Usage: "total number of shares",
				},
				cli.StringFlag{
					Name:  "k",
					Usage: "total number of shares needed to re-create secret",
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
					Name:  "protected, p",
					Usage: "optionally hide console input for word shares",
				},
				cli.StringFlag{
					Name:  "passphrase, pp",
					Value: "",
					Usage: "an optional passphrase for generating a seed",
				},
				// cli.IntFlag{
				// 	Name:  "size, s",
				// 	Usage: "the size in bits of the master secret",
				// },
			},
			Action: recover,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func create(context *cli.Context) {
	n, err := strconv.Atoi(context.String("n"))
	if err != nil || n < 2 {
		log.Fatal("n must be a number greater than or equal to 2")
	}
	k, err := strconv.Atoi(context.String("k"))
	if err != nil || k < 2 || k > n {
		log.Fatal("k must be greater than or equal to 2 and less than or equal to n")
	}
	secretHex := context.String("secrethex")
	if secretHex == "" {
		log.Fatalf("must provide a secret in valid hex form")
	}
	secretBytes, err := hex.DecodeString(secretHex)
	if err != nil {
		log.Fatal("an error occurred decoding the hex string to bytes")
	}
	shares := secretsharing.CreateMnemonicWordsList(uint(n), uint(k), secretBytes, "")
	replacer := strings.NewReplacer("[", "", "]", "")
	for _, share := range shares {
		formattedShare := replacer.Replace(fmt.Sprint(share))
		fmt.Println(formattedShare)
	}
}

func recover(context *cli.Context) {
	protected := context.Bool("protected")
	xValues, yValues := readShares(protected)
	// passPhrase := context.String("passphrase")

	recoveredSecretBytes := secretsharing.RecoverSecretFromShamirData(xValues, yValues)
	recoveredSecret := hex.EncodeToString(recoveredSecretBytes)
	// generatedSeed := cryptos.CreatePbkdf2Seed(recoveredSecretBytes, passPhrase)
	// generatedSeedHex := hex.EncodeToString(generatedSeed)

	fmt.Printf("recovered secret: %v\n", recoveredSecret)
	// fmt.Printf("generated seed: %v\n", generatedSeedHex)
}

func readShares(protected bool) (xValues []uint, yValues [][]byte) {
	var wordLists [][]string
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("please enter first share:\n")
	firstShare := readShare(reader, protected)

	wordLists = append(wordLists, firstShare)
	secretNonce, index, secretThreshold, shamirBuffer := secretsharing.RecoverShare(firstShare)
	var shareThreshold uint
	var shareNonce uint
	xValues = append(xValues, index)
	yValues = append(yValues, shamirBuffer)
	indexMap := map[uint]bool{index: true}
	fmt.Printf("index: '%v' threshold: '%v'\n\n", index, secretThreshold)

	for i := uint(1); i < secretThreshold; i++ {
		fmt.Printf("please enter share %v/%v:\n", i+1, secretThreshold)
		share := readShare(reader, protected)
		shareNonce, index, shareThreshold, shamirBuffer = secretsharing.RecoverShare(share)
		fmt.Printf("index: '%v' threshold: '%v'\n\n", index, shareThreshold)
		if shareNonce != secretNonce {
			log.Fatalf("the share's nonce '%v', did not match the first share's nonce '%v'", shareNonce, secretNonce)
		}
		if shareThreshold != secretThreshold {
			log.Fatalf("the share's threshold '%v', did not match the first share's threshold '%v'", shareThreshold, secretThreshold)
		}
		if _, exists := indexMap[index]; exists {
			log.Fatalf("share with index '%v' was already entered.  Each share must have a unique index", index)
		}
		indexMap[index] = true
		xValues = append(xValues, index)
		yValues = append(yValues, shamirBuffer)
		wordLists = append(wordLists, share)
	}
	fmt.Println()
	return xValues, yValues
}

func readShare(reader *bufio.Reader, protected bool) []string {
	var data string
	if protected {
		dataBytes, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal(err)
		}
		data = string(dataBytes)
	} else {
		var err error
		dataBytes, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		data = string(dataBytes)
	}
	split := strings.Split(data, " ")
	return split
}
