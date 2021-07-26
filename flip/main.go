package main

import (
	cr "crypto/rand"
	"flag"
	"fmt"
	"math/big"
)

func main() {
	flips := flag.Int("f", 1000, "number of flips to perform")
	flag.Parse()
	fmt.Println(findSide(*flips))
}

func findSide(flips int) string {
	half := flips / 2
	rand := genCryptoRng()
	z := 0
	for i := 0; i < flips; i++ {
		if n, _ := rand(1); n == 0 {
			z++
		}
	}
	if z > half {
		return "heads"
	}
	return "tails"
}

func genCryptoRng() func(int) (int64, error) {
	return func(n int) (int64, error) {
		bn, err := cr.Int(cr.Reader, big.NewInt(int64(n)))
		if err != nil {
			return -1, err
		}
		return bn.Int64(), nil
	}
}
