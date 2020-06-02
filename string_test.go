package com

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTitleCase(t *testing.T) {
	assert.Equal(t, `Webx_Top`, TitleCase(`webx_top`))
	assert.Equal(t, `Webx Top`, TitleCase(`webx top`))
}

func TestSlashes(t *testing.T) {
	assert.Equal(t, `webx\'top\\`, AddSlashes(`webx'top\`))
	assert.Equal(t, `webx'top\`, StripSlashes(`webx\'top\\`))
	assert.Equal(t, `webx\\\'top\\\\`, AddSlashes(`webx\'top\\`))
	assert.Equal(t, `webx\'top\\`, StripSlashes(`webx\\\'top\\\\`))
}

func TestSafeBase64(t *testing.T) {
	s, e := SafeBase64Decode(SafeBase64Encode(`webx_top`))
	if e != nil {
		panic(e)
	}
	assert.Equal(t, `webx_top`, s)
}
