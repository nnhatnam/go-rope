package go_rope

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestRopeLeaf(t *testing.T) {
	s1 := "The quick brown fox jumped over"
	leaf := NewRopeLeaf(s1)
	assert.Equal(t, s1, leaf.String())
	assert.Equal(t, len(s1), leaf.Len())
	assert.Equal(t, leaf.Depth(), 0)
	assert.Equal(t, leaf.Index(5), s1[5])

	leaf2 := leaf.Concat(leaf)
	//leaf remains unchange
	assert.Equal(t, s1, leaf.String())
	assert.Equal(t, len(s1), leaf.Len())
	assert.Equal(t, leaf.Depth(), 0)
	assert.Equal(t, leaf.Index(5), s1[5])

	//leaf2
	assert.Equal(t, s1 + s1, leaf2.String())
	assert.Equal(t, len(s1 + s1), leaf2.Len())
	assert.Equal(t, leaf2.Depth(), 0)
	assert.Equal(t, leaf2.Index(30), (s1 + s1)[30])

}

