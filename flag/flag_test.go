package astiflag_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/go-astitools/flag"
)

func TestSubcommand(t *testing.T) {
	os.Args = []string{"bite"}
	assert.Equal(t, "", astiflag.Subcommand())
	os.Args = []string{"bite", "-caca"}
	assert.Equal(t, "", astiflag.Subcommand())
	os.Args = []string{"bite", "caca"}
	assert.Equal(t, "caca", astiflag.Subcommand())
}
