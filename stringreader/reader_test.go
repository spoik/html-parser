package stringreader_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/spoik/html-parser/stringreader"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestImplementsIoReaderInterface(t *testing.T) {
	assert.Implements(t, (*io.Reader)(nil), new(stringreader.StringReader))
}

func TestRead(t *testing.T) {
	string := "Test"

	type ExpectedResult struct {
		BytesRead      int
		Bytes          []byte
		Error          error
		ReaderPosition int
	}

	type TestCase struct {
		ReadLen         int
		ExpectedResults []ExpectedResult
	}

	testCases := []TestCase{
		{
			ReadLen: 1,
			ExpectedResults: []ExpectedResult{
				{
					BytesRead:      1,
					Bytes:          []byte{'T'},
					ReaderPosition: 0,
				},
				{
					BytesRead:      1,
					Bytes:          []byte{'e'},
					ReaderPosition: 1,
				},
				{
					BytesRead:      1,
					Bytes:          []byte{'s'},
					ReaderPosition: 2,
				},
				{
					BytesRead:      1,
					Bytes:          []byte{'t'},
					ReaderPosition: 3,
				},
				{
					BytesRead:      0,
					Bytes:          []byte{'t'},
					Error:          io.EOF,
					ReaderPosition: 3,
				},
			},
		},
		{
			ReadLen: 2,
			ExpectedResults: []ExpectedResult{
				{
					BytesRead:      2,
					Bytes:          []byte{'T', 'e'},
					ReaderPosition: 1,
				},
				{
					BytesRead:      2,
					Bytes:          []byte{'s', 't'},
					ReaderPosition: 3,
				},
				{
					BytesRead:      0,
					Bytes:          []byte{'s', 't'},
					Error:          io.EOF,
					ReaderPosition: 3,
				},
			},
		},
		{
			ReadLen: 3,
			ExpectedResults: []ExpectedResult{
				{
					BytesRead:      3,
					Bytes:          []byte{'T', 'e', 's'},
					ReaderPosition: 2,
				},
				{
					BytesRead:      1,
					Bytes:          []byte{'t', 'e', 's'},
					Error:          io.EOF,
					ReaderPosition: 3,
				},
				{
					BytesRead:      0,
					Bytes:          []byte{'t', 'e', 's'},
					Error:          io.EOF,
					ReaderPosition: 3,
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Read length %d", testCase.ReadLen), func(t *testing.T) {
			t.Parallel()

			bytes := make([]byte, testCase.ReadLen)
			sr := stringreader.New(string)

			for _, result := range testCase.ExpectedResults {
				numRead, err := sr.Read(bytes)

				assert.Equal(t, result.Bytes, bytes)
				assert.Equal(t, result.BytesRead, numRead)
				assert.Equal(t, result.ReaderPosition, sr.Position())

				if result.Error == nil {
					require.NoError(t, err)
				} else {
					require.Error(t, err)
					assert.ErrorIs(t, err, io.EOF)
				}
			}
		})
	}
}
