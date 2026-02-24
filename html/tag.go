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
		for _, tag := range t.Tags.Tags {
			b.WriteString(tag.FullText())
		}
	}

	return b.String()
}

func (t *Tag) Attribute(name string) *Attribute {
	return t.Attributes.Attribute(name)
}
