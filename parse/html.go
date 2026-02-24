package parse

import (
	"bufio"
	"errors"
	"fmt"
	"io"

	"github.com/spoik/html-parser/html"
	"github.com/spoik/html-parser/stringreader"
)

// Returns a slice of html.Tag instances that represent the html provide in s.
func ParseHtml(s *string) (*html.Tags, error) {
	r := bufio.NewReaderSize(
		stringreader.New(*s),
		2,
	)

	tags, err := parseTags(r)

	if err != nil {
		return nil, fmt.Errorf("Error parsing HTML: %w", err)
	}

	if len(tags) == 0 {
		return nil, fmt.Errorf("No HTML tags found in \"%s\"", *s)
	}

	return &html.Tags{Tags: tags}, nil
}

func parseTags(r *bufio.Reader) ([]*html.Tag, error) {
	// TODO: Change this slice to have a fixed size to avoid constantly resizing the slice.
	// Not sure what the best size would be though: make([]*html.Tag, 50)
	var tags []*html.Tag

	for {
		byte, err := r.ReadByte()

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return nil, err
		}

		if byte != '<' {
			continue
		}

		tag, err := ParseTag(r)

		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	return tags, nil
}
