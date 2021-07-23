package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
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
	t := parseMs(n)
	s := t.Format("2 Jan 2006 15:04:05 MST")
	fmt.Println(s)

}

func parseMs(ms int64) time.Time {
	s := ms / 1000
	rem := ms - (s * 1000)
	return time.Unix(s, rem)
}

func errorf(s string, args ...interface{}) (int, error) {
	return fmt.Fprintf(os.Stderr, s, args...)
}

func nowUnixMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
