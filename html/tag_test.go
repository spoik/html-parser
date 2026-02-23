package html_test

import (
	"testing"

	"github.com/spoik/html-parser/html"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Name           string
	Tag            html.Tag
	ExpectedString string
}

func TestFullText(t *testing.T) {
	testCases := []testCase{
		{
			Name:           "Single tag with no text",
			Tag:            html.Tag{},
			ExpectedString: "",
		},
		{
			Name:           "Single tag with text",
			Tag:            html.Tag{Text: "Hello"},
			ExpectedString: "Hello",
		},
		{
			Name: "Tag with text and child tag with no text",
			Tag: html.Tag{
				Text: "Hello",
				Tags: []*html.Tag{{}},
			},
			ExpectedString: "Hello",
		},
		{
			Name: "Tag with text and child tag with text",
			Tag: html.Tag{
				Text: "Hello",
				Tags: []*html.Tag{{Text: "World"}},
			},
			ExpectedString: "HelloWorld",
		},
		{
			Name: "Tag with text and child tag with text and leading space",
			Tag: html.Tag{
				Text: "Hello",
				Tags: []*html.Tag{{Text: " World"}},
			},
			ExpectedString: "Hello World",
		},
		{
			Name: "Tag with text and child tag with text and leading space",
			Tag: html.Tag{
				Text: "Hello",
				Tags: []*html.Tag{{Text: " World"}},
			},
			ExpectedString: "Hello World",
		},
		{
			Name: "Tag with text and multiple child tags",
			Tag: html.Tag{
				Text: "Hello",
				Tags: []*html.Tag{
					{Text: "There"},
					{Text: "World"},
				},
			},
			ExpectedString: "HelloThereWorld",
		},
		{
			Name: "Tag with text and multiple child tags with their own child tags",
			Tag: html.Tag{
				Text: "Hello",
				Tags: []*html.Tag{
					{
						Text: "There",
						Tags: []*html.Tag{{Text: "How"}},
					},
					{
						Text: "Are",
						Tags: []*html.Tag{{Text: "You"}},
					},
				},
			},
			ExpectedString: "HelloThereHowAreYou",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			result := testCase.Tag.FullText()

			assert.Equal(t, testCase.ExpectedString, result)
		})
	}
}
