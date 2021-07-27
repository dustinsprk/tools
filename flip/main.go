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
	rand := genCryptoRng(2)
	c := 0
	for i := 0; i < flips; i++ {
		if n, _ := rand(); n == 0 {
			c++
		}
	}
	half := float32(float32(flips) / 2.0)
	cf := float32(c)
	if cf == half {
		return "draw"
	}
	if cf > half {
		return "heads"
	}
	return "tails"
}

func genCryptoRng(n int) func() (int64, error) {
	return func() (int64, error) {
		bn, err := cr.Int(cr.Reader, big.NewInt(int64(n)))
		if err != nil {
			return -1, err
		}
		return bn.Int64(), nil
	}
}
