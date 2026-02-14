package parse

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

func getText(r *bufio.Reader) (string, error) {
	hasText, err := seekToText(r)

	if err != nil {
		return "", err
	}

	if hasText {
		return readText(r)
	}

	return "", nil
}

func seekToText(r *bufio.Reader) (hasText bool, err error) {
	for {
		bytes, err := r.Peek(1)

		if err != nil {
			if errors.Is(err, io.EOF) {
				return false, nil
			}

			return false, err
		}

		b := bytes[0]

		if !isTagEndChar(b) {
			_, err := r.Discard(1)

			if err != nil {
				return false, err
			}

			continue
		}

		// If this is a self closing tag, there is no text to extract.
		if b == '/' {
			return false, nil
		}

		_, err = r.Discard(1)

		if err != nil {
			return false, err
		}

		return true, nil
	}
}

func readText(r *bufio.Reader) (string, error) {
	var text strings.Builder

	for {
		bytes, err := r.Peek(1)

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

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
