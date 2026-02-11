package parse_test

import (
	"testing"

	"github.com/spoik/html-parser/parse"
	"github.com/spoik/html-parser/stringreader"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSuccessfulAtPosition(t *testing.T) {
	type testCase struct {
		string                 string
		expectedTag            parse.Tag
		expectedReaderPosition int
	}

	testCases := []testCase{
		{
			"<a href=\"https://example.com\">Example</a>",
			parse.Tag{"a"},
			1,
		},
		{
			"<html lang=\"en\">Example</a>",
			parse.Tag{"html"},
			4,
		},
		{
			"<html>",
			parse.Tag{"html"},
			4,
		},
		{
			"<hr/>",
			parse.Tag{"hr"},
			2,
		},
		{
			"<hr",
			parse.Tag{"hr"},
			2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.string, func(t *testing.T) {
			t.Parallel()

			sr := stringreader.New(testCase.string)
			tag, err := parse.TagAtPosition(sr)

			require.NoError(t, err)

			assert.Equal(t, testCase.expectedTag, tag)
			assert.Equal(t, testCase.expectedReaderPosition, sr.Position())
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
		sr := stringreader.New(testCase.string)
		_, err := parse.TagAtPosition(sr)

		assert.EqualError(t, err, testCase.errorMessage)
	}
}
