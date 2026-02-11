package stringreader_test

import (
	"testing"

	"github.com/spoik/html-parser/stringreader"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadNext(t *testing.T) {
	sr := stringreader.New("Te")
	assert.Equal(t, sr.Position(), -1)

	byte, err := sr.ReadNext()
	require.NoError(t, err)

	assert.Equal(t, 'T', byte)
	assert.Equal(t, sr.Position(), 0)

	byte, err = sr.ReadNext()
	require.NoError(t, err)

	assert.Equal(t, 'e', byte)
	assert.Equal(t, sr.Position(), 1)
}
