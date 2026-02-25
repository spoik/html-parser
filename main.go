package main

import (
	"fmt"

	"github.com/spoik/html-parser/parse"
)

func main() {
	html := "<a href=\"https://example.com\">Example</a>"
	tags, err := parse.ParseHtml(&html)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	tag, err := tags.Get(0)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", tag.Attribute("href").Value)
}
