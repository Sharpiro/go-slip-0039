package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"go-slip-0039/cryptos"
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
				cli.IntFlag{
					Name:  "size, s",
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
	shares := secretsharing.CreateWordShares(uint(n), uint(k), secretBytes)

	fmt.Printf("secret: %v\n", secretHex)
	fmt.Println(shares)
}

func recover(context *cli.Context) {
	secretSizeBits := context.Int("size")
	if secretSizeBits < 1 {
		log.Fatal("must provide size in bits of master secert to be recovered")
	}
	protected := context.Bool("protected")
	shares := readShares(protected)
	passPhrase := context.String("passphrase")

	// totalBitLength := (len(secretBytes) + 4) << 3
	totalBitLength := secretSizeBits + 32
	recoveredSecretBytes := secretsharing.RecoverFromWordShares(shares, totalBitLength)
	recoveredSecret := hex.EncodeToString(recoveredSecretBytes)
	generatedSeed := cryptos.CreatePbkdf2Seed(recoveredSecretBytes, passPhrase)
	generatedSeedHex := hex.EncodeToString(generatedSeed)

	fmt.Printf("recovered secret: %v\n", recoveredSecret)
	fmt.Printf("generated seed: %v\n", generatedSeedHex)
}

func readShares(protected bool) [][]string {
	var wordLists [][]string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("please enter first share:")
	firstShare := readShare(reader, protected)

	wordLists = append(wordLists, firstShare)
	index, threshold, _ := secretsharing.AnalyzeShare(firstShare)
	fmt.Printf("index: %v\tthreshold: %v\n", index, threshold)

	for i := 1; i < threshold; i++ {
		fmt.Printf("please enter share %v/%v:\n", i+1, threshold)
		share := readShare(reader, protected)
		wordLists = append(wordLists, share)
	}
	return wordLists
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
