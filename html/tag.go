package html

import "strings"

type Tag struct {
	Type       string
	Attributes *Attributes
	Text       string
	Tags       *Tags
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

func (t *Tag) FindTags(tagType string) []*Tag {
	if t.Tags == nil {
		return []*Tag{}
	}

	return t.Tags.Find(tagType)
}
