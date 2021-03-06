package main

import (
	"flag"
	"fmt"
	"time"
)

var suppressNL bool

func init() {
	flag.BoolVar(&suppressNL, "suppress-new-line", false, "suppress a final newline")
	flag.BoolVar(&suppressNL, "s", false, "(abbrv) suppress a final newline")
	flag.Parse()
}
func main() {
	yd := time.Now().UTC().YearDay()
	suff := "\n"
	if suppressNL {
		suff = ""
	}
	fmt.Printf("%d%s", yd, suff)
}
