package go_rope

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestRopeConcat(t *testing.T) {
	leaf1 := NewRopeLeaf("123456789")
	concat := NewRopeConcat(leaf1, leaf1)
	assert.Equal(t, concat.String(), "123456789123456789")
	assert.Equal(t, concat.Len(), 18)
	assert.Equal(t, concat.Depth(), 1)
	assert.Equal(t, concat.Index(10), "123456789123456789"[10])
}
