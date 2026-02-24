package html_test

import (
	"testing"

	"github.com/spoik/html-parser/html"
	"github.com/stretchr/testify/assert"
)

func TestNewTags(t *testing.T) {
	type testCase struct {
		Name           string
		Tags           []*html.Tag
		ExpectedResult *html.Tags
	}

	testCases := []testCase{
		{
			Name:           "Empty tag slice.",
			Tags:           []*html.Tag{},
			ExpectedResult: nil,
		},
		{
			Name: "Populated tag slice.",
			Tags: []*html.Tag{
				{Type: "a"},
				{Type: "p"},
			},
			ExpectedResult: &html.Tags{
				Tags: []*html.Tag{
					{Type: "a"},
					{Type: "p"},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			result := html.NewTags(testCase.Tags)
			assert.Equal(t, testCase.ExpectedResult, result)
		})
	}
}

func TestFind(t *testing.T) {
	type testCase struct {
		Name           string
		Tags           html.Tags
		TagType        string
		ExpectedResult []*html.Tag
	}

	testCases := []testCase{
		{
			Name:           "With no tags.",
			Tags:           html.Tags{},
			TagType:        "a",
			ExpectedResult: []*html.Tag{},
		},
		{
			Name: "With one matching tag.",
			Tags: *html.NewTags([]*html.Tag{
				{Type: "a"},
				{Type: "p"},
			}),
			TagType: "a",
			ExpectedResult: []*html.Tag{
				{Type: "a"},
			},
		},
		{
			Name: "With one matching tag, returns a copy with all attributes except Tags.",
			Tags: *html.NewTags([]*html.Tag{
				{
					Type: "a",
					Text: "Text",
					Attributes: html.NewAttributes([]*html.Attribute{
						{
							Name:  "Name",
							Value: "Value",
						},
					}),
					Tags: html.NewTags([]*html.Tag{
						{Type: "p"},
					}),
				},
			}),
			TagType: "a",
			ExpectedResult: []*html.Tag{
				{
					Type: "a",
					Text: "Text",
					Attributes: html.NewAttributes([]*html.Attribute{
						{
							Name:  "Name",
							Value: "Value",
						},
					}),
				},
			},
		},
		{
			Name: "With multiple matching tag.",
			Tags: *html.NewTags([]*html.Tag{
				{Type: "a"},
				{Type: "a"},
				{Type: "p"},
			}),
			TagType: "a",
			ExpectedResult: []*html.Tag{
				{Type: "a"},
				{Type: "a"},
			},
		},
		{
			Name: "With deep matching tags.",
			Tags: *html.NewTags([]*html.Tag{
				{
					Type: "a",
					Tags: html.NewTags([]*html.Tag{
						{
							Type: "a",
							Tags: html.NewTags([]*html.Tag{
								{Type: "a"},
							}),
						},
					}),
				},
				{Type: "p"},
			}),
			TagType: "a",
			ExpectedResult: []*html.Tag{
				{Type: "a"},
				{Type: "a"},
				{Type: "a"},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			result := testCase.Tags.Find(testCase.TagType)

			assert.Equal(t, testCase.ExpectedResult, result)
		})
	}
}

func TestFindDoesNotModifyOriginalTags(t *testing.T) {
	tags := html.NewTags([]*html.Tag{
		{
			Type: "a",
			Tags: html.NewTags([]*html.Tag{
				{Type: "p"},
				{Type: "img"},
			}),
		},
	})

	assert.Equal(t, 2, len(tags.Tags[0].Tags.Tags))
	result := tags.Find("a")
	assert.Equal(t, 2, len(tags.Tags[0].Tags.Tags))

	expectedTags := []*html.Tag{
		{Type: "a"},
	}
	assert.Equal(t, expectedTags, result)
}
