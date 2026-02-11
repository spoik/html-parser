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
		expectedTag            *parse.Tag
		expectedReaderPosition int
	}

	testCases := []testCase{
		{
			"<a href=\"https://example.com\">Example</a>",
			&parse.Tag{"a"},
			2,
		},
		{
			"<html lang=\"en\">Example</a>",
			&parse.Tag{"html"},
			5,
		},
		{
			"<html>",
			&parse.Tag{"html"},
			5,
		},
		{
			"<hr/>",
			&parse.Tag{"hr"},
			3,
		},
		{
			"<hr",
			&parse.Tag{"hr"},
			2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.string, func(t *testing.T) {
			t.Parallel()

			sr := stringreader.New(testCase.string)
			sr.Read(make([]byte, 1))
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
		{"<>", "Unable to find tag type."},
		{"", "Unable to find tag type."},
		{" ", "Unable to find tag type."},
	}

	for _, testCase := range testCases {
		t.Run(testCase.string, func(t *testing.T) {
			t.Parallel()

			sr := stringreader.New(testCase.string)
			sr.Read(make([]byte, 1))

			_, err := parse.TagAtPosition(sr)

			assert.EqualError(t, err, testCase.errorMessage)
		})
	}
}
