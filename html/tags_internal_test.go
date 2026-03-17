package html

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTags(t *testing.T) {
	type testCase struct {
		Name           string
		Tags           []Tag
		ExpectedResult Tags
	}

	testCases := []testCase{
		{
			Name:           "Empty tag slice.",
			Tags:           []Tag{},
			ExpectedResult: Tags{tags: []Tag{}},
		},
		{
			Name: "Populated tag slice.",
			Tags: []Tag{
				{Type: "a"},
				{Type: "p"},
			},
			ExpectedResult: Tags{
				tags: []Tag{
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
			assert.Condition(
				t,
				func() bool { return testCase.ExpectedResult.Equal(result) },
				fmt.Sprintf(
					"Tags are not equal: Expected \"%s\" got \"%s\".",
					testCase.ExpectedResult,
					result,
				),
			)
		})
	}
}

func TestEqual(t *testing.T) {
	type testCase struct {
		Name  string
		TagsA Tags
		TagsB Tags
	}

	testCases := []testCase{
		{
			Name:  "Both html.Tags have a no tags.",
			TagsA: Tags{},
			TagsB: Tags{},
		},
		{
			Name:  "Both html.Tags have a nil tags.",
			TagsA: Tags{tags: nil},
			TagsB: Tags{tags: nil},
		},
		{
			Name:  "Both html.Tags have a nil TagIndex.",
			TagsA: Tags{tags: []Tag{{Type: "a"}}},
			TagsB: Tags{tags: []Tag{{Type: "a"}}},
		},
		{
			Name: "One html.Tags has a nil TagIndex and the other doesn't.",
			TagsA: Tags{
				tags:     []Tag{{Type: "a"}},
				tagIndex: &TagIndex{},
			},
			TagsB: Tags{
				tags: []Tag{{Type: "a"}},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			assert.Condition(
				t,
				func() bool {
					return testCase.TagsA.Equal(testCase.TagsB)
				},
			)
		})
	}
}

func TestNotEqual(t *testing.T) {
	type testCase struct {
		Name  string
		TagsA Tags
		TagsB Tags
	}

	testCases := []testCase{
		{
			Name:  "Tags have different tags.",
			TagsA: Tags{tags: []Tag{{Type: "a"}}},
			TagsB: Tags{tags: []Tag{{Type: "img"}}},
		},
		{
			Name:  "Tags have different amount tags.",
			TagsA: Tags{tags: []Tag{{Type: "a"}}},
			TagsB: Tags{tags: []Tag{{Type: "a"}, {Type: "img"}}},
		},
		{
			Name:  "Tags have different amount tags.",
			TagsA: Tags{tags: []Tag{{Type: "a"}, {Type: "img"}}},
			TagsB: Tags{tags: []Tag{{Type: "a"}}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			assert.Condition(
				t,
				func() bool {
					return !testCase.TagsA.Equal(testCase.TagsB)
				},
			)
		})
	}
}
