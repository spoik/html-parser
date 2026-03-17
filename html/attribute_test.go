package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAttribute(t *testing.T) {
	type testCase struct {
		Name       string
		Attributes []Attribute
		Expected   Attributes
	}

	testCases := []testCase{
		{
			Name: "One attribute",
			Attributes: []Attribute{
				{
					Name:  "class",
					Value: "primary",
				},
			},
			Expected: Attributes{
				attributes: map[string]Attribute{
					"class": {
						Name:  "class",
						Value: "primary",
					},
				},
			},
		},
		{
			Name: "Two attributes",
			Attributes: []Attribute{
				{
					Name:  "class",
					Value: "primary",
				},
				{
					Name:  "id",
					Value: "name",
				},
			},
			Expected: Attributes{
				attributes: map[string]Attribute{
					"class": {
						Name:  "class",
						Value: "primary",
					},
					"id": {
						Name:  "id",
						Value: "name",
					},
				},
			},
		},
		{
			Name: "Duplicate attributes",
			Attributes: []Attribute{
				{
					Name:  "class",
					Value: "primary",
				},
				{
					Name:  "class",
					Value: "secondary",
				},
			},
			Expected: Attributes{
				attributes: map[string]Attribute{
					"class": {
						Name:  "class",
						Value: "primary",
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			result := NewAttributes(testCase.Attributes)

			assert.Equal(t, testCase.Expected, result)
		})
	}
}

func TestAttribute(t *testing.T) {
	type testCase struct {
		Name               string
		Attributes         Attributes
		AttrName           string
		ExpectedAttriburte Attribute
		ExpectedOk         bool
	}

	testCases := []testCase{
		{
			Name: "Attribute is present in the attirbutes.",
			Attributes: NewAttributes([]Attribute{
				{Name: "id", Value: "profile"},
			}),
			AttrName:           "id",
			ExpectedAttriburte: Attribute{Name: "id", Value: "profile"},
			ExpectedOk: true,
		},
		{
			Name: "Attribute is not present in the attirbutes.",
			Attributes: NewAttributes([]Attribute{
				{Name: "id", Value: "profile"},
			}),
			AttrName:           "class",
			ExpectedAttriburte: Attribute{},
			ExpectedOk: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()

			result, ok := testCase.Attributes.Attribute(testCase.AttrName)

			assert.Equal(t, testCase.ExpectedAttriburte, result)
			assert.Equal(t, testCase.ExpectedOk, ok)
		})
	}
}
