# Go-SLIP-0039

This is an implementation of the [SLIP-0039](https://github.com/satoshilabs/slips/blob/master/slip-0039.md) for Shamir's Secret Sharing.  This repo is a work in progress, and neither the SLIP nor this codebase is certified for use at this time.

## Installation

```
go get
go build
```

## Usage
```bash
go-slip-0039 -h
```
### Creating Shares
```bash
# create 3 total secret shares, where 2 shares are needed to recreate the master secret "ff" in hexadecimal format
go-slip-0039 create -n 3 -k 2 --secrethex ff
```
### Recovering Shares
```bash
# recover a master secret of size 8 in bits.  Optionally add 'protected' to hide share words in console.  Optionally add 'passphrase' to add a pass phrase to the auto generated seed output
go-slip-0039 recover --size 8 --protected --passphrase phrase
```

## High-level Process

### Create Mnemonics

* Input secret, passphrase, threshold, and number of shares
* ...

### Recover Mnemonics

* ...

## Issues

* confusion with string concatentaiton operator [#497](https://github.com/satoshilabs/slips/issues/497)
* curve.png f(0) points to Master secret when in the explanation it points to pre-master secret [#501](https://github.com/satoshilabs/slips/issues/501)
* lacks a simple high level overview of the exact steps to create and recover shares
* Notation table lacks detail
    * ```s``` is unclear
    * need a better explanation of pre-master secret and master secret
        * perhaps more clear naming
* unclear of different processes when generating master secret vs providing own master secret