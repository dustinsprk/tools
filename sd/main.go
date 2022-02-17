package main

import (
	"flag"
	"fmt"
	"math"
	"time"
)

var suppressNL bool
var hours bool

func init() {
	flag.BoolVar(&suppressNL, "suppress-new-line", false, "suppress a final newline")
	flag.BoolVar(&suppressNL, "s", false, "(abbrv) suppress a final newline")
	flag.BoolVar(&hours, "hours", false, "show hours instead of days")
	flag.BoolVar(&hours, "h", false, "(abbrv) show hours instead of days")

	flag.Parse()
}

func main() {
	now := time.Now().In(time.Local)
	then := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local)
	divisor := float64(24)
	if hours {
		divisor = float64(1)
	}
	r := now.Sub(then).Hours() / divisor
	rem := math.Ceil(r)
	suff := "\n"
	if suppressNL {
		suff = ""
	}
	fmt.Printf("%d%s", int(rem), suff)
}
