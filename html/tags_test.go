package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTags(t *testing.T) {
	type testCase struct {
		Name           string
		Tags           []*Tag
		ExpectedResult *Tags
	}

	testCases := []testCase{
		{
			Name:           "Empty tag slice.",
			Tags:           []*Tag{},
			ExpectedResult: nil,
		},
		{
			Name: "Populated tag slice.",
			Tags: []*Tag{
				{Type: "a"},
				{Type: "p"},
			},
			ExpectedResult: &Tags{
				tags: []*Tag{
					{Type: "a"},
					{Type: "p"},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			result := NewTags(testCase.Tags)
			assert.Equal(t, testCase.ExpectedResult, result)
		})
	}
}

func TestFind(t *testing.T) {
	type testCase struct {
		Name           string
		Tags           Tags
		TagType        string
		ExpectedResult []*Tag
	}

	testCases := []testCase{
		{
			Name:           "With no tags.",
			Tags:           Tags{},
			TagType:        "a",
			ExpectedResult: []*Tag{},
		},
		{
			Name: "With one matching tag.",
			Tags: *NewTags([]*Tag{
				{Type: "a"},
				{Type: "p"},
			}),
			TagType: "a",
			ExpectedResult: []*Tag{
				{Type: "a"},
			},
		},
		{
			Name: "With one matching tag, returns a copy with all attributes except Tags.",
			Tags: *NewTags([]*Tag{
				{
					Type: "a",
					Text: "Text",
					Attributes: NewAttributes([]*Attribute{
						{
							Name:  "Name",
							Value: "Value",
						},
					}),
					Tags: NewTags([]*Tag{
						{Type: "p"},
					}),
				},
			}),
			TagType: "a",
			ExpectedResult: []*Tag{
				{
					Type: "a",
					Text: "Text",
					Attributes: NewAttributes([]*Attribute{
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
			Tags: *NewTags([]*Tag{
				{Type: "a"},
				{Type: "a"},
				{Type: "p"},
			}),
			TagType: "a",
			ExpectedResult: []*Tag{
				{Type: "a"},
				{Type: "a"},
			},
		},
		{
			Name: "With deep matching tags.",
			Tags: *NewTags([]*Tag{
				{
					Type: "a",
					Tags: NewTags([]*Tag{
						{
							Type: "a",
							Tags: NewTags([]*Tag{
								{Type: "a"},
							}),
						},
					}),
				},
				{Type: "p"},
			}),
			TagType: "a",
			ExpectedResult: []*Tag{
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
	tags := NewTags([]*Tag{
		{
			Type: "a",
			Tags: NewTags([]*Tag{
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

	expectedTags := []*Tag{
		{Type: "a"},
	}
	assert.Equal(t, expectedTags, result)
}

func TestGet(t *testing.T) {
	type testCase struct {
		Name           string
		Tags           Tags
		Index          int
		ExpectedResult *Tag
		ExpectedError  error
	}

	testCases := []testCase{
		{
			Name:           "Empty Tags.",
			Tags:           Tags{},
			Index:          0,
			ExpectedResult: nil,
			ExpectedError:  NoTagAtIndex,
		},
		{
			Name:           "Out of bounds index.",
			Tags:           *NewTags([]*Tag{{Type: "a"}}),
			Index:          4,
			ExpectedResult: nil,
			ExpectedError:  NoTagAtIndex,
		},
		{
			Name:           "Valid index.",
			Tags:           *NewTags([]*Tag{{Type: "a"}}),
			Index:          0,
			ExpectedResult: &Tag{Type: "a"},
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
			assert.ErrorIs(t, err, NoTagAtIndex)
		})
	}
}
