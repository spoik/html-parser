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

	return readText(r)
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

func readText(r *bufio.Reader) (string, error) {
	var text strings.Builder

	for {
		bytes, err := r.Peek(1)

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return "", err
		}

		b := bytes[0]

		if b == '<' {
			break
		}

		_, err = r.Discard(1)

		if err != nil {
			return "", err
		}

		err = text.WriteByte(b)

		if err != nil {
			return "", err
		}
	}

	return text.String(), nil
}
