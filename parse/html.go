package parse

import (
	"bufio"
	"errors"
	"fmt"
	"io"

	"github.com/spoik/html-parser/html"
	"github.com/spoik/html-parser/stringreader"
)

// Returns a html.Tags instances that represent the html provide in s.
func ParseHtml(s *string) (html.Tags, error) {
	index := &html.TagIndex{}
	tags, err := parseHtml(s, index)

	if err != nil {
		return html.Tags{}, err
	}

	return html.NewTags(html.NewTagsOpts{
		Tags: tags,
		TagIndex: index,
	}), nil
}

func parseHtml(s *string, index *html.TagIndex) ([]html.Tag, error) {
	r := bufio.NewReaderSize(
		stringreader.New(*s),
		2,
	)

	tags, err := parseTags(r, index)

	if err != nil {
		return nil, fmt.Errorf("Error parsing HTML: %w", err)
	}

	if len(tags) == 0 {
		return nil, fmt.Errorf("No HTML tags found in \"%s\"", *s)
	}

	return tags, nil
}

func parseTags(r *bufio.Reader, index *html.TagIndex) ([]html.Tag, error) {
	// TODO: Change this slice to have a fixed size to avoid constantly resizing the slice.
	// Not sure what the best size would be though: make([]*html.Tag, 50)
	var tags []html.Tag

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

		tag, err := ParseTag(r, index)

		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	return tags, nil
}
