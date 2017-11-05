package astitemplate_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/go-astitools/template"
)

func TestParseDirectory(t *testing.T) {
	tmpl, err := astitemplate.ParseDirectory("tests", "")
	assert.NoError(t, err)
	assert.Len(t, tmpl.Templates(), 4)
	tmpl, err = astitemplate.ParseDirectory("tests", ".html")
	assert.NoError(t, err)
	assert.Len(t, tmpl.Templates(), 3)
}
