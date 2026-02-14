package parse

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"iter"
	"slices"
	"strings"

	"github.com/spoik/html-parser/html"
)

func isAttrNameEndChar(byte byte) bool {
	return byte == ' '
}

func parseAttributes(r *bufio.Reader) ([]*html.Attribute, error) {
	// Make an educated guess that most html elements will have ~5 attributes
	attributes := make([]*html.Attribute, 5)

	for {
		bytes, err := r.Peek(1)

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, err
		}

		byte := bytes[0]

		if isTagEndChar(byte) {
			break
		}

		if isAttrNameEndChar(byte) {
			r.Discard(1)
			continue
		}

		attribute, err := parseAttribute(r)

		if err != nil {
			return nil, err
		}

		attributes = append(attributes, attribute)
	}

	attributes = slices.DeleteFunc(attributes, func(a *html.Attribute) bool {
		return a == nil
	})

	return attributes, nil
}

func parseAttribute(r *bufio.Reader) (*html.Attribute, error) {
	nextByte, checkErr := attributeNameSeq(r)

	var attributeNameBuilder strings.Builder

	for byte := range nextByte {
		attributeNameBuilder.WriteByte(byte)
	}

	if attributeNameBuilder.Len() == 0 {
		return nil, fmt.Errorf("Unable to find attribute")
	}

	err := checkErr()

	if err != nil {
		return nil, err
	}

	attribute := &html.Attribute{Name: attributeNameBuilder.String()}

	bytes, err := r.Peek(1)

	if err != nil {
		if errors.Is(err, io.EOF) {
			return attribute, nil
		}

		return nil, err
	}

	byte := bytes[0]

	var value string

	if byte == '=' {
		_, err := r.Discard(1)

		if err != nil {
			return nil, err
		}

		value, err = parseValue(r)

		if err != nil {
			return nil, err
		}
	}

	attribute.Value = value

	return attribute, nil
}

func attributeNameSeq(r *bufio.Reader) (iter.Seq[byte], func() error) {
	var err error

	seq := func(yield func(byte) bool) {
		for {
			bytes, e := r.Peek(1)

			if e != nil {
				if errors.Is(e, io.EOF) {
					break
				}

				e = err
				break
			}

			byte := bytes[0]

			if byte == '=' {
				break
			}

			if isTagEndChar(byte) || isAttrNameEndChar(byte) {
				break
			}

			_, e = r.Discard(1)

			if e != nil {
				err = e
				break
			}

			if !yield(byte) {
				break
			}
		}
	}

	return seq, func() error { return err }
}
