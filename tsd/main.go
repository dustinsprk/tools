package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dustinsprk/tools/time"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		errorf("usage: %s timestamp_ms\n", args[0])
		return
	}
	n, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		errorf("cannot parse %s: %v", args[1], err)
		return
	}
	t := time.ParseMs(n)
	s := t.Format("2 Jan 2006 15:04:05 MST")
	fmt.Println(s)

}

func errorf(s string, args ...interface{}) (int, error) {
	return fmt.Fprintf(os.Stderr, s, args...)
}
