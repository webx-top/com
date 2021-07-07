package com

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMkdirAll(t *testing.T) {
	os.RemoveAll(`testdata`)
	dir := `testdata/m/k/d/i`
	//mode := os.ModePerm
	mode := os.FileMode(0766)
	err := MkdirAll(dir, mode)
	assert.NoError(t, err)

	fi, err := os.Stat(dir)
	assert.NoError(t, err)
	assert.Equal(t, mode, fi.Mode().Perm())

	fi, err = os.Stat(filepath.Dir(dir))
	assert.NoError(t, err)
	assert.Equal(t, mode, fi.Mode().Perm())

	fi, err = os.Stat(filepath.Dir(dir))
	assert.NoError(t, err)
	assert.Equal(t, mode, fi.Mode().Perm())

	fi, err = os.Stat(filepath.Dir(dir))
	assert.NoError(t, err)
	assert.Equal(t, mode, fi.Mode().Perm())

	fi, err = os.Stat(filepath.Dir(dir))
	assert.NoError(t, err)
	assert.Equal(t, mode, fi.Mode().Perm())
}
