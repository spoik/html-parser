package html

import (
	"fmt"
	"strings"
)

var SelfClosingTags = []string{"br", "hr", "img", "input", "link", "meta", "source"}

type Tag struct {
	Type       string
	Attributes *Attributes
	Text       string
	Tags       *Tags
}

// Returns a string of HTML representing this Tag and all its children.
func (t *Tag) String() string {
	var tagsStr string

	if t.Tags == nil {
		tagsStr = ""
	} else {
		tagsStr = t.Tags.String()
	}

	return fmt.Sprintf(
		"<%s>%s%s</%s>",
		t.Type,
		t.Text,
		tagsStr,
		t.Type,
	)
}

// Returns the Tag's Text concatenated with all it's children's Text.
func (t *Tag) FullText() string {
	var b strings.Builder

	b.WriteString(t.Text)

	if t.Tags != nil {
		for _, tag := range t.Tags.tags {
			b.WriteString(tag.FullText())
		}
	}

	return b.String()
}

// Returns the html.Attribute belonging to the html.Tag with the given name. If the html.Tag has no matching attribute, nil is returned.
func (t *Tag) Attribute(name string) *Attribute {
	return t.Attributes.Attribute(name)
}

// Returns all children tag of the given type.
func (t *Tag) FindTags(tagType string) []*Tag {
	if t.Tags == nil {
		return []*Tag{}
	}

	return t.Tags.Find(tagType)
}
