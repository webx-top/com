package com

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTitleCase(t *testing.T) {
	assert.Equal(t, `Webx_Top`, TitleCase(`webx_top`))
	assert.Equal(t, `Webx Top`, TitleCase(`webx top`))
}
