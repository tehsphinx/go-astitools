package astibyte_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/go-astitools/byte"
)

func TestPadLeft(t *testing.T) {
	assert.Equal(t, []byte("test"), astibyte.PadLeft([]byte("test"), ' ', 2))
	assert.Equal(t, []byte("test"), astibyte.PadLeft([]byte("test"), ' ', 4))
	assert.Equal(t, []byte("  test"), astibyte.PadLeft([]byte("test"), ' ', 6))
}

func TestPadRight(t *testing.T) {
	assert.Equal(t, []byte("test"), astibyte.PadRight([]byte("test"), ' ', 2))
	assert.Equal(t, []byte("test"), astibyte.PadRight([]byte("test"), ' ', 4))
	assert.Equal(t, []byte("test  "), astibyte.PadRight([]byte("test"), ' ', 6))
}
