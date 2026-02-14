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
		expectedTags []*html.Tag
	}

	testCases := []testCase{
		{
			"<a href=\"https://example.com\">",
			[]*html.Tag{{
				Type: "a",
				Attributes: []*html.Attribute{
					{
						Name: "href",
						Value: "https://example.com",
					},
				},
			}},
		},
		{
			"<html>",
			[]*html.Tag{{
				Type:       "html",
				Attributes: []*html.Attribute(nil),
			}},
		},
		{
			"<hr/>",
			[]*html.Tag{{
				Type:       "hr",
				Attributes: []*html.Attribute(nil),
			}},
		},
		{
			"<hr",
			[]*html.Tag{
				{
					Type:       "hr",
					Attributes: []*html.Attribute(nil),
				},
			},
		},
		{
			"<div><hr>",
			[]*html.Tag{
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
			[]*html.Tag{
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
			[]*html.Tag{
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
			[]*html.Tag{
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
		{
			"<div><p>Paragraph text",
			[]*html.Tag{
				{
					Type:       "div",
					Attributes: []*html.Attribute(nil),
				},
				{
					Type:       "p",
					Text:       "Paragraph text",
					Attributes: []*html.Attribute(nil),
				},
			},
		},
		{
			"<div>Div text<p>Paragraph text",
			[]*html.Tag{
				{
					Type:       "div",
					Text:       "Div text",
					Attributes: []*html.Attribute(nil),
				},
				{
					Type:       "p",
					Text:       "Paragraph text",
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
			"Error parsing HTML: Unable to find tag.",
		},
		{
			"Example",
			"No HTML tags found in \"Example\"",
		},
		{
			"",
			"No HTML tags found in \"\"",
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
