package parse

import (
	"bufio"
	"errors"
	"io"
	"slices"

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
	tagType, err := getTagType(r)

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
