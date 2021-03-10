package com

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFloat(t *testing.T) {
	assert.True(t, IsFloat(`0.123`))
	assert.True(t, IsFloat(`1.0`))
	assert.False(t, IsFloat(`a.0`))
	assert.True(t, IsFloat(`0.0`))
	assert.True(t, IsFloat(`-0.1`))
	assert.Equal(t, -0.1, Float64(`-0.1`))
}
