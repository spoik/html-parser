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

func getAttributes(r *bufio.Reader) (html.Attributes, error) {
	var attributes html.Attributes

	bytes, err := r.Peek(1)

	if errors.Is(err, io.EOF) {
		return attributes, nil
	}

	if err != nil {
		return html.Attributes{}, err
	}

	byte := bytes[0]

	if !isAttributeDeliminer(byte) {
		return attributes, nil
	}

	return parseAttributes(r)
}

func parseAttributes(r *bufio.Reader) (html.Attributes, error) {
	// Make an educated guess that most html elements will have ~5 attributes
	attributes := make([]html.Attribute, 5)

	for {
		bytes, err := r.Peek(1)

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return html.Attributes{}, err
		}

		byte := bytes[0]

		if isTagEndChar(byte) {
			break
		}

		if isAttrNameEndChar(byte) {
			_, err := r.Discard(1)

			if err != nil {
				return html.Attributes{}, err
			}

			continue
		}

		attribute, err := parseAttribute(r)

		if err != nil {
			return html.Attributes{}, err
		}

		attributes = append(attributes, attribute)
	}

	attributes = slices.DeleteFunc(attributes, func(a html.Attribute) bool {
		return a == html.Attribute{}
	})

	if len(attributes) == 0 {
		return html.Attributes{}, nil
	}

	return html.NewAttributes(attributes), nil
}

func parseAttribute(r *bufio.Reader) (html.Attribute, error) {
	attributeName, err := parseAttributeName(r)

	if err != nil {
		return html.Attribute{}, err
	}

	if attributeName == "" {
		return html.Attribute{}, err
	}

	value, err := parseValue(r)

	if err != nil {
		return html.Attribute{}, err
	}

	attribute := html.Attribute{
		Name:  attributeName,
		Value: value,
	}

	return attribute, nil
}

func parseAttributeName(r *bufio.Reader) (string, error) {
	var attributeName strings.Builder

	for {
		bytes, err := r.Peek(1)

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
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

		err = attributeName.WriteByte(byte)

		if err != nil {
			return "", err
		}
	}

	if attributeName.Len() == 0 {
		return "", fmt.Errorf("Unable to find attribute")
	}

	return attributeName.String(), nil
}
