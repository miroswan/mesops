package v1

import (
	"bufio"
	"bytes"
	"strconv"
	"testing"
)

type byteReader struct {
	buf *bytes.Buffer
}

// Read returns the buffered data a byte at a time.
func (m byteReader) Read(p []byte) (n int, err error) {
	b, err := m.buf.ReadByte()
	if err != nil {
		return 0, err
	}

	p[0] = b
	return 1, nil
}

// AddMessage buffers the given string in RecordIO format.
func (b byteReader) AddMessage(msg string) {
	b.buf.WriteString(strconv.Itoa(len(msg)))
	b.buf.WriteString("\n")
	b.buf.WriteString(msg)
}

func TestRecordioReader(t *testing.T) {
	cases := []string{
		"short",
		"medium",
		"looooooooooonnngggg",
	}

	r := byteReader{buf: &bytes.Buffer{}}

	for _, c := range cases {
		r.AddMessage(c)
	}

	b := bufio.NewReader(r)
	for _, c := range cases {
		msg, err := readRecordioMessage(b)
		if err != nil {
			t.Errorf("got error %s, wanted nil", err)
		}
		if string(msg) != c {
			t.Errorf("got message '%s', wanted '%s'", msg, c)
		}
	}
}
