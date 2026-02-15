package parse

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

var errNoText = errors.New("No text found in tag")

func getText(r *bufio.Reader) (string, error) {
	err := seekToText(r)

	if errors.Is(err, errNoText) {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	return readText(r, &strings.Builder{})
}

func seekToText(r *bufio.Reader) error {
	bytes, err := r.Peek(1)

	if errors.Is(err, io.EOF) {
		return errNoText
	}

	if err != nil {
		return err
	}

	b := bytes[0]

	if !isTagEndChar(b) {
		_, err := r.Discard(1)

		if err != nil {
			return err
		}

		return seekToText(r)
	}

	// If this is a self closing tag, there is no text to extract.
	if b == '/' {
		return errNoText
	}

	_, err = r.Discard(1)

	if err != nil {
		return err
	}

	return nil
}

func readText(r *bufio.Reader, builder *strings.Builder) (string, error) {
	bytes, err := r.Peek(1)

	if errors.Is(err, io.EOF) {
		return builder.String(), nil
	}

	if err != nil {
		return "", err
	}

	b := bytes[0]

	if b == '<' {
		return builder.String(), nil
	}

	_, err = r.Discard(1)

	if err != nil {
		return "", err
	}

	err = builder.WriteByte(b)

	if err != nil {
		return "", err
	}

	return readText(r, builder)
}
