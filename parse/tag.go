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

var tagEndBytes = []byte{'>', '/'}

func isTagEndChar(b byte) bool {
	return slices.Contains(tagEndBytes, b)
}

var attributeDeliminer byte = ' '

func isAttributeDeliminer(b byte) bool {
	return b == attributeDeliminer
}

func ParseTag(r *bufio.Reader) (*html.Tag, error) {
	tagType, err := parseTagType(r)

	if err != nil {
		return nil, err
	}

	attributes, err := getAttributes(r)

	if err != nil {
		return nil, err
	}

	tag := &html.Tag{
		Type:       tagType,
		Attributes: attributes,
	}

	return tag, nil
}

func parseTagType(r *bufio.Reader) (string, error) {
	var typeBuilder strings.Builder

	for {
		bytes, err := r.Peek(1)

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return "", err
		}

		byte := bytes[0]

		if isTagEndChar(byte) || isAttributeDeliminer(byte) {
			break
		}

		r.Discard(1)

		err = typeBuilder.WriteByte(byte)

		if err != nil {
			return "", err
		}
	}

	if typeBuilder.Len() == 0 {
		return "", fmt.Errorf("Unable to find tag.")
	}

	return typeBuilder.String(), nil
}

func getAttributes(r *bufio.Reader) ([]*html.Attribute, error) {
	var attributes []*html.Attribute

	bytes, err := r.Peek(1)

	if err != nil {
		if errors.Is(err, io.EOF) {
			return attributes, nil
		}

		return nil, err
	}

	byte := bytes[0]

	if !isAttributeDeliminer(byte) {
		return attributes, nil
	}

	return parseAttributes(r)
}
