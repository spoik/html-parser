package html

type Attribute struct {
	Name  string
	Value string
}

type Attributes struct {
	attributes map[string]*Attribute
}

func NewAttributes(atrs []*Attribute) *Attributes {
	attributes := make(map[string]*Attribute, len(atrs))

	for _, atr := range atrs {
		// If an attribute with this name has already been added, skip it.
		if attributes[atr.Name] != nil {
			continue
		}

		attributes[atr.Name] = atr
	}

	return &Attributes{attributes: attributes}
}
