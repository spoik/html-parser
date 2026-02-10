package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	html := "<a href=\"https://example.com\">Example</a>"
	ParseHtml(&html)
}

type Tag struct {
	Type string
}

func ParseHtml(html *string) (*Tag, error) {
	for index, char := range *html {
		if char == '<' {
			tagTypeResult, err := TagType(html, index+1)

			if err != nil {
				return nil, err
			}

			return &Tag{
				Type: tagTypeResult.TagType,
			}, nil
		}
	}

	return nil, fmt.Errorf("No HTML found in \"%s\"", *html)
}

type TagTypeResult struct {
	TagType string
	EndPos  int
}

var tagEndCharaacters = []byte{' ', '>', '/'}

// TagType returns the tag type in html starting at startPos and the position in html where the tag type definition ends. For example, if html is "<html>", this function will return "html" and 4. Returns an error if a tag type can not be found.
func TagType(html *string, startPos int) (*TagTypeResult, error) {
	var tagType strings.Builder
	index := startPos

	for index < len(*html) {
		char := (*html)[index]

		if slices.Contains(tagEndCharaacters, char) {
			break
		}

		tagType.WriteByte(char)
		index++
	}

	if tagType.Len() == 0 {
		return nil, fmt.Errorf("Unable to find tag type in \"%s\"", *html)
	}

	return &TagTypeResult{
		TagType: tagType.String(),
		EndPos:  index - 1,
	}, nil
}
