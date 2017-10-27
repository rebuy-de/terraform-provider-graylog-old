package graylog

import (
	"bytes"
	"io"
)

func ReaderToString(r io.Reader) string {
	if r == nil {
		return ""
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	return buf.String()
}
