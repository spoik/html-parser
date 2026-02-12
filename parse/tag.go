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

var tagEndCharacters = []byte{'>', '/'}

func ParseTag(r *bufio.Reader) (*html.Tag, error) {
	var tagType strings.Builder
	var attributes []*html.Attribute

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

		if slices.Contains(tagEndCharacters, byte) {
			break
		}

		if byte == ' ' {
			attributes, err = parseAttributes(r)

			if err != nil {
				return nil, err
			}

			continue
		}

		tagType.WriteByte(byte)
	}

	if tagType.Len() == 0 {
		return nil, fmt.Errorf("Unable to find tag.")
	}

	tag := &html.Tag{
		Type:       tagType.String(),
		Attributes: attributes,
	}

	return tag, nil
}
