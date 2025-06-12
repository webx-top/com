package com

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberTrim(t *testing.T) {
	s := `2.123987666`
	i := Float2int(s)
	assert.Equal(t, 2, i)
	r := NumberTrim(s, 5)
	assert.Equal(t, `2.12398`, r)
	r = NumberTrim(s, 1, ``)
	assert.Equal(t, `2.1`, r)
}

func TestNumberTrimZero(t *testing.T) {
	s := fmt.Sprintf("%.7f", 2.123)
	assert.Equal(t, `2.1230000`, s)
	assert.Equal(t, 2.123, Float64(s))

	r := NumberTrimZero(s)
	assert.Equal(t, `2.123`, r)

	r = FormatBytes(12344566, 2, true)
	assert.Equal(t, `11.77MB`, r)
}

func TestToFloat64(t *testing.T) {
	v := `100000000000000163.0100`
	n := Float64(v)
	assert.Equal(t, 100000000000000163.01, n)
}
