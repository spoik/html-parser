package parse

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"iter"
	"strings"
)

func getTagType(r *bufio.Reader) (string, error) {
	var typeBuilder strings.Builder

	nextByte, checkErr := tagTypeSeq(r)

	for byte := range nextByte {
		typeBuilder.WriteByte(byte)
	}

	if err := checkErr(); err != nil {
		return "", err
	}

	if typeBuilder.Len() == 0 {
		return "", fmt.Errorf("Unable to find tag.")
	}

	return typeBuilder.String(), nil
}

func tagTypeSeq(r *bufio.Reader) (iter.Seq[byte], func() error) {
	var err error

	seq := func(yield func(byte) bool) {
		for {
			bytes, e := r.Peek(1)

			if e != nil {
				if errors.Is(e, io.EOF) {
					break
				}

				err = e
				break
			}

			byte := bytes[0]

			if isTagEndChar(byte) || isAttributeDeliminer(byte) {
				break
			}

			r.Discard(1)

			if !yield(byte) {
				break
			}
		}
	}

	return seq, func() error { return err }
}
