package html

import "slices"

type Tags struct {
	tags []*Tag
}

func NewTags(tags []*Tag) *Tags {
	if len(tags) == 0 {
		return nil
	}

	return &Tags{tags: tags}
}

// Returns the tag at the given index. If there is no Tag at the index, nil is returned.
func (t *Tags) Get(index int) *Tag {
	if index > len(t.tags) {
		// TODO: Return error instead.
		return nil
	}

	return t.tags[index]
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
