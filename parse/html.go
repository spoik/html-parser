package parse

import "fmt"

type Tag struct {
	Type string
}

func ParseHtml(html *string) (*[]Tag, error) {
	var tags []Tag

	for index := 0; index < len(*html); index++ {
		char := (*html)[index]

		if char != '<' {
			continue
		}

		result, err := TagAtPosition(html, index+1)

		if err != nil {
			return nil, err
		}

		tags = append(tags, *result.Tag)

		index = result.EndPos
	}

	if len(tags) == 0 {
		return nil, fmt.Errorf("No HTML found in \"%s\"", *html)
	}

	return &tags, nil
}
