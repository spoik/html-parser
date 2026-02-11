package stringreader

import (
	"io"
)

type StringReader struct {
	str   string
	pos int
	strLen   int
}

func New(string string) *StringReader {
	return &StringReader{
		str:   string,
		strLen:   len(string),
		pos: -1,
	}
}

func (sr *StringReader) Position() int {
	return sr.pos
}

func (sr *StringReader) Read(p []byte) (n int, err error) {
	var readLen int

	for readLen = range len(p) {
		err = sr.advancePosition()

		if err != nil {
			return readLen, err
		}

		p[readLen] = sr.str[sr.pos]
	}

	return readLen + 1, nil
}

func (sr *StringReader) advancePosition() error {
	if sr.pos == sr.strLen-1 {
		return io.EOF
	}

	sr.pos++
	return nil
}
