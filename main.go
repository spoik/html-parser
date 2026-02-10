package main

import (
	"fmt"
	"slices"
	"strings"
)

var tagEndCharaacters = []byte{' ', '>', '/'}

// Returns the tag type. If html is "<html>", this function will return "html". Returns an error if a tag type can not be found.
func TagType(html *string, startPos int) (string, error) {
	length := len(*html)
	var tagType strings.Builder

	for i := startPos; i < length; i++ {
		char := (*html)[i]

		if slices.Contains(tagEndCharaacters, char) {
			break
		}

		tagType.WriteByte(char)
	}

	if tagType.Len() == 0 {
		return "", fmt.Errorf("Unable to find tag type in \"%s\"", *html)
	}

	return tagType.String(), nil
}
