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
				Attributes: html.NewAttributes([]*html.Attribute{
					{
						Name:  "href",
						Value: "https://example.com",
					},
				}),
			}},
		},
		{
			"<html>",
			[]*html.Tag{{Type: "html"}},
		},
		{
			"<hr/>",
			[]*html.Tag{{Type: "hr"}},
		},
		{
			"<hr",
			[]*html.Tag{{Type: "hr"}},
		},
		{
			"<div><hr>",
			[]*html.Tag{
				{
					Type: "div",
					Tags: []*html.Tag{{Type: "hr"}},
				},
			},
		},
		{
			"<img src><hr data>",
			[]*html.Tag{
				{
					Type:       "img",
					Attributes: html.NewAttributes([]*html.Attribute{{Name: "src"}}),
					Tags: []*html.Tag{
						{
							Type: "hr",
							Attributes: html.NewAttributes([]*html.Attribute{
								{Name: "data"},
							}),
						},
					},
				},
			},
		},
		{
			"<img src/><hr data/>",
			[]*html.Tag{
				{
					Type:       "img",
					Attributes: html.NewAttributes([]*html.Attribute{{Name: "src"}}),
				},
				{
					Type:       "hr",
					Attributes: html.NewAttributes([]*html.Attribute{{Name: "data"}}),
				},
			},
		},
		{
			"<img/><hr/>",
			[]*html.Tag{
				{Type: "img"},
				{Type: "hr"},
			},
		},
		{
			"<div><p>Paragraph text",
			[]*html.Tag{
				{
					Type: "div",
					Tags: []*html.Tag{
						{
							Type: "p",
							Text: "Paragraph text",
						},
					},
				},
			},
		},
		{
			"<div>Div text<p>Paragraph text</p>",
			[]*html.Tag{
				{
					Type: "div",
					Text: "Div text",
					Tags: []*html.Tag{
						{
							Type: "p",
							Text: "Paragraph text",
						},
					},
				},
			},
		},
		{
			"<div>Div text<p>Paragraph text</p></div>",
			[]*html.Tag{
				{
					Type: "div",
					Text: "Div text",
					Tags: []*html.Tag{
						{
							Type: "p",
							Text: "Paragraph text",
						},
					},
				},
			},
		},
		{
			"<div><p><a></a></p></div>",
			[]*html.Tag{
				{
					Type: "div",
					Tags: []*html.Tag{
						{
							Type: "p",
							Tags: []*html.Tag{
								{
									Type: "a",
								},
							},
						},
					},
				},
			},
		},
		{
			"<div><p></p><a></a></div>",
			[]*html.Tag{
				{
					Type: "div",
					Tags: []*html.Tag{
						{Type: "p"},
						{Type: "a"},
					},
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
		{
			"<a></p>",
			"Error parsing HTML: End tag type does not matching opening tag type. Expected end tag type a, but got p",
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
