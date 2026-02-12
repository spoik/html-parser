package parse_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/spoik/html-parser/html"
	"github.com/spoik/html-parser/parse"
)

func TestSuccessfulParseHtml(t *testing.T) {
	type testCase struct {
		html         string
		expectedTags *[]*html.Tag
	}

	testCases := []testCase{
		{
			"<a href>",
			&[]*html.Tag{{
				Type: "a",
				Attributes: []*html.Attribute{
					{Name: "href"},
				},
			}},
		},
		{
			"<html>",
			&[]*html.Tag{{
				Type:       "html",
				Attributes: []*html.Attribute(nil),
			}},
		},
		{
			"<hr/>",
			&[]*html.Tag{{
				Type:       "hr",
				Attributes: []*html.Attribute(nil),
			}},
		},
		{
			"<hr",
			&[]*html.Tag{
				{
					Type:       "hr",
					Attributes: []*html.Attribute(nil),
				},
			},
		},
		{
			"<div><hr>",
			&[]*html.Tag{
				{
					Type:       "div",
					Attributes: []*html.Attribute(nil),
				},
				{
					Type:       "hr",
					Attributes: []*html.Attribute(nil),
				},
			},
		},
		{
			"<img src><hr data>",
			&[]*html.Tag{
				{
					Type:       "img",
					Attributes: []*html.Attribute{{Name: "src"}},
				},
				{
					Type:       "hr",
					Attributes: []*html.Attribute{{Name: "data"}},
				},
			},
		},
		{
			"<img src/><hr data/>",
			&[]*html.Tag{
				{
					Type:       "img",
					Attributes: []*html.Attribute{{Name: "src"}},
				},
				{
					Type:       "hr",
					Attributes: []*html.Attribute{{Name: "data"}},
				},
			},
		},
		{
			"<img/><hr/>",
			&[]*html.Tag{
				{
					Type:       "img",
					Attributes: []*html.Attribute(nil),
				},
				{
					Type:       "hr",
					Attributes: []*html.Attribute(nil),
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.html, func(t *testing.T) {
			t.Parallel()
			tags, err := parse.ParseHtml(&testCase.html)

			require.NoError(t, err)
			assert.Equal(t, testCase.expectedTags, tags)
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
