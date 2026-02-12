package parse

import (
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"
)

var tagEndCharaacters = []byte{'>', '/'}

func ParseTag(r io.Reader) (*Tag, error) {
	var tagType strings.Builder
	var attributes []Attribute

	bytes := make([]byte, 1)

	for {
		_, err := r.Read(bytes)

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, err
		}

		byte := bytes[0]

		if slices.Contains(tagEndCharaacters, byte) {
			attributes = make([]Attribute, 0)
			break
		}

		if byte == ' ' {
			attributes, err = parseAttributes(r)

			if err != nil {
				return nil, err
			}
		}

		tagType.WriteByte(byte)
	}

	if tagType.Len() == 0 {
		return nil, fmt.Errorf("Unable to find tag.")
	}

	tag := &Tag{
		tagType.String(),
		attributes,
	}

	return tag, nil
}

func parseAttributes(r io.Reader) ([]Attribute, error) {
	return make([]Attribute, 0), nil
}
