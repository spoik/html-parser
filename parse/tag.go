package parse

import (
	"fmt"
	"slices"
	"strings"
)

type TagAtPositionResult struct {
	Tag    *Tag
	EndPos int
}

func TagAtPosition(html *string, position int) (*TagAtPositionResult, error) {
	result, err := TagType(html, position)

	if err != nil {
		return nil, err
	}

	tag := Tag{Type: result.TagType}

	return &TagAtPositionResult{
		Tag:    &tag,
		EndPos: result.EndPos,
	}, nil
}

type TagTypeResult struct {
	TagType string
	EndPos  int
}

var tagTypeEndCharaacters = []byte{' ', '>', '/'}

// TagType returns the tag type in html starting at position and the position in html where the tag type definition ends. For example, if html is "<html>", this function will return "html" and 4. Returns an error if a tag type can not be found.
func TagType(html *string, position int) (*TagTypeResult, error) {
	var tagType strings.Builder
	index := position

	for index < len(*html) {
		char := (*html)[index]

		if slices.Contains(tagTypeEndCharaacters, char) {
			break
		}

		tagType.WriteByte(char)
		index++
	}

	if tagType.Len() == 0 {
		return nil, fmt.Errorf("Unable to find tag type in \"%s\" starting at position %d.", *html, position)
	}

	return &TagTypeResult{
		TagType: tagType.String(),
		EndPos:  index - 1,
	}, nil
}
