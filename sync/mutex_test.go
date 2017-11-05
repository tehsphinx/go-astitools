package astisync_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/go-astitools/sync"
)

func TestRWMutex_IsDeadlocked(t *testing.T) {
	var m = astisync.NewRWMutex("test")
	d, _ := m.IsDeadlocked(time.Millisecond)
	assert.False(t, d)
	m.Lock()
	d, c := m.IsDeadlocked(time.Millisecond)
	assert.True(t, d)
	assert.Contains(t, c, "github.com/tehsphinx/go-astitools/sync/mutex_test.go:15")
}
