package parse

import (
	"bufio"
	"fmt"

	"github.com/spoik/html-parser/html"
	"github.com/spoik/html-parser/stringreader"
)

// Returns a slice of html.Tag instances that represent the html provide in s.
func ParseHtml(s *string) ([]*html.Tag, error) {
	r := bufio.NewReaderSize(
		stringreader.New(*s),
		2,
	)


	tags, err := ParseTags(r)

	if err != nil {
		return nil, fmt.Errorf("Error parsing HTML: %w", err)
	}

	if len(tags) == 0 {
		return nil, fmt.Errorf("No HTML tags found in \"%s\"", *s)
	}

	return tags, nil
}
