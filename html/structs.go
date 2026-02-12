package html

type Attribute struct {
	Name string
	Value string
}

type Tag struct {
	Type       string
	Attributes []*Attribute
}
