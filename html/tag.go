package html

import (
	"fmt"
	"strings"
)

var SelfClosingTags = []string{"br", "hr", "img", "input", "link", "meta", "source"}

type Tag struct {
	Type       string
	Attributes Attributes
	Text       string
	Tags       Tags
}

// Returns a string of HTML representing this Tag and all its children.
func (t Tag) String() string {
	return fmt.Sprintf(
		"<%s>%s%s</%s>",
		t.Type,
		t.Text,
		t.Tags,
		t.Type,
	)
}

// Returns the Tag's Text concatenated with all it's children's Text.
func (t *Tag) FullText() string {
	var b strings.Builder

	b.WriteString(t.Text)

	for _, tag := range t.Tags.tags {
		b.WriteString(tag.FullText())
	}

	return b.String()
}

// Returns the html.Attribute belonging to the html.Tag with the given name. If the html.tag has a matching attribute,
// ok will be true. If the html.Tag has no matching attribute, ok will be false.
func (t *Tag) Attribute(name string) (attribute Attribute, ok bool) {
	return t.Attributes.Attribute(name)
}

// Returns all children tag of the given type.
func (t *Tag) FindTags(tagType string) []Tag {
	return t.Tags.Find(tagType)
}
