package html

import (
	"errors"
	"reflect"
)

type Tags struct {
	tags     []*Tag
	tagIndex *TagIndex
}

func EmptyTags() *Tags {
	return &Tags{tags: []*Tag{}, tagIndex: &TagIndex{}}
}
func NewTags(tags []*Tag) *Tags {
	return &Tags{tags: tags, tagIndex: &TagIndex{}}
}

func TagsWithIndex(tags []*Tag, tagIndex *TagIndex) *Tags {
	return &Tags{
		tags:     tags,
		tagIndex: tagIndex,
	}
}

var NoTagAtIndex = errors.New("No tag at index.")

// Returns the tag at the given index. If there is no Tag at the index, nil is returned.
func (t *Tags) Get(index int) (*Tag, error) {
	if t.tags == nil {
		return nil, NoTagAtIndex
	}

	if index > t.Length() {
		return nil, NoTagAtIndex
	}

	return t.tags[index], nil
}

func (t *Tags) Equal(other *Tags) bool {
	return reflect.DeepEqual(t.tags, other.tags)
}

// Returns the number of Tags.
func (t *Tags) Length() int {
	return len(t.tags)
}

// Returns the number of Tags in addition to the number of child tags each Tag has.
func (t *Tags) FullLength() int {
	fullLen := t.Length()

	for _, tag := range t.tags {
		if tag.Tags == nil {
			continue
		}

		fullLen += tag.Tags.FullLength()
	}

	return fullLen
}

func (t *Tags) Find(tagType string) []*Tag {
	if t.tagIndex == nil {
		panic("Unable to Find tag by type. TagIndex is nil.")
	}

	return t.tagIndex.Get(tagType)
}
