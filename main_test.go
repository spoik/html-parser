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
		expectedResult string
	}

	testCases := []testCase{
		{"<a href=\"https://example.com\">Example</a>", "a"},
		{"<html lang=\"en\">Example</a>", "html"},
		{"<html>", "html"},
		{"<hr/>", "hr"},
		{"<hr", "hr"},
	}

	for _, testCase := range testCases {
		result, err := main.TagType(&testCase.string, 1)

		require.NoError(t, err)

		assert.Equal(
			t,
			testCase.expectedResult,
			result,
			"Test case: \"%s\"", testCase.string,
		)
	}
}

func TestFailureTagTypes(t *testing.T) {
	type testCase struct {
		string       string
		errorMessage string
	}

	testCases := []testCase{
		{"<>", "Unable to find tag type in \"<>\""},
		{"", "Unable to find tag type in \"\""},
		{" ", "Unable to find tag type in \" \""},
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
