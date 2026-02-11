package main

import "github.com/spoik/html-parser/parse"

func main() {
	html := "<a href=\"https://example.com\">Example</a>"
	parse.ParseHtml(&html)
}
