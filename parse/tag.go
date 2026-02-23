package parse

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/spoik/html-parser/html"
)

var tagEndBytes = []byte{'>', '/'}

func isTagEndChar(b byte) bool {
	return slices.Contains(tagEndBytes, b)
}

var attributeDeliminer byte = ' '

func isAttributeDeliminer(b byte) bool {
	return b == attributeDeliminer
}

func ParseTag(r *bufio.Reader) (*html.Tag, error) {
	tagType, err := parseTagType(r)

	if err != nil {
		return nil, err
	}

	attributes, err := getAttributes(r)

	if err != nil {
		return nil, err
	}

	text, err := getText(r)

	if err != nil {
		return nil, err
	}

	internalTags, err := parseInternalTags(r)

	if err != nil {
		return nil, err
	}

	err = parseClosingTag(tagType, r)

	if err != nil {
		return nil, err
	}

	tag := &html.Tag{
		Type:       tagType,
		Text:       text,
		Attributes: attributes,
		Tags:       internalTags,
	}

	return tag, nil
}

func parseTagType(r *bufio.Reader) (string, error) {
	var typeBuilder strings.Builder

	for {
		bytes, err := r.Peek(1)

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return "", err
		}

		byte := bytes[0]

		if isTagEndChar(byte) || isAttributeDeliminer(byte) {
			break
		}

		_, err = r.Discard(1)

		if err != nil {
			return "", err
		}

		err = typeBuilder.WriteByte(byte)

		if err != nil {
			return "", err
		}
	}

	if typeBuilder.Len() == 0 {
		return "", fmt.Errorf("Unable to find tag.")
	}

	return typeBuilder.String(), nil
}

func parseInternalTags(r *bufio.Reader) ([]*html.Tag, error) {
	// TODO: Initialize childTags with a starting size to
	//minimize how often the slice is resized
	var childTags []*html.Tag

	for {
		bytes, err := r.Peek(2)

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return nil, err
		}

		// If the next two bytes are a tag's closing tags, stop parsing internal tags.
		// There are no more internal tags.
		if string(bytes) == "</" {
			break
		}

		// If the next two bytes are a tag's self closing bytes, stop parsing internal tags.
		// There are no more internal tags.
		if string(bytes) == "/>" {
			break
		}

		// If the next byte is not "<", continue advancing the bufio.Reader to continue
		// the search for internal tags.
		if bytes[0] != '<' {
			_, err = r.Discard(1)

			if err != nil {
				return nil, err
			}

			continue
		}

		// If the next byte is '<', this indicates the beginning of an internal tag. Discard
		// the "<" and call ParseTag to parse the internal tag.
		_, err = r.Discard(1)

		if err != nil {
			return nil, err
		}

		childTag, err := ParseTag(r)

		if err != nil {
			return nil, err
		}

		childTags = append(childTags, childTag)
	}

	return childTags, nil
}

// Advance the bufio.Reader past the closing tag. This would advance the reader just past the
// "&lt;/a&gt;" in "&lt;a&gt;text&lt;/a&gt;&lt;p&gt;" to the "&lt;" in "&lt;p&gt;".
//
// If tag currently being parsed in s a self closing tag, this function will advance the
// reader past the "/&gt;" in "&lt;img/&gt;&lt;a" to the "&lt;" in "&lt;a&gt;".
func parseClosingTag(tagType string, r *bufio.Reader) error {
	// Peak the next two bytes to see if they are the beginning of the closing
	// tag (the "</" in "</p>") or the end of a self closing tag (the "/>" in "<img/>")
	bytes, err := r.Peek(2)

	if errors.Is(err, io.EOF) {
		return nil
	}

	if err != nil {
		return err
	}

	// If the next two bytes are the end of a self closing tag. Example: <img/>
	if string(bytes) == "/>" {
		// Discard the "/>" bytes and return. The end tag has now been parsed.
		_, err = r.Discard(2)

		if err != nil {
			return err
		}

		return nil
	}

	// If the next two bytes are not the beginning of a closing tag. Example: </p>
	if string(bytes) != "</" {
		// Discard the next byte
		_, err = r.Discard(1)

		if err != nil {
			return err
		}

		// Continue the search for the closing tag.
		return parseClosingTag(tagType, r)
	}

	// If the beginning of the closing tag has been found (</), Discard it.
	_, err = r.Discard(2)

	// Peek next n bytes after the </ to see if it matches the tagType.
	endTagTypeBytes, err := r.Peek(len(tagType))

	if errors.Is(err, io.EOF) {
		// Since the end of file has been reached, discard what's left in the buffer. There is
		// Nothing left to parse.
		_, err := r.Discard(len(tagType))

		if err != nil && !errors.Is(err, io.EOF) {
			return err
		}

		return nil
	}

	endTagType := string(endTagTypeBytes)

	if endTagType != tagType {
		return fmt.Errorf(
			"End tag type does not matching opening tag type. Expected end tag type %s, but got %s",
			tagType,
			string(endTagType),
		)
	}

	if err != nil {
		return err
	}

	// The closing tag has been successfully parsed.
	return nil
}
