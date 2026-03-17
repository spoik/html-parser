package html

type TagIndex struct {
	byType map[string][]*Tag
}

func (i *TagIndex) Get(t string) []Tag {
	tags := i.byType[t]

	result := make([]Tag, len(tags))

	for i, tag := range tags {
		tagCopy := *tag
		tagCopy.Tags = Tags{}

		result[i] = tagCopy
	}

	return result
}

func (i *TagIndex) Add(t *Tag) {
	if i.byType == nil {
		i.byType = make(map[string][]*Tag)
	}

	i.byType[t.Type] = append(i.byType[t.Type], t)
}

func (i *TagIndex) AddAll(t []Tag) {
	for _, tag := range t {
		i.Add(&tag)
		i.AddAll(tag.Tags.tags)
	}
}
