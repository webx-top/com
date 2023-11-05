package com

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceExtractCallback(t *testing.T) {
	parts := []string{`1`, `2`, `3`}
	var n1N, n2N, n3N int
	SliceExtractCallback(parts, func(v string) int {
		return Int(v)
	}, &n1N, &n2N, &n3N)
	assert.Equal(t, 1, n1N)
	assert.Equal(t, 2, n2N)
	assert.Equal(t, 3, n3N)
}

func TestReverseSortIndex(t *testing.T) {
	parts := []string{`1`, `2`, `3`}
	ReverseSortIndex(parts)
	assert.Equal(t, `3`, parts[0])
	assert.Equal(t, `2`, parts[1])
	assert.Equal(t, `1`, parts[2])
}
