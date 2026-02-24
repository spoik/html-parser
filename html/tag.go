package html

import "strings"

type Tag struct {
	Type       string
	Attributes *Attributes
	Text       string
	Tags       []*Tag
}

// Returns the Tag's Text concatenated with all it's children's Text.
func (t *Tag) FullText() string {
	var b strings.Builder

	b.WriteString(t.Text)

	for _, tag := range t.Tags {
		b.WriteString(tag.FullText())
	}

	return b.String()
}
