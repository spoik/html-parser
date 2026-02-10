package main_test

import (
	"testing"

	"github.com/spoik/html-parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSuccessfulParseHtml(t *testing.T) {
	type testCase struct {
		html        string
		expectedTag []main.Tag
	}

	testCases := []testCase{
		{
			"<a href=\"https://example.com\">",
			[]main.Tag{{Type: "a"}},
		},
		{
			"<html lang=\"en\">",
			[]main.Tag{{Type: "html"}},
		},
		{
			"<html>",
			[]main.Tag{{Type: "html"}},
		},
		{
			"<hr/>",
			[]main.Tag{{Type: "hr"}},
		},
		{
			"<hr",
			[]main.Tag{{Type: "hr"}},
		},
		{
			"<div><hr>",
			[]main.Tag{
				{Type: "div"},
				{Type: "hr"},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.html, func(t *testing.T) {
			t.Parallel()
			tag, err := main.ParseHtml(&testCase.html)

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
			"Unable to find tag type in \"<>\" starting at position 1.",
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
			tag, err := main.ParseHtml(&testCase.html)

			require.Error(t, err)
			assert.Nil(t, tag)
			assert.EqualError(t, err, testCase.errorMessage)
		})
	}
}
