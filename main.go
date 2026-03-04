package main

import (
	"fmt"

	"github.com/spoik/html-parser/parse"
)

func main() {
	html := "<p>Example 1<span>Example 2</span><span>Example 3</span></p>"
	tags, _ := parse.ParseHtml(&html)

	tag, _ := tags.Get(0)
	fmt.Println(tag.FullText())
}
