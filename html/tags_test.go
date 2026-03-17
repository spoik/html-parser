package html_test

import (
	"testing"

	"github.com/spoik/html-parser/html"
	"github.com/stretchr/testify/assert"
)

func tagsWithIndex(tags []html.Tag) html.Tags {
	index := &html.TagIndex{}
	index.AddAll(tags)

	return html.NewTags(tags, html.WithIndex(index))
}

func emptyTags() html.Tags {
	return html.NewTags([]html.Tag{})
}

func TestFind(t *testing.T) {
	type testCase struct {
		Name           string
		Tags           html.Tags
		TagType        string
		ExpectedResult []html.Tag
	}

	testCases := []testCase{
		{
			Name:           "With no tags.",
			Tags:           tagsWithIndex([]html.Tag{}),
			TagType:        "a",
			ExpectedResult: []html.Tag{},
		},
		{
			Name: "With one matching tag.",
			Tags: tagsWithIndex([]html.Tag{
				{Type: "a"},
				{Type: "p"},
			}),
			TagType: "a",
			ExpectedResult: []html.Tag{
				{Type: "a"},
			},
		},
		{
			Name: "With one matching tag, returns a copy with all attributes except Tags.",
			Tags: tagsWithIndex([]html.Tag{
				{
					Type: "a",
					Text: "Text",
					Attributes: html.NewAttributes([]html.Attribute{
						{
							Name:  "Name",
							Value: "Value",
						},
					}),
					Tags: html.NewTags([]html.Tag{
						{Type: "p"},
					}),
				},
			}),
			TagType: "a",
			ExpectedResult: []html.Tag{
				{
					Type: "a",
					Text: "Text",
					Attributes: html.NewAttributes([]html.Attribute{
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
			Tags: tagsWithIndex([]html.Tag{
				{Type: "a"},
				{Type: "a"},
				{Type: "p"},
			}),
			TagType: "a",
			ExpectedResult: []html.Tag{
				{Type: "a"},
				{Type: "a"},
			},
		},
		{
			Name: "With deep matching tags.",
			Tags: tagsWithIndex([]html.Tag{
				{
					Type: "a",
					Tags: html.NewTags([]html.Tag{
						{
							Type: "a",
							Tags: html.NewTags([]html.Tag{
								{Type: "a"},
							}),
						},
					}),
				},
				{Type: "p"},
			}),
			TagType: "a",
			ExpectedResult: []html.Tag{
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
	tags := tagsWithIndex([]html.Tag{
		{
			Type: "a",
			Tags: html.NewTags([]html.Tag{
				{Type: "p"},
				{Type: "img"},
			}),
		},
	})

	tag, err := tags.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, 2, tag.Tags.Len())

	result := tags.Find("a")

	tag, err = tags.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, 2, tag.Tags.Len())

	expectedTags := []html.Tag{
		{Type: "a"},
	}
	assert.Equal(t, expectedTags, result)
}

func TestGet(t *testing.T) {
	type testCase struct {
		Name           string
		Tags           html.Tags
		Index          int
		ExpectedResult html.Tag
		ExpectedError  error
	}

	testCases := []testCase{
		{
			Name:           "Empty Tags.",
			Tags:           html.Tags{},
			Index:          0,
			ExpectedResult: html.Tag{},
			ExpectedError:  html.NoTagAtIndex,
		},
		{
			Name:           "Out of bounds index.",
			Tags:           html.NewTags([]html.Tag{{Type: "a"}}),
			Index:          4,
			ExpectedResult: html.Tag{},
			ExpectedError:  html.NoTagAtIndex,
		},
		{
			Name:           "Valid index.",
			Tags:           html.NewTags([]html.Tag{{Type: "a"}}),
			Index:          0,
			ExpectedResult: html.Tag{Type: "a"},
			ExpectedError:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			result, err := testCase.Tags.Get(testCase.Index)

			if testCase.ExpectedError == nil {
				assert.NoError(t, err)
				assert.Equal(t, testCase.ExpectedResult, result)
				return
			}

			assert.Error(t, err)
			assert.ErrorIs(t, err, html.NoTagAtIndex)
		})
	}
}

func TestLen(t *testing.T) {
	type testCase struct {
		Name           string
		Tags           html.Tags
		ExpectedResult int
	}

	testCases := []testCase{
		{
			Name:           "Tags with nil tags slice.",
			Tags:           html.Tags{},
			ExpectedResult: 0,
		},
		{
			Name:           "Tags with empty tags slice.",
			Tags:           emptyTags(),
			ExpectedResult: 0,
		},
		{
			Name: "Tags with tags slice.",
			Tags: html.NewTags([]html.Tag{
				{Type: "a"},
				{Type: "p"},
				{Type: "img"},
			}),
			ExpectedResult: 3,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			result := testCase.Tags.Len()
			assert.Equal(t, testCase.ExpectedResult, result)
		})
	}
}

func TestFullLen(t *testing.T) {
	type testCase struct {
		Name           string
		Tags           html.Tags
		ExpectedResult int
	}

	testCases := []testCase{
		{
			Name:           "Empty tags.",
			Tags:           html.Tags{},
			ExpectedResult: 0,
		},
		{
			Name: "One tag.",
			Tags: html.NewTags([]html.Tag{
				{Type: "a"},
			}),
			ExpectedResult: 1,
		},
		{
			Name: "Two tag.",
			Tags: html.NewTags([]html.Tag{
				{Type: "a"},
				{Type: "p"},
			}),
			ExpectedResult: 2,
		},
		{
			Name: "Deeply nested tags.",
			Tags: html.NewTags([]html.Tag{
				{
					Type: "a",
					Tags: html.NewTags([]html.Tag{
						{Type: "img"},
						{Type: "p"},
					},
					),
				},
				{
					Type: "p",
					Tags: html.NewTags([]html.Tag{
						{Type: "p"},
					}),
				},
			}),
			ExpectedResult: 5,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			result := testCase.Tags.FullLen()
			assert.Equal(t, testCase.ExpectedResult, result)
		})
	}
}

func TestTagsString(t *testing.T) {
	type testCase struct {
		Tags           html.Tags
		ExpectedResult string
	}

	testCases := []testCase{
		{
			Tags: html.NewTags([]html.Tag{
				{
					Type: "a",
					Text: "Anchor text",
				},
				{
					Type: "span",
					Text: "Span text",
				},
			}),
			ExpectedResult: "<a>Anchor text</a><span>Span text</span>",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.ExpectedResult, func(t *testing.T) {
			t.Parallel()

			result := testCase.Tags.String()

			assert.Equal(t, testCase.ExpectedResult, result)
		})
	}
}

func TestIterator(t *testing.T) {
	type testCase struct {
		Tags           html.Tags
		ExpectedResult []html.Tag
	}

	testCases := []testCase{
		{
			Tags:           html.NewTags([]html.Tag{}),
			ExpectedResult: []html.Tag{},
		},
		{
			Tags: html.NewTags([]html.Tag{
				{
					Type: "a",
					Text: "text",
				},
			}),
			ExpectedResult: []html.Tag{
				{
					Type: "a",
					Text: "text",
				},
			},
		},
		{
			Tags: html.NewTags([]html.Tag{
				{
					Type: "div",
					Text: "text",
				},
				{
					Type: "a",
					Text: "text",
				},
			}),
			ExpectedResult: []html.Tag{
				{
					Type: "div",
					Text: "text",
				},
				{
					Type: "a",
					Text: "text",
				},
			},
		},
		{
			Tags: html.NewTags([]html.Tag{
				{
					Type: "div",
					Text: "text",
					Tags: html.NewTags([]html.Tag{
						{
							Type: "a",
							Text: "text",
						},
					}),
				},
			}),
			ExpectedResult: []html.Tag{
				{
					Type: "div",
					Text: "text",
					Tags: html.NewTags([]html.Tag{
						{
							Type: "a",
							Text: "text",
						},
					}),
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Tags.String(), func(t *testing.T) {
			t.Parallel()

			results := make([]html.Tag, testCase.Tags.Len())

			for i, tag := range testCase.Tags.AllTags() {
				results[i] = tag
			}

			assert.Equal(
				t,
				testCase.ExpectedResult,
				results,
			)
		})
	}
}
