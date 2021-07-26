package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dustinsprk/tools/time"
)

func main() {
	var newLine bool
	var negative bool
	flag.BoolVar(&newLine, "suppress-new-line", false, "suppress a final newline")
	flag.BoolVar(&newLine, "s", false, "(abbrv) suppress a final newline")
	flag.BoolVar(&negative, "negative", false, "use negative to go into the past")
	flag.BoolVar(&negative, "n", false, "(abbrv) use negative to go into the past")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		usage()
		return
	}
	n, err := strconv.Atoi(args[0])
	if err != nil {
		errorf("cannot parse %s as an integer: %v", args[1], err)
		return
	}
	ms := int64(n * 1000)
	if negative {
		ms = ms * -1
	}
	now := time.NowUnixMillis()
	suffix := "\n"
	if newLine {
		suffix = ""
	}
	fmt.Printf("%d%s", now+ms, suffix)
}

func usage() {
	b := strings.Builder{}
	b.WriteString("get the time n seconds from now in milliseconds (UTC)\n")
	b.WriteString(fmt.Sprintf("  usage: %s [OPTIONS] seconds\n\n", os.Args[0]))
	b.WriteString("positional:\n")
	b.WriteString("  seconds: number of seconds in the future\n")
	b.WriteString("optional:\n")
	errorf(b.String())
	flag.PrintDefaults()
}

func errorf(s string, args ...interface{}) (int, error) {
	return fmt.Fprintf(os.Stderr, s, args...)
}
