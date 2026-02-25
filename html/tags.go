package html

import (
	"errors"
	"slices"
)

type Tags struct {
	tags []*Tag
}

func NewTags(tags []*Tag) *Tags {
	if len(tags) == 0 {
		return nil
	}

	return &Tags{tags: tags}
}

var NoTagAtIndex = errors.New("No tag at index.")

// Returns the tag at the given index. If there is no Tag at the index, nil is returned.
func (t *Tags) Get(index int) (*Tag, error) {
	if t.tags == nil {
		return nil, NoTagAtIndex
	}

	if index > len(t.tags) {
		return nil, NoTagAtIndex
	}

	return t.tags[index], nil
}

// Returns the number of Tags.
func (t *Tags) Length() int {
	return len(t.tags)
}

func (t *Tags) Find(tagType string) []*Tag {
	matches := make([]*Tag, 5)

	for _, tag := range t.tags {
		if tag.Type == tagType {
			tagCopy := *tag
			tagCopy.Tags = nil

			matches = append(matches, &tagCopy)
		}

		matches = append(matches, tag.FindTags(tagType)...)
	}

	matches = slices.DeleteFunc(matches, func(tag *Tag) bool {
		return tag == nil
	})

	return matches
}
