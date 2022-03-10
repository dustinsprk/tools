package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"time"
)

var suppressNL bool
var hours bool
var weeks bool

func init() {
	flag.BoolVar(&suppressNL, "suppress-new-line", false, "suppress a final newline")
	flag.BoolVar(&suppressNL, "s", false, "(abbrv) suppress a final newline")
	flag.BoolVar(&hours, "hours", false, "show hours instead of days")
	flag.BoolVar(&hours, "H", false, "show hours instead of days")
	flag.BoolVar(&weeks, "weeks", false, "(abbrv) show seeks instead of days")
	flag.BoolVar(&weeks, "w", false, "(abbrv) show weeks instead of days")
	flag.Parse()
}

func main() {
	if hours && weeks {
		fmt.Fprintln(os.Stderr, "hours and weeks cannot both be set")
		os.Exit(1)
	}

	now := time.Now().In(time.Local)
	then := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.Local)
	divisor := float64(24)
	if hours {
		divisor = float64(1)
	}
	if weeks {
		divisor = float64(168)
	}
	r := now.Sub(then).Hours() / divisor
	rem := math.Round(r)
	suff := "\n"
	if suppressNL {
		suff = ""
	}
	fmt.Printf("%d%s", int(rem), suff)
}
