package html

import (
	"errors"
	"reflect"
)

type Tags struct {
	tags     []*Tag
	tagIndex *TagIndex
}

type NewTagsOpts struct {
	Tags     []*Tag
	TagIndex *TagIndex
}

func EmptyTags() *Tags {
	return NewTags(NewTagsOpts{})
}

func NewTags(o NewTagsOpts) *Tags {
	tags := o.Tags

	if tags == nil {
		tags = []*Tag{}
	}

	tagIndex := o.TagIndex

	if tagIndex == nil {
		tagIndex = &TagIndex{}
	}
	return &Tags{tags: tags, tagIndex: tagIndex}
}

var NoTagAtIndex = errors.New("No tag at index.")

// Returns the tag at the given index. If there is no Tag at the index, nil is returned.
func (t *Tags) Get(index int) (*Tag, error) {
	if t.tags == nil {
		return nil, NoTagAtIndex
	}

	if index > t.Len() {
		return nil, NoTagAtIndex
	}

	return t.tags[index], nil
}

func (t *Tags) Equal(other *Tags) bool {
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
		if tag.Tags == nil {
			continue
		}

		fullLen += tag.Tags.FullLen()
	}

	return fullLen
}

func (t *Tags) Find(tagType string) []*Tag {
	if t.tagIndex == nil {
		panic("Unable to Find tag by type. TagIndex is nil.")
	}

	return t.tagIndex.Get(tagType)
}
