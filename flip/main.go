package main

import (
	cr "crypto/rand"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"os"
)

func main() {
	flips := flag.Int("f", 1000, "number of flips to perform")
	flag.Parse()
	s, err := findSide(*flips)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unexpected error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(s)
}

func findSide(flips int) (string, error) {
	rand := genCryptoRng(2)
	c := 0
	for i := 0; i < flips; i++ {
		n, err := rand()
		if err != nil {
			return "", errors.New(fmt.Sprintf("error generating random number %v", err))
		}
		if n == 0 {
			c++
		}
	}
	half := float32(float32(flips) / 2.0)
	cf := float32(c)
	if cf == half {
		return "draw", nil
	}
	if cf > half {
		return "heads", nil
	}
	return "tails", nil
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
