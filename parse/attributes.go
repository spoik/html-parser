package parse

import (
	"bufio"
	"errors"
	"fmt"
	"io"
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
	attributeName, err := parseAttributeName(r)

	if err != nil {
		return nil, err
	}

	if len(attributeName) == 0 {
		return nil, fmt.Errorf("Unable to find attribute")
	}

	attribute := &html.Attribute{Name: attributeName}

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

		value, err = parseAttributeValue(r)

		if err != nil {
			return nil, err
		}
	}

	attribute.Value = value

	return attribute, nil
}

func parseAttributeName(r *bufio.Reader) (string, error) {
	var attributeNameBuilder strings.Builder

	for {
		bytes, err := r.Peek(1)

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return "", err
		}

		byte := bytes[0]

		if byte == '=' {
			break
		}

		if isTagEndChar(byte) || isAttrNameEndChar(byte) {
			break
		}

		_, err = r.Discard(1)

		if err != nil {
			return "", err
		}

		err = attributeNameBuilder.WriteByte(byte)

		if err != nil {
			return "", err
		}
	}

	return attributeNameBuilder.String(), nil
}
