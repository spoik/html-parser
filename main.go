package main

import (
	"fmt"

	"github.com/spoik/html-parser/parse"
)

func main() {
	html := "<p>Example 2</p><section>Section <p>Paragraph</p>"
	tags, _ := parse.ParseHtml(&html)

	for t := range tags.AllTagsDeep() {
		fmt.Println(t)
	}
}
