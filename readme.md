# Basic HTML Parser
Does this need to exist? No. I created this to get more familiar with golang.

## Features & Usage
### Initialization
Use `html.ParseHtml` to parse an HTML string. Returns a `html.Tags` that represents the HTML.

```go
import "github.com/spoik/html-parser/parse"

func main() {
	html := "<a href=\"https://example.com\">Example</a>"
	tags, err := parse.ParseHtml(&html)
}
```

### Querying
`html.Tags` have a `Find` method that returns all `Tag` instance of a specific type.

```go
import "github.com/spoik/html-parser/parse"

func main() {
	html := "<a>Example 1</a><a>Example 2</a><p>Example 3</p>"
	tags, _ := parse.ParseHtml(&html)

    tags := tags.Find('a')

    // tags will be []*Tag{{ html.Tag{Type: "a", Text: "Example 1"}, html.Tag{Type: "a", Text: "Example 2" }}}
}
```
