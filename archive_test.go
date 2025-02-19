package com

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZip(t *testing.T) {
	MkdirAll(`testdata`, os.ModePerm)
	n, err := Zip(`./encoding`, `testdata/test.zip`)
	assert.NoError(t, err)
	assert.NotEqual(t, n, 0)

	err = Unzip(`testdata/test.zip`, `testdata/unarchive`)
	assert.NoError(t, err)
}
