package main_test

import (
	"testing"

	"github.com/spoik/html-parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSuccessfulTagTypes(t *testing.T) {
	type testCase struct {
		string         string
		expectedResult *main.TagTypeResult
	}

	testCases := []testCase{
		{
			"<a href=\"https://example.com\">Example</a>",
			&main.TagTypeResult{
				TagType: "a",
				EndPos:  1,
			},
		},
		{
			"<html lang=\"en\">Example</a>",
			&main.TagTypeResult{
				TagType: "html",
				EndPos:  4,
			},
		},
		{
			"<html>",
			&main.TagTypeResult{
				TagType: "html",
				EndPos:  4,
			},
		},
		{
			"<hr/>",
			&main.TagTypeResult{
				TagType: "hr",
				EndPos:  2,
			},
		},
		{
			"<hr",
			&main.TagTypeResult{
				TagType: "hr",
				EndPos:  2,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.string, func(t *testing.T) {
			t.Parallel()
			result, err := main.TagType(&testCase.string, 1)

			require.NoError(t, err)

			assert.Equal(
				t,
				testCase.expectedResult,
				result,
				"Test case: \"%s\"", testCase.string,
			)
		})
	}
}

func TestFailureTagTypes(t *testing.T) {
	type testCase struct {
		string       string
		errorMessage string
	}

	testCases := []testCase{
		{"<>", "Unable to find tag type in \"<>\" starting at position 1."},
		{"", "Unable to find tag type in \"\" starting at position 1."},
		{" ", "Unable to find tag type in \" \" starting at position 1."},
	}

	for _, testCase := range testCases {
		_, err := main.TagType(&testCase.string, 1)

		assert.EqualError(
			t,
			err,
			testCase.errorMessage,
			"Test case: \"%s\"", testCase.string,
		)
	}
}
