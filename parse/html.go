package parse

import (
	"fmt"

	"github.com/spoik/html-parser/stringreader"
)

type Tag struct {
	Type string
}

func ParseHtml(html *string) (*[]Tag, error) {
	sr := stringreader.New(*html)
	var tags []Tag

	for !sr.AtEnd() {
		char, err := sr.ReadNext()

		if err != nil {
			return nil, err
		}

		if char != '<' {
			continue
		}

		tag, err := TagAtPosition(sr)

		if err != nil {
			return nil, err
		}

		tags = append(tags, *tag)
	}

	if len(tags) == 0 {
		return nil, fmt.Errorf("No HTML found in \"%s\"", *html)
	}

	return &tags, nil
}
