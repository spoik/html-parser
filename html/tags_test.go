package html_test

import (
	"testing"

	"github.com/spoik/html-parser/html"
	"github.com/stretchr/testify/assert"
)

func createTagsWithIndex(tags []*html.Tag) *html.Tags {
	index := &html.TagIndex{}
	index.AddAll(tags)
	return html.TagsWithIndex(tags, index)
}

func TestFind(t *testing.T) {
	type testCase struct {
		Name           string
		Tags           *html.Tags
		TagType        string
		ExpectedResult []*html.Tag
	}

	testCases := []testCase{
		{
			Name:           "With no tags.",
			Tags:           createTagsWithIndex([]*html.Tag{}),
			TagType:        "a",
			ExpectedResult: []*html.Tag{},
		},
		{
			Name: "With one matching tag.",
			Tags: createTagsWithIndex([]*html.Tag{
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
			Tags: createTagsWithIndex([]*html.Tag{
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
			Tags: createTagsWithIndex([]*html.Tag{
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
			Tags: createTagsWithIndex([]*html.Tag{
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
	tags := createTagsWithIndex([]*html.Tag{
		{
			Type: "a",
			Tags: html.NewTags([]*html.Tag{
				{Type: "p"},
				{Type: "img"},
			}),
		},
	})

	tag, err := tags.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, 2, tag.Tags.Length())

	result := tags.Find("a")

	tag, err = tags.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, 2, tag.Tags.Length())

	expectedTags := []*html.Tag{
		{Type: "a"},
	}
	assert.Equal(t, expectedTags, result)
}

func TestGet(t *testing.T) {
	type testCase struct {
		Name           string
		Tags           html.Tags
		Index          int
		ExpectedResult *html.Tag
		ExpectedError  error
	}

	testCases := []testCase{
		{
			Name:           "Empty Tags.",
			Tags:           html.Tags{},
			Index:          0,
			ExpectedResult: nil,
			ExpectedError:  html.NoTagAtIndex,
		},
		{
			Name:           "Out of bounds index.",
			Tags:           *html.NewTags([]*html.Tag{{Type: "a"}}),
			Index:          4,
			ExpectedResult: nil,
			ExpectedError:  html.NoTagAtIndex,
		},
		{
			Name:           "Valid index.",
			Tags:           *html.NewTags([]*html.Tag{{Type: "a"}}),
			Index:          0,
			ExpectedResult: &html.Tag{Type: "a"},
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

func TestLength(t *testing.T) {
	type testCase struct {
		Name           string
		Tags           *html.Tags
		ExpectedResult int
	}

	testCases := []testCase{
		{
			Name:           "Tags with nil tags slice.",
			Tags:           &html.Tags{},
			ExpectedResult: 0,
		},
		{
			Name:           "Tags with empty tags slice.",
			Tags:           html.NewTags([]*html.Tag{}),
			ExpectedResult: 0,
		},
		{
			Name: "Tags with tags slice.",
			Tags: html.NewTags(
				[]*html.Tag{
					{Type: "a"},
					{Type: "p"},
					{Type: "img"},
				},
			),
			ExpectedResult: 3,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			result := testCase.Tags.Length()
			assert.Equal(t, testCase.ExpectedResult, result)
		})
	}
}

func TestFullLength(t *testing.T) {
	type testCase struct {
		Name           string
		Tags           *html.Tags
		ExpectedResult int
	}

	testCases := []testCase{
		{
			Name:           "Empty tags.",
			Tags:           &html.Tags{},
			ExpectedResult: 0,
		},
		{
			Name: "One tag.",
			Tags: html.NewTags(
				[]*html.Tag{
					{Type: "a"},
				},
			),
			ExpectedResult: 1,
		},
		{
			Name: "Two tag.",
			Tags: html.NewTags(
				[]*html.Tag{
					{Type: "a"},
					{Type: "p"},
				},
			),
			ExpectedResult: 2,
		},
		{
			Name: "Deeply nested tags.",
			Tags: html.NewTags(
				[]*html.Tag{
					{
						Type: "a",
						Tags: html.NewTags(
							[]*html.Tag{
								{Type: "img"},
								{Type: "p"},
							},
						),
					},
					{
						Type: "p",
						Tags: html.NewTags(
							[]*html.Tag{
								{Type: "p"},
							},
						),
					},
				},
			),
			ExpectedResult: 5,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			result := testCase.Tags.FullLength()
			assert.Equal(t, testCase.ExpectedResult, result)
		})
	}
}
