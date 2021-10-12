package main

import (
	"flag"
	"fmt"
	"math"
	"time"
)

var suppressNL bool

func init() {
	flag.BoolVar(&suppressNL, "suppress-new-line", false, "suppress a final newline")
	flag.BoolVar(&suppressNL, "s", false, "(abbrv) suppress a final newline")
	flag.Parse()
}

func main() {
	now := time.Now().In(time.Local)
	then := time.Date(2021, time.October, 12, 8, 0, 0, 0, time.Local)
	r := now.Sub(then).Hours() / float64(24)
	rem := math.Ceil(r)
	suff := "\n"
	if suppressNL {
		suff = ""
	}
	fmt.Printf("%d%s", int(rem), suff)
}
