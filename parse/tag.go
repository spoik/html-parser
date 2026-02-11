package parse

import (
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"
)

func TagAtPosition(r io.Reader) (*Tag, error) {
	tagType, err := tagType(r)

	if err != nil {
		return nil, err
	}

	return &Tag{Type: tagType}, nil
}

var tagTypeEndCharaacters = []byte{' ', '>', '/'}

func tagType(r io.Reader) (string, error) {
	var tagType strings.Builder

	bytes := make([]byte, 1)

	for {
		_, err := r.Read(bytes)

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return "", err
		}

		byte := bytes[0]

		if slices.Contains(tagTypeEndCharaacters, byte) {
			break
		}

		tagType.WriteByte(byte)
	}

	if tagType.Len() == 0 {
		return "", fmt.Errorf("Unable to find tag type.")
	}

	return tagType.String(), nil
}
