package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	b, err := toBytes(os.Stdin)
	if err != nil {
		panic(err)
	}
	s := string(b)
	s = strings.TrimRight(s, "\n")
	fmt.Print(s)
}

func toBytes(r io.Reader) ([]byte, error) {
	buf := &bytes.Buffer{}
	_, err := io.Copy(buf, r)
	if err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}
