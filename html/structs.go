package html

type Attribute struct {
	Name string
}

type Tag struct {
	Type       string
	Attributes []*Attribute
}
