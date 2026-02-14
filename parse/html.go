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

// Returns a slice of html.Tag instances that represent the html provide in s.
func ParseHtml(s *string) ([]*html.Tag, error) {
	r := bufio.NewReaderSize(
		stringreader.New(*s),
		2,
	)

	nextTag, checkErr := tagIterator(r)

	// TODO: Change this slice to have a fixed size to avoid constantly resizing the slice.
	// Not sure what the best size would be though: make([]*html.Tag, 50)
	var tags []*html.Tag

	for tag := range nextTag {
		tags = append(tags, tag)
	}

	if err := checkErr(); err != nil {
		return nil, fmt.Errorf("Error parsing HTML: %w", err)
	}

	if len(tags) == 0 {
		return nil, fmt.Errorf("No HTML tags found in \"%s\"", *s)
	}

	return tags, nil
}

// Returns an iterator that returns each tag present in the reader's text and an closure that
// returns any errors that occurred. The error closure should be checked before using the values
// returned by the iterator.
func tagIterator(r *bufio.Reader) (iter.Seq[*html.Tag], func() error) {
	var err error

	seq := func(yield func(*html.Tag) bool) {
		for {
			e := seekToNextTag(r)

			if e != nil {
				// If the end of the reader has been reached, there are no more tags to be found.
				// In this case, io.EOF isn't an error. It is instead an indication that all tags
				// have been found successfully. No need to return an error in this case.
				if errors.Is(e, io.EOF) {
					break
				}

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

// Advance the reader to where the next tag begins.
func seekToNextTag(r *bufio.Reader) error {
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
