package parse_test

import (
	"testing"

	"github.com/spoik/html-parser/parse"
	"github.com/spoik/html-parser/stringreader"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSuccessfulParseTag(t *testing.T) {
	type testCase struct {
		string                 string
		expectedTag            *parse.Tag
		expectedReaderPosition int
	}

	testCases := []testCase{
		{
			"<a href=\"https://example.com\">Example</a>",
			&parse.Tag{
				"a",
				[]parse.Attribute{{
					"href",
					"https://example.com",
				}},
			},
			2,
		},
		{
			"<html lang=\"en\">Example</a>",
			&parse.Tag{
				"html",
				[]parse.Attribute{{
					"lang",
					"en",
				}},
			},
			5,
		},
		{
			"<html>",
			&parse.Tag{"html", []parse.Attribute{}},
			5,
		},
		{
			"<hr/>",
			&parse.Tag{"hr", []parse.Attribute{}},
			3,
		},
		{
			"<hr",
			&parse.Tag{"hr", []parse.Attribute{}},
			2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.string, func(t *testing.T) {
			t.Parallel()

			sr := stringreader.New(testCase.string)
			sr.Read(make([]byte, 1))
			tag, err := parse.ParseTag(sr)

			require.NoError(t, err)

			assert.Equal(t, testCase.expectedTag, tag)
			assert.Equal(t, testCase.expectedReaderPosition, sr.Position())
		})
	}
}

func TestFailureParseTag(t *testing.T) {
	type testCase struct {
		string       string
		errorMessage string
	}

	testCases := []testCase{
		{"<>", "Unable to find tag."},
		{"", "Unable to find tag."},
		{" ", "Unable to find tag."},
	}

	for _, testCase := range testCases {
		t.Run(testCase.string, func(t *testing.T) {
			t.Parallel()

			sr := stringreader.New(testCase.string)
			sr.Read(make([]byte, 1))

			_, err := parse.ParseTag(sr)

			assert.EqualError(t, err, testCase.errorMessage)
		})
	}
}
