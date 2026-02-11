package parse_test

import (
	"testing"

	"github.com/spoik/html-parser/parse"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSuccessfulParseHtml(t *testing.T) {
	type testCase struct {
		html        string
		expectedTag []parse.Tag
	}

	testCases := []testCase{
		{
			"<a href=\"https://example.com\">",
			[]parse.Tag{{Type: "a"}},
		},
		{
			"<html lang=\"en\">",
			[]parse.Tag{{Type: "html"}},
		},
		{
			"<html>",
			[]parse.Tag{{Type: "html"}},
		},
		{
			"<hr/>",
			[]parse.Tag{{Type: "hr"}},
		},
		{
			"<hr",
			[]parse.Tag{{Type: "hr"}},
		},
		{
			"<div><hr>",
			[]parse.Tag{
				{Type: "div"},
				{Type: "hr"},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.html, func(t *testing.T) {
			t.Parallel()
			tag, err := parse.ParseHtml(&testCase.html)

			require.NoError(t, err)
			assert.Equal(t, testCase.expectedTag, *tag)
		})
	}
}

func TestUnsuccessfulParseHtml(t *testing.T) {
	type testCase struct {
		html         string
		errorMessage string
	}

	testCases := []testCase{
		{
			"<>",
			"Unable to find tag type.",
		},
		{
			"Example",
			"No HTML found in \"Example\"",
		},
		{
			"",
			"No HTML found in \"\"",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.html, func(t *testing.T) {
			t.Parallel()
			tag, err := parse.ParseHtml(&testCase.html)

			require.Error(t, err)
			assert.Nil(t, tag)
			assert.EqualError(t, err, testCase.errorMessage)
		})
	}
}
