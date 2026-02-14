package parse

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"iter"

	"github.com/spoik/html-parser/html"
	"github.com/spoik/html-parser/stringreader"
)

func ParseHtml(htmlStr *string) ([]*html.Tag, error) {
	r := bufio.NewReaderSize(
		stringreader.New(*htmlStr),
		2,
	)

	nextTag, checkErr := tagIterator(r)

	// TODO: Change this slice to have a fixed size to avoid constantly resizing the slice.
	// Not sure what the best size would be though: make([]*html.Tag, 50)
	var tags []*html.Tag

	for tag := range nextTag {
		println("Found tag:", tag)
		tags = append(tags, tag)
	}

	if err := checkErr(); err != nil {
		if !errors.Is(err, io.EOF) {
			return nil, fmt.Errorf("Error parsing HTML: %w", err)
		}
	}

	if len(tags) == 0 {
		return nil, fmt.Errorf("No HTML found in \"%s\"", *htmlStr)
	}

	return tags, nil
}

func tagIterator(r *bufio.Reader) (iter.Seq[*html.Tag], func() error) {
	var err error

	seq := func(yield func(*html.Tag) bool) {
		for {
			e := seekToTag(r)

			if e != nil {
				err = e
				break
			}

			tag, e := ParseTag(r)

			if e != nil {
				err = e
				break
			}

			if !yield(tag) {
				break
			}
		}
	}

	return seq, func() error { return err }
}

func seekToTag(r *bufio.Reader) error {
	for {
		byte, err := r.ReadByte()

		if err != nil {
			return err
		}

		if byte == '<' {
			return nil
		}
	}
}
