package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "expected integer status code\n")
		return
	}
	n, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid status code %s\n", args[1])
		return
	}
	t := http.StatusText(n)
	if len(t) == 0 {
		fmt.Fprintf(os.Stderr, "invalid status code %d\n", n)
		return
	}

	fmt.Println(t)
}
