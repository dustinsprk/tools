package main

import (
	cr "crypto/rand"
	"fmt"
	"math/big"
	"os"
)

func main() {
	choices := os.Args[1:]
	if len(choices) < 1 {
		fmt.Fprintln(os.Stderr, "expected > 0 choices")
		os.Exit(1)
	}
	rng := genCryptoRng(int64(len(choices)))
	fmt.Println(choices[rng()])
}

func genCryptoRng(nOpts int64) func() int64 {
	max := big.NewInt(nOpts)
	return func() int64 {
		bn, err := cr.Int(cr.Reader, max)
		if err != nil {
			panic(err)
		}
		return bn.Int64()
	}
}
