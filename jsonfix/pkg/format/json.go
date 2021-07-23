package format

import (
	"bytes"
	"encoding/json"
	"strings"
)

func IndentJSON(b []byte, indentText string, level int) ([]byte, error) {
	buf := bytes.Buffer{}
	indent := strings.Repeat(indentText, level)
	err := json.Indent(&buf, b, "", indent)
	return buf.Bytes(), err
}

func FlattenJSON(b []byte) ([]byte, error) {
	buf := bytes.Buffer{}
	err := json.Compact(&buf, b)
	return buf.Bytes(), err
}
