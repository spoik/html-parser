package parse

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spoik/html-parser/stringreader"
)

func TagAtPosition(sr *stringreader.Reader) (*Tag, error) {
	tagType, err := tagType(sr)

	if err != nil {
		return nil, err
	}

	return &Tag{Type: tagType}, nil
}

var tagTypeEndCharaacters = []byte{' ', '>', '/'}

func tagType(sr *stringreader.Reader) (string, error) {
	var tagType strings.Builder

	for !sr.AtEnd() {
		char, err := sr.ReadNext()

		if err != nil {
			return "", err
		}

		if slices.Contains(tagTypeEndCharaacters, char) {
			break
		}

		tagType.WriteByte(char)
	}

	if tagType.Len() == 0 {
		return "", fmt.Errorf(
			"Unable to find tag type in \"%s\" starting at position %d.",
			sr.String(),
			sr.Position(),
		)
	}

	return tagType.String(), nil
}
