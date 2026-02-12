package parse

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

func parseValue(r *bufio.Reader) (string, error) {
	err := skipOpeningQuote(r)

	if err != nil {
		return "", err
	}

	var value strings.Builder

	for {
		bytes, err := r.Peek(1)

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return "", err
		}

		byte := bytes[0]

		if byte == '/' {
			selfClosingTag, err := nextTwoBytesAreSelfClosingTag(r)

			if err != nil {
				return "", err
			}

			if selfClosingTag {
				break
			}
		}

		if byte == '"' {
			_, err = r.Discard(1)

			if err != nil {
				return "", err
			}

			break
		}

		_, err = r.Discard(1)

		if err != nil {
			return "", err
		}

		value.WriteByte(byte)
	}

	return value.String(), nil
}

func skipOpeningQuote(r *bufio.Reader) error {
	bytes, err := r.Peek(1)

	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}

		return err
	}

	if bytes[0] == '"' {
		_, err = r.Discard(1)

		if err != nil {
			return err
		}
	}

	return nil
}

func nextTwoBytesAreSelfClosingTag(r *bufio.Reader) (bool, error) {
	bytes, err := r.Peek(2)

	if err != nil {
		if errors.Is(err, io.EOF) {
			return false, nil
		}

		return false, err
	}

	return bytes[1] == '>', nil
}
