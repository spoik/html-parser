package html_test

import (
	"testing"

	"github.com/spoik/html-parser/html"
	"github.com/stretchr/testify/assert"
)

func TestFullText(t *testing.T) {
	type testCase struct {
		Name           string
		Tag            html.Tag
		ExpectedString string
	}

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
				Tags: html.NewTags([]html.Tag{{Type: "a"}}),
			},
			ExpectedString: "Hello",
		},
		{
			Name: "Tag with text and child tag with text",
			Tag: html.Tag{
				Text: "Hello",
				Tags: html.NewTags([]html.Tag{{Text: "World"}}),
			},
			ExpectedString: "HelloWorld",
		},
		{
			Name: "Tag with text and child tag with text and leading space",
			Tag: html.Tag{
				Text: "Hello",
				Tags: html.NewTags([]html.Tag{{Text: " World"}}),
			},
			ExpectedString: "Hello World",
		},
		{
			Name: "Tag with text and child tag with text and leading space",
			Tag: html.Tag{
				Text: "Hello",
				Tags: html.NewTags([]html.Tag{{Text: " World"}}),
			},
			ExpectedString: "Hello World",
		},
		{
			Name: "Tag with text and multiple child tags",
			Tag: html.Tag{
				Text: "Hello",
				Tags: html.NewTags([]html.Tag{
					{Text: "There"},
					{Text: "World"},
				}),
			},
			ExpectedString: "HelloThereWorld",
		},
		{
			Name: "Tag with text and multiple child tags with their own child tags",
			Tag: html.Tag{
				Text: "Hello",
				Tags: html.NewTags([]html.Tag{
					{
						Text: "There",
						Tags: html.NewTags([]html.Tag{{Text: "How"}}),
					},
					{
						Text: "Are",
						Tags: html.NewTags([]html.Tag{{Text: "You"}}),
					},
				}),
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
		Name              string
		Tag               html.Tag
		AttrName          string
		ExpectedAttribute html.Attribute
		EpxectedOk        bool
	}

	testCases := []testCase{
		{
			Name: "Attribute is present in the attributes.",
			Tag: html.Tag{
				Attributes: html.NewAttributes([]html.Attribute{
					{Name: "id", Value: "profile"},
				}),
			},
			AttrName:          "id",
			ExpectedAttribute: html.Attribute{Name: "id", Value: "profile"},
			EpxectedOk:        true,
		},
		{
			Name: "Attribute is not present in the attributes.",
			Tag: html.Tag{
				Attributes: html.NewAttributes([]html.Attribute{
					{Name: "id", Value: "profile"},
				}),
			},
			AttrName:          "class",
			ExpectedAttribute: html.Attribute{},
			EpxectedOk:        false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			result, ok := testCase.Tag.Attribute(testCase.AttrName)

			assert.Equal(t, testCase.EpxectedOk, ok)
			assert.Equal(t, testCase.ExpectedAttribute, result)
		})
	}
}

func TestFindTags(t *testing.T) {
	type testCase struct {
		Name           string
		Tag            html.Tag
		TagType        string
		ExpectedResult []html.Tag
	}

	testCases := []testCase{
		{
			Name: "With no tags.",
			Tag: html.Tag{
				Tags: createTagsWithIndex([]html.Tag{}),
			},
			TagType:        "a",
			ExpectedResult: []html.Tag{},
		},
		{
			Name: "With one matching tag.",
			Tag: html.Tag{
				Tags: createTagsWithIndex([]html.Tag{
					{Type: "a"},
				}),
			},
			TagType: "a",
			ExpectedResult: []html.Tag{
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

func TestTagString(t *testing.T) {
	type testCase struct {
		Tag            html.Tag
		ExpectedResult string
	}

	testCases := []testCase{
		{
			Tag: html.Tag{
				Type: "a",
			},
			ExpectedResult: "<a></a>",
		},
		{
			Tag: html.Tag{
				Type: "div",
				Text: "Testing 123",
			},
			ExpectedResult: "<div>Testing 123</div>",
		},
		{
			Tag: html.Tag{
				Type: "section",
				Text: "Section text",
				Tags: html.NewTags([]html.Tag{
					{
						Type: "div",
						Text: "Div text",
					},
					{
						Type: "span",
						Text: "Span text",
					},
				}),
			},
			ExpectedResult: "<section>Section text<div>Div text</div><span>Span text</span></section>",
		},
		{
			Tag: html.Tag{
				Type: "section",
				Text: "Section text",
				Tags: html.NewTags([]html.Tag{
					{
						Type: "div",
						Text: "Div text",
						Tags: html.NewTags([]html.Tag{
							{
								Type: "span",
								Text: "Span text",
							},
						}),
					},
				}),
			},
			ExpectedResult: "<section>Section text<div>Div text<span>Span text</span></div></section>",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.ExpectedResult, func(t *testing.T) {
			t.Parallel()

			result := testCase.Tag.String()
			assert.Equal(t, testCase.ExpectedResult, result)
		})
	}
}
