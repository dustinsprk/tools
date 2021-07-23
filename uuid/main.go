package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func main() {
	var newLine bool
	var n int
	flag.BoolVar(&newLine, "suppress-new-line", false, "suppress a final newline")
	flag.BoolVar(&newLine, "s", false, "suppress a final newline (abbrv)")
	flag.IntVar(&n, "number", 1, "`NUMBER` of UUIDs to get")
	flag.IntVar(&n, "n", 1, "`NUMBER` of UUIDs to get")
	flag.Parse()
	if n > 1 && newLine {
		s := "supress-new-line and number > 1 are mutually exclusive"
		fmt.Fprintln(os.Stderr, s)
		return
	}
	for i := 0; i < n; i++ {

		v, err := uuid.NewRandom()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error getting an UUID: %v", err)
			return
		}
		suff := ""
		if !newLine {
			suff = "\n"
		}
		fmt.Printf("%s%s", v, suff)
	}
}
