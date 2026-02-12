package parse

import (
	"bufio"
	"errors"
	"fmt"
	"io"

	"github.com/spoik/html-parser/html"
	"github.com/spoik/html-parser/stringreader"
)

func ParseHtml(htmlStr *string) (*[]*html.Tag, error) {
	r := bufio.NewReaderSize(
		stringreader.New(*htmlStr),
		2,
	)

	var tags []*html.Tag
	bytes := make([]byte, 1)

	for {
		_, err := r.Read(bytes)

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, err
		}

		char := bytes[0]

		if char != '<' {
			continue
		}

		tag, err := ParseTag(r)

		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	if len(tags) == 0 {
		return nil, fmt.Errorf("No HTML found in \"%s\"", *htmlStr)
	}

	return &tags, nil
}
