package parse_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/spoik/html-parser/html"
	"github.com/spoik/html-parser/parse"
)

func TestSuccessfulParseHtml(t *testing.T) {
	type testCase struct {
		html         string
		expectedTags *html.Tags
	}

	testCases := []testCase{
		{
			"<a href=\"https://example.com\">",
			html.NewTags(html.NewTagsOpts{
				Tags: []*html.Tag{{
					Type: "a",
					Attributes: html.NewAttributes([]*html.Attribute{
						{
							Name:  "href",
							Value: "https://example.com",
						},
					}),
					Tags: html.EmptyTags(),
				}},
			},
			),
		},
		{
			"<html>",
			html.NewTags(html.NewTagsOpts{
				Tags: []*html.Tag{{
					Type: "html",
					Tags: html.EmptyTags(),
				}},
			}),
		},
		{
			"<hr/>",
			html.NewTags(html.NewTagsOpts{
				Tags: []*html.Tag{{
					Type: "hr",
					Tags: html.EmptyTags(),
				}},
			}),
		},
		{
			"<hr",
			html.NewTags(html.NewTagsOpts{
				Tags: []*html.Tag{{
					Type: "hr",
					Tags: html.EmptyTags(),
				}},
			}),
		},
		{
			"<img src><hr data>",
			html.NewTags(html.NewTagsOpts{
				Tags: []*html.Tag{
					{
						Type:       "img",
						Attributes: html.NewAttributes([]*html.Attribute{{Name: "src"}}),
						Tags: html.NewTags(html.NewTagsOpts{
							Tags: []*html.Tag{
								{
									Type: "hr",
									Attributes: html.NewAttributes([]*html.Attribute{
										{Name: "data"},
									}),
									Tags: html.EmptyTags(),
								},
							},
						}),
					},
				},
			}),
		},
		{
			"<img src/><hr data/>",
			html.NewTags(html.NewTagsOpts{
				Tags: []*html.Tag{
					{
						Type:       "img",
						Attributes: html.NewAttributes([]*html.Attribute{{Name: "src"}}),
						Tags:       html.EmptyTags(),
					},
					{
						Type:       "hr",
						Attributes: html.NewAttributes([]*html.Attribute{{Name: "data"}}),
						Tags:       html.EmptyTags(),
					},
				},
			}),
		},
		{
			"<img/><hr/>",
			html.NewTags(html.NewTagsOpts{
				Tags: []*html.Tag{
					{
						Type: "img",
						Tags: html.EmptyTags(),
					},
					{
						Type: "hr",
						Tags: html.EmptyTags(),
					},
				},
			}),
		},
		{
			"<div><p>Paragraph text",
			html.NewTags(html.NewTagsOpts{
				Tags: []*html.Tag{
					{
						Type: "div",
						Tags: html.NewTags(html.NewTagsOpts{
							Tags: []*html.Tag{
								{
									Type: "p",
									Text: "Paragraph text",
									Tags: html.EmptyTags(),
								},
							},
						}),
					},
				},
			}),
		},
		{
			"<div>Div text<p>Paragraph text</p>",
			html.NewTags(html.NewTagsOpts{
				Tags: []*html.Tag{
					{
						Type: "div",
						Text: "Div text",
						Tags: html.NewTags(html.NewTagsOpts{
							Tags: []*html.Tag{
								{
									Type: "p",
									Text: "Paragraph text",
									Tags: html.EmptyTags(),
								},
							},
						}),
					},
				},
			}),
		},
		{
			"<div>Div text<p>Paragraph text</p></div>",
			html.NewTags(html.NewTagsOpts{
				Tags: []*html.Tag{
					{
						Type: "div",
						Text: "Div text",
						Tags: html.NewTags(html.NewTagsOpts{
							Tags: []*html.Tag{
								{
									Type: "p",
									Text: "Paragraph text",
									Tags: html.EmptyTags(),
								},
							},
						}),
					},
				},
			}),
		},
		{
			"<div><p><a></a></p></div>",
			html.NewTags(html.NewTagsOpts{
				Tags: []*html.Tag{
					{
						Type: "div",
						Tags: html.NewTags(html.NewTagsOpts{
							Tags: []*html.Tag{
								{
									Type: "p",
									Tags: html.NewTags(html.NewTagsOpts{
										Tags: []*html.Tag{
											{
												Type: "a",
												Tags: html.EmptyTags(),
											},
										},
									}),
								},
							},
						}),
					},
				},
			}),
		},
		{
			"<div><p></p><a></a></div>",
			html.NewTags(html.NewTagsOpts{
				Tags: []*html.Tag{
					{
						Type: "div",
						Tags: html.NewTags(html.NewTagsOpts{
							Tags: []*html.Tag{
								{
									Type: "p",
									Tags: html.EmptyTags(),
								},
								{
									Type: "a",
									Tags: html.EmptyTags(),
								},
							},
						}),
					},
				},
			}),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.html, func(t *testing.T) {
			t.Parallel()
			tags, err := parse.ParseHtml(&testCase.html)

			require.NoError(t, err)
			assert.Condition(t, func() bool {
				return testCase.expectedTags.Equal(tags)
			})
		})
	}
}

func TestParseHtmlWithTagsWithAnOptionalClosingTag(t *testing.T) {
	type testCase struct {
		html         string
		expectedTags *html.Tags
	}

	tagTypes := []string{"br", "hr", "img", "input", "link", "meta", "source"}
	testCases := make([]testCase, len(tagTypes)*2)

	for i, tagType := range tagTypes {
		index := i * 2

		testCases[index] = testCase{
			html: fmt.Sprintf("<section><%s></section>", tagType),
			expectedTags: html.NewTags(html.NewTagsOpts{
				Tags: []*html.Tag{
					{
						Type: "section",
						Tags: html.NewTags(html.NewTagsOpts{
							Tags: []*html.Tag{{
								Type: tagType,
								Tags: html.EmptyTags(),
							}},
						}),
					},
				},
			}),
		}

		testCases[index+1] = testCase{
			html: fmt.Sprintf("<section><%s></%s></section>", tagType, tagType),
			expectedTags: html.NewTags(html.NewTagsOpts{
				Tags: []*html.Tag{
					{
						Type: "section",
						Tags: html.NewTags(html.NewTagsOpts{
							Tags: []*html.Tag{{
								Type: tagType,
								Tags: html.EmptyTags(),
							}},
						}),
					},
				},
			}),
		}
	}

	for _, testCase := range testCases {
		t.Run(testCase.html, func(t *testing.T) {
			t.Parallel()
			tags, err := parse.ParseHtml(&testCase.html)

			require.NoError(t, err)
			assert.Condition(t, func() bool {
				return testCase.expectedTags.Equal(tags)
			})
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
			"Error parsing HTML: End tag type does not matching opening tag type. Expected end tag type \"a\", but got \"p\".",
		},
		{
			"<a><longtagname></a>",
			"Error parsing HTML: End tag missing for tag type \"longtagname\". The end of HTML reached before finding the end tag.",
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
