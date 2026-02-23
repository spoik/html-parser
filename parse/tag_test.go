package parse_test

import (
	"bufio"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/spoik/html-parser/html"
	"github.com/spoik/html-parser/parse"
	"github.com/spoik/html-parser/stringreader"
)

func TestSuccessfulParseTag(t *testing.T) {
	type testCase struct {
		string      string
		expectedTag *html.Tag
	}

	testCases := []testCase{
		{
			"<a href=\"http://www.example.com\">Example</a>",
			&html.Tag{
				Type: "a",
				Text: "Example",
				Attributes: []*html.Attribute{
					{
						Name:  "href",
						Value: "http://www.example.com",
					},
				},
			},
		},
		{
			"<a href class>Second example</a>",
			&html.Tag{
				Type: "a",
				Text: "Second example",
				Attributes: []*html.Attribute{
					{Name: "href"},
					{Name: "class"},
				},
			},
		},
		{
			"<a href=\"http://www.example.com\" class=\"btn btn-primary\">Example</a>",
			&html.Tag{
				Type: "a",
				Text: "Example",
				Attributes: []*html.Attribute{
					{
						Name:  "href",
						Value: "http://www.example.com",
					},
					{
						Name:  "class",
						Value: "btn btn-primary",
					},
				},
			},
		},
		{
			"<a class=btn>Example</a>",
			&html.Tag{
				Type: "a",
				Text: "Example",
				Attributes: []*html.Attribute{
					{
						Name:  "class",
						Value: "btn",
					},
				},
			},
		},
		{
			"<a class=\"btn>Example</a>",
			&html.Tag{
				Type: "a",
				Text: "Example",
				Attributes: []*html.Attribute{
					{
						Name:  "class",
						Value: "btn",
					},
				},
			},
		},
		{
			"<a class=btn\">Example</a>",
			&html.Tag{
				Type: "a",
				Text: "Example",
				Attributes: []*html.Attribute{
					{
						Name:  "class",
						Value: "btn",
					},
				},
			},
		},
		{
			"<a class=btn btn-primary>Example</a>",
			&html.Tag{
				Type: "a",
				Text: "Example",
				Attributes: []*html.Attribute{
					{
						Name:  "class",
						Value: "btn",
					},
					{
						Name: "btn-primary",
					},
				},
			},
		},
		{
			"<html lang>Example</a>",
			&html.Tag{
				Type: "html",
				Text: "Example",
				Attributes: []*html.Attribute{
					{Name: "lang"},
				},
			},
		},
		{
			"<html>",
			&html.Tag{Type: "html"},
		},
		{
			"<hr class=bold/>",
			&html.Tag{
				Type: "hr",
				Attributes: []*html.Attribute{
					{
						Name:  "class",
						Value: "bold",
					},
				},
			},
		},
		{
			"<hr data-test/>",
			&html.Tag{
				Type: "hr",
				Attributes: []*html.Attribute{
					{Name: "data-test"},
				},
			},
		},
		{
			"<hr/>",
			&html.Tag{Type: "hr"},
		},
		{
			"<hr  />",
			&html.Tag{Type: "hr"},
		},
		{
			"<hr",
			&html.Tag{Type: "hr"},
		},
		{
			"<hr data",
			&html.Tag{
				Type: "hr",
				Attributes: []*html.Attribute{
					{Name: "data"},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.string, func(t *testing.T) {
			t.Parallel()

			sr := stringreader.New(testCase.string)
			r := bufio.NewReaderSize(sr, 2)
			_, err := r.Discard(1)

			if err != nil {
				t.Fatalf("Error discarding byte: %v", err)
			}

			tag, err := parse.ParseTag(r)

			require.NoError(t, err)

			assert.Equal(t, testCase.expectedTag, tag)
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
			r := bufio.NewReaderSize(sr, 2)
			_, err := r.Discard(1)

			if err != nil && !errors.Is(err, io.EOF) {
				t.Fatalf("Error discarding byte: %v", err)
			}

			_, err = parse.ParseTag(r)

			assert.EqualError(t, err, testCase.errorMessage)
		})
	}
}
