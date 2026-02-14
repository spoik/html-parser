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

func getTagType(r *bufio.Reader) (string, error) {
	var typeBuilder strings.Builder

	nextByte, checkErr := tagTypeIterator(r)

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

func tagTypeIterator(r *bufio.Reader) (iter.Seq[byte], func() error) {
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
