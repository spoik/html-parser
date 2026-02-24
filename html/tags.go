package html


type Tags struct {
	Tags []*Tag
}

func NewTags(tags []*Tag) *Tags {
	if len(tags) == 0 {
		return nil
	}

	return &Tags{Tags: tags}
}
