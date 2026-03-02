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

	aTags := tags.Find("im")

	fmt.Println(len(aTags))
	fmt.Printf("%+v", aTags)
}
