# Go-SLIP-0039

This is an implementation of the [SLIP-0039](https://github.com/satoshilabs/slips/blob/master/slip-0039.md) for Shamir's Secret Sharing.  This repo is a work in progress, and neither the SLIP nor this codebase is certified for use at this time.

## Installation

```
Go Get
Go build
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