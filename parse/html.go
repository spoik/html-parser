package parse

import (
	"errors"
	"fmt"
	"io"

	"github.com/spoik/html-parser/stringreader"
)

type Tag struct {
	Type string
}

func ParseHtml(html *string) (*[]Tag, error) {
	sr := stringreader.New(*html)
	var tags []Tag
	bytes := make([]byte, 1)

	for {
		_, err := sr.Read(bytes)

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return nil, err
		}

		char := bytes[0]

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
