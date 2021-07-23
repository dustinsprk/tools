package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/dustinsprk/tools/jsonfix/pkg/format"
)

func main() {
	var level int
	var text bool
	var indentText string
	var flatten bool
	flag.Usage = usage
	levelUsage := "number of spaces per indentation level"
	flag.IntVar(&level, "level", 4, levelUsage)
	flag.IntVar(&level, "l", 4, "(abbrv) "+levelUsage)
	flag.BoolVar(&text, "text", false, "read from command line arg")
	flag.BoolVar(&text, "t", false, "(abbrv) read from command line arg")
	flag.StringVar(&indentText, "indent-text", " ", "one or more `CHARACTERS` to be used for indentation. this will be repeated `level` times")
	flag.StringVar(&indentText, "i", " ", "(abbrv) `CHARACTERS` to be used for indentation")
	flag.BoolVar(&flatten, "flatten", false, "flatten/compress input")
	flag.BoolVar(&flatten, "f", false, "(abbrv) flatten/compress input")
	flag.Parse()
	args := flag.Args()
	if indentText == "" {
		indentText = " "
	}
	if flatten && level != 4 {
		errorf("level and flatten are mutually exclusive\n")
		flag.Usage()
		return
	}
	var raw []byte
	var err error
	switch text {
	case false:
		raw, err = toBytes(os.Stdin)
		if err != nil {
			errorf("error reading from stdin: %v", err)
			return
		}
	case true:
		raw = []byte(args[0])
	}
	var b []byte
	switch flatten {
	case true:
		b, err = format.FlattenJSON(raw)
		if err != nil {
			errorf("invalid json: %v\n", err)
			errorf(string(raw))
			return
		}
	case false:
		b, err = format.IndentJSON(raw, indentText, level)
		if err != nil {
			errorf("invalid json: %v\n", err)
			errorf(string(raw))
			return
		}
	}
	fmt.Println(string(b))
}

func toBytes(r io.Reader) ([]byte, error) {
	buf := &bytes.Buffer{}
	_, err := io.Copy(buf, r)
	if err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

func errorf(s string, args ...interface{}) (int, error) {
	return fmt.Fprintf(os.Stderr, s, args...)
}

func usage() {
	errorf("usage: %s [OPTIONS] JSONDATA\n", os.Args[0])
	flag.PrintDefaults()
}
