package com

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceExtract(t *testing.T) {
	parts := []string{`1`, `2`, `3`}
	var n1, n2, n3 string
	SliceExtract(parts, &n1, &n2, &n3)
	assert.Equal(t, parts[0], n1)
	assert.Equal(t, parts[1], n2)
	assert.Equal(t, parts[2], n3)
}
