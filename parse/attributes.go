package parse

import (
	"bufio"
	"errors"
	"io"
	"slices"
	"strings"

	"github.com/spoik/html-parser/html"
)

var attrNameEndChar = []byte{' ', '='}

func isAttrNameEndChar(byte byte) bool {
	return slices.Contains(attrNameEndChar, byte)
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
	var attributeName strings.Builder

	for {
		bytes, err := r.Peek(1)

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, err
		}

		byte := bytes[0]

		if isTagEndChar(byte) || isAttrNameEndChar(byte) {
			break
		}

		_, err = r.Discard(1)

		if err != nil {
			return nil, err
		}

		attributeName.WriteByte(byte)
	}

	return &html.Attribute{Name: attributeName.String()}, nil
}
