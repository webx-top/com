package com

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFixDateTimeString(t *testing.T) {
	result := FixDateTimeString(`2001-1-1 1:1:1`)
	assert.Equal(t, []string{
		`2001-01-01`,
		`01:01:01`,
	}, result)
	result = FixDateTimeString(`2001-1-1 01:01`)
	assert.Equal(t, []string{
		`2001-01-01`,
		`01:01:00`,
	}, result)
	result = FixDateTimeString(`2001-12-31 23:00`)
	assert.Equal(t, []string{`2001-12-31`, `23:00:00`}, result)
	result = FixDateTimeString(`2001-12-31 23:59:59`)
	assert.Equal(t, []string{`2001-12-31`, `23:59:59`}, result)

	// wrong date
	result = FixDateTimeString(`2001-13-1 01:01`)
	assert.Equal(t, []string(nil), result)
	result = FixDateTimeString(`2001-12-60 01:01`)
	assert.Equal(t, []string(nil), result)
	result = FixDateTimeString(`999-12-31 01:01`)
	assert.Equal(t, []string(nil), result)
	result = FixDateTimeString(`10000-12-31 01:01`)
	assert.Equal(t, []string(nil), result)

	// wrong time
	result = FixDateTimeString(`2001-12-31 24:01`)
	assert.Equal(t, []string{`2001-12-31`}, result)
	result = FixDateTimeString(`2001-12-31 23:60`)
	assert.Equal(t, []string{`2001-12-31`}, result)
	result = FixDateTimeString(`2001-12-31 23:59:60`)
	assert.Equal(t, []string{`2001-12-31`}, result)
}
