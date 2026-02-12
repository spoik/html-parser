package parse_test

import (
	"bufio"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/spoik/html-parser/html"
	"github.com/spoik/html-parser/parse"
	"github.com/spoik/html-parser/stringreader"
)

func TestSuccessfulParseTag(t *testing.T) {
	type testCase struct {
		string                 string
		expectedTag            *html.Tag
		expectedReaderPosition int
	}

	testCases := []testCase{
		{
			"<a href=\"http://www.example.com\">Example</a>",
			&html.Tag{
				Type: "a",
				Attributes: []*html.Attribute{
					{
						Name: "href",
						Value: "http://www.example.com",
					},
				},
			},
			43,
		},
		{
			"<a href class>Example</a>",
			&html.Tag{
				Type: "a",
				Attributes: []*html.Attribute{
					{Name: "href"},
					{Name: "class"},
				},
			},
			15,
		},
		{
			"<a href=\"http://www.example.com\" class=\"btn btn-primary\">Example</a>",
			&html.Tag{
				Type: "a",
				Attributes: []*html.Attribute{
					{
						Name: "href",
						Value: "http://www.example.com",
					},
					{
						Name: "class",
						Value: "btn btn-primary",
					},
				},
			},
			62,
		},
		{
			"<html lang>Example</a>",
			&html.Tag{
				Type: "html",
				Attributes: []*html.Attribute{
					{Name: "lang"},
				},
			},
			15,
		},
		{
			"<html>",
			&html.Tag{
				Type:       "html",
				Attributes: []*html.Attribute(nil)},
			5,
		},
		{
			"<hr data-test/>",
			&html.Tag{
				Type: "hr",
				Attributes: []*html.Attribute{
					{Name: "data-test"},
				},
			},
			14,
		},
		{
			"<hr/>",
			&html.Tag{
				Type:       "hr",
				Attributes: []*html.Attribute(nil)},
			4,
		},
		{
			"<hr",
			&html.Tag{
				Type:       "hr",
				Attributes: []*html.Attribute(nil)},
			2,
		},
		{
			"<hr data",
			&html.Tag{
				Type: "hr",
				Attributes: []*html.Attribute{
					{Name: "data"},
				},
			},
			7,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.string, func(t *testing.T) {
			t.Parallel()

			sr := stringreader.New(testCase.string)
			r := bufio.NewReaderSize(sr, 2)
			r.Discard(1)

			tag, err := parse.ParseTag(r)

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
			r := bufio.NewReaderSize(sr, 2)
			r.Discard(1)

			_, err := parse.ParseTag(r)

			assert.EqualError(t, err, testCase.errorMessage)
		})
	}
}
