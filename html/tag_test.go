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
				Tags: html.NewTags(nil),
			},
			ExpectedString: "Hello",
		},
		{
			Name: "Tag with text and child tag with text",
			Tag: html.Tag{
				Text: "Hello",
				Tags: html.NewTags([]*html.Tag{{Text: "World"}}),
			},
			ExpectedString: "HelloWorld",
		},
		{
			Name: "Tag with text and child tag with text and leading space",
			Tag: html.Tag{
				Text: "Hello",
				Tags: html.NewTags([]*html.Tag{{Text: " World"}}),
			},
			ExpectedString: "Hello World",
		},
		{
			Name: "Tag with text and child tag with text and leading space",
			Tag: html.Tag{
				Text: "Hello",
				Tags: html.NewTags([]*html.Tag{{Text: " World"}}),
			},
			ExpectedString: "Hello World",
		},
		{
			Name: "Tag with text and multiple child tags",
			Tag: html.Tag{
				Text: "Hello",
				Tags: html.NewTags(
					[]*html.Tag{
						{Text: "There"},
						{Text: "World"},
					},
				),
			},
			ExpectedString: "HelloThereWorld",
		},
		{
			Name: "Tag with text and multiple child tags with their own child tags",
			Tag: html.Tag{
				Text: "Hello",
				Tags: html.NewTags(
					[]*html.Tag{
						{
							Text: "There",
							Tags: html.NewTags([]*html.Tag{{Text: "How"}}),
						},
						{
							Text: "Are",
							Tags: html.NewTags([]*html.Tag{{Text: "You"}}),
						},
					},
				),
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

func TestAttribute(t *testing.T) {
	type testCase struct {
		Name           string
		Tag            *html.Tag
		AttrName       string
		ExpectedResult *html.Attribute
	}

	testCases := []testCase{
		{
			Name: "Attribute is present in the attirbutes.",
			Tag: &html.Tag{
				Attributes: html.NewAttributes([]*html.Attribute{
					{Name: "id", Value: "profile"},
				}),
			},
			AttrName:       "id",
			ExpectedResult: &html.Attribute{Name: "id", Value: "profile"},
		},
		{
			Name: "Attribute is not present in the attirbutes.",
			Tag: &html.Tag{
				Attributes: html.NewAttributes([]*html.Attribute{
					{Name: "id", Value: "profile"},
				}),
			},
			AttrName:       "class",
			ExpectedResult: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			result := testCase.Tag.Attribute(testCase.AttrName)

			assert.Equal(t, testCase.ExpectedResult, result)
		})
	}
}

func TestFindTags(t *testing.T) {
	type testCase struct {
		Name           string
		Tag            html.Tag
		TagType        string
		ExpectedResult []*html.Tag
	}

	testCases := []testCase{
		{
			Name:           "With no tags.",
			Tag:            html.Tag{},
			TagType:        "a",
			ExpectedResult: []*html.Tag{},
		},
		{
			Name: "With one matching tag.",
			Tag: html.Tag{
				Tags: html.NewTags([]*html.Tag{
					{Type: "a"},
				}),
			},
			TagType: "a",
			ExpectedResult: []*html.Tag{
				{Type: "a"},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			result := testCase.Tag.FindTags(testCase.TagType)

			assert.Equal(t, testCase.ExpectedResult, result)
		})
	}
}
