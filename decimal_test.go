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
	t.Logf(`convert %s => %f`, v, n)
	assert.Equal(t, 100000000000000163.01, n)
	//assert.Equal(t, v, String(n)) // 100000000000000160
	v2 := `95000000000000005.0100`
	n2 := Float64(v2)
	t.Logf(`convert %s => %f`, v2, n2)
	assert.Equal(t, 95000000000000005.0100, n2)
	//assert.Equal(t, v2, String(n2)) // 95000000000000000
	vY := `100000000000163.0100`
	nY := Float64(vY)
	t.Logf(`convert %s => %f`, vY, nY)
	assert.Equal(t, 100000000000163.01, nY) // 100000000000163.015625

	// dec, err := decimal.NewFromString(vY)
	// assert.NoError(t, err)
	// nY = dec.InexactFloat64()
	// t.Logf(`convert %s => %f`, vY, nY)
	// assert.Equal(t, 100000000000163.01, nY)
}
