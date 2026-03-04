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
#### `html.Tags.Find()`
Returns all `html.Tag` instance of a specific type.

```go
import "github.com/spoik/html-parser/parse"

func main() {
	html := "<a>Example 1</a><a>Example 2</a><p>Example 3</p>"
	tags, _ := parse.ParseHtml(&html)

    tags := tags.Find('a')

    // tags will be []*Tag{{ html.Tag{Type: "a", Text: "Example 1"}, html.Tag{Type: "a", Text: "Example 2" }}}
}
```

#### `html.Tags.Get()`
Returns the `html.Tag` at a given index.
```go
import "github.com/spoik/html-parser/parse"

func main() {
	html := "<a>Example 1</a><p>Example 2</p>"
	tags, _ := parse.ParseHtml(&html)

    tag, _:= tags.Get(1)

    // tag will be *html.Tag{Type: "p", Text: "Example 2"}
}
```

#### `html.Tags.Len()`
Returns the number of immediate children `html.Tag`s.
```go
import "github.com/spoik/html-parser/parse"

func main() {
	html := "<p>Example 2</p><p>Example <span>2</span></p>"
	tags, _ := parse.ParseHtml(&html)

    length := tags.Len()

    // length == 2
}
```

#### `html.Tags.FullLen()`
Returns the total number of children `html.Tag`s.
```go
import "github.com/spoik/html-parser/parse"

func main() {
	html := "<p>Example 2</p><p>Example <span>2</span></p>"
	tags, _ := parse.ParseHtml(&html)

    length := tags.FullLen()

    // length == 3
}
```
