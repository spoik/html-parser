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
			[]parse.Tag{{"a", []parse.Attribute{}}},
		},
		{
			"<html>",
			[]parse.Tag{{"html", []parse.Attribute{}}},
		},
		{
			"<hr/>",
			[]parse.Tag{{"hr", []parse.Attribute{}}},
		},
		{
			"<hr",
			[]parse.Tag{{"hr", []parse.Attribute{}}},
		},
		{
			"<div><hr>",
			[]parse.Tag{
				{"div", []parse.Attribute{}},
				{"hr", []parse.Attribute{}},
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
			"Unable to find tag.",
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
