package html

import (
	"errors"
	"iter"
	"reflect"
	"strings"
)

type Tags struct {
	tags     []Tag
	tagIndex *TagIndex
}

type newTagsOption func(*Tags)

func WithIndex(i *TagIndex) newTagsOption {
	return func(t *Tags) {
		t.tagIndex = i
	}
}

func NewTags(tags []Tag, options ...newTagsOption) Tags {
	t := Tags{tags: tags}

	for _, option := range options {
		option(&t)
	}

	return t
}

func (t Tags) String() string {
	var b strings.Builder

	for _, tag := range t.tags {
		b.WriteString(tag.String())
	}

	return b.String()
}

var NoTagAtIndex = errors.New("No tag at index.")

// Returns the tag at the given index. If there is no Tag at the index, nil is returned.
func (t *Tags) Get(index int) (Tag, error) {
	if t.tags == nil {
		return Tag{}, NoTagAtIndex
	}

	if index > t.Len() {
		return Tag{}, NoTagAtIndex
	}

	return t.tags[index], nil
}

func (t *Tags) AllTags() iter.Seq2[int, Tag] {
	return func(yield func(int, Tag) bool) {
		for i, tag := range t.tags {
			if !yield(i, tag) {
				return
			}
		}
	}
}

func (t *Tags) Equal(other Tags) bool {
	return reflect.DeepEqual(t.tags, other.tags)
}

// Returns the number of Tags.
func (t *Tags) Len() int {
	return len(t.tags)
}

// Returns the number of Tags in addition to the number of child tags each Tag has.
func (t *Tags) FullLen() int {
	fullLen := t.Len()

	for _, tag := range t.tags {
		fullLen += tag.Tags.FullLen()
	}

	return fullLen
}

func (t *Tags) Find(tagType string) []Tag {
	if t.tagIndex == nil {
		panic("Unable to Find tag by type. TagIndex is nil.")
	}

	return t.tagIndex.Get(tagType)
}
