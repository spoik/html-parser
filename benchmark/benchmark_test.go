package benchmark_test

import (
	"os"
	"testing"

	"github.com/spoik/html-parser/html"
	"github.com/spoik/html-parser/parse"
)

func testFileContents(b *testing.B) *string {
	content, err := os.ReadFile("test.html")

	if err != nil {
		b.Fatalf("Failed to open test html file: %v", err)
	}

	str := string(content)

	return &str
}

func testFileTags(b *testing.B) html.Tags {
	html := testFileContents(b)
	tags, err := parse.ParseHtml(html)

	if err != nil {
		b.Fatalf("Failed to parse Html: %v", err)
	}

	return tags
}

func BenchmarkParseHtml(b *testing.B) {
	html := testFileContents(b)
	b.ResetTimer()

	for b.Loop() {
		_, err := parse.ParseHtml(html)

		if err != nil {
			b.Fatalf("Failed to parse Html: %v", err)
		}
	}
}

func BenchmarkTagsFind(b *testing.B) {
	tags := testFileTags(b)
	b.ResetTimer()

	for b.Loop() {
		tags.Find("a")
	}
}
