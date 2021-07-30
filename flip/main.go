package main

import (
	cr "crypto/rand"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"os"
)

type Winner string

const (
	Draw  Winner = "draw"
	Heads Winner = "heads"
	Tails Winner = "tails"
)

type FlipResult struct {
	Total int
	Heads int
	Tails int
}

func (r FlipResult) Winner() Winner {
	if r.Heads > r.Tails {
		return Heads
	}
	if r.Heads < r.Tails {
		return Tails
	}
	return Draw
}

func main() {
	flips := flag.Int("f", 1000, "number of flips to perform")
	verbose := flag.Bool("v", false, "verbose output")
	flag.Parse()
	r, err := findSide(*flips)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unexpected error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(r.Winner())
	if *verbose {
		fmt.Printf("heads: %d\ntails: %d\n", r.Heads, r.Tails)
	}
}

func findSide(flips int) (FlipResult, error) {
	rand := genCryptoRng(2)
	result := FlipResult{Total: flips, Heads: 0, Tails: 0}
	for i := 0; i < flips; i++ {
		n, err := rand()
		if err != nil {
			return FlipResult{}, errors.New(fmt.Sprintf("error generating random number %v", err))
		}
		if n == 0 {
			result.Heads++
		}
	}
	result.Tails = flips - result.Heads
	return result, nil
}

func genCryptoRng(n int64) func() (int64, error) {
	max := big.NewInt(n)
	return func() (int64, error) {
		bn, err := cr.Int(cr.Reader, max)
		if err != nil {
			return -1, err
		}
		return bn.Int64(), nil
	}
}
