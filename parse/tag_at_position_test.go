package parse_test

import (
	"testing"

	"github.com/spoik/html-parser/parse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSuccessfulAtPosition(t *testing.T) {
	type testCase struct {
		string         string
		expectedResult *parse.TagAtPositionResult
	}

	testCases := []testCase{
		{
			"<a href=\"https://example.com\">Example</a>",
			&parse.TagAtPositionResult{
				Tag: &parse.Tag{"a"},
				EndPos:  1,
			},
		},
		{
			"<html lang=\"en\">Example</a>",
			&parse.TagAtPositionResult{
				Tag: &parse.Tag{"html"},
				EndPos:  4,
			},
		},
		{
			"<html>",
			&parse.TagAtPositionResult{
				Tag: &parse.Tag{"html"},
				EndPos:  4,
			},
		},
		{
			"<hr/>",
			&parse.TagAtPositionResult{
				Tag: &parse.Tag{"hr"},
				EndPos:  2,
			},
		},
		{
			"<hr",
			&parse.TagAtPositionResult{
				Tag: &parse.Tag{"hr"},
				EndPos:  2,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.string, func(t *testing.T) {
			t.Parallel()
			result, err := parse.TagAtPosition(&testCase.string, 1)

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

func TestFailureTagAtPosition(t *testing.T) {
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
		_, err := parse.TagAtPosition(&testCase.string, 1)

		assert.EqualError(
			t,
			err,
			testCase.errorMessage,
			"Test case: \"%s\"", testCase.string,
		)
	}
}
