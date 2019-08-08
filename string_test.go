package com

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTitleCase(t *testing.T) {
	assert.Equal(t, `Webx_Top`, TitleCase(`webx_top`))
	assert.Equal(t, `Webx Top`, TitleCase(`webx top`))
}

func TestSafeBase64(t *testing.T) {
	s, e := SafeBase64Decode(SafeBase64Encode(`webx_top`))
	if e != nil {
		panic(e)
	}
	assert.Equal(t, `webx_top`, s)
}
