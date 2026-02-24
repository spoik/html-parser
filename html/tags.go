package html

import "slices"

type Tags struct {
	Tags []*Tag
}

func NewTags(tags []*Tag) *Tags {
	if len(tags) == 0 {
		return nil
	}

	return &Tags{Tags: tags}
}

func (t *Tags) Find(tagType string) []*Tag {
	matches := make([]*Tag, 5)

	for _, tag := range t.Tags {
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
