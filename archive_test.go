package com

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIllegalFilePath(t *testing.T) {
	assert.Equal(t, `/a/b/c`, filepath.Dir(`/a/b/c/..a.txt`))
	assert.Equal(t, `/a/b`, filepath.Dir(`/a/b/c/../a.txt`))
	assert.Equal(t, `/abc/a/b/a.txt`, filepath.Join(`/abc`, `/a/b/c/../a.txt`))
	assert.Equal(t, `/abc/a\b\c\..\a.txt`, filepath.Join(`/abc`, `/a\b\c\..\a.txt`))
	assert.Equal(t, `/abc/a/b/c/..a.txt`, filepath.Join(`/abc`, `/a/b/c/..a.txt`))
	assert.Equal(t, `/a/b`, filepath.Join(`/a/b/c`, `..`))
	assert.Equal(t, `/a/b`, filepath.Join(`/a/b/c/..`))
	assert.True(t, IllegalFilePath(`a/b/c/../a.txt`))
	assert.True(t, IllegalFilePath(`a/b/c/..\a.txt`))
	assert.False(t, IllegalFilePath(`a/b/c/..a.txt`))
	assert.True(t, IllegalFilePath(`..`))
	assert.True(t, IllegalFilePath(`/..`))
}

func TestZip(t *testing.T) {
	MkdirAll(`testdata`, os.ModePerm)
	n, err := Zip(`./encoding`, `testdata/test.zip`)
	assert.NoError(t, err)
	assert.NotEqual(t, n, 0)

	err = Unzip(`testdata/test.zip`, `testdata/unarchive`)
	assert.NoError(t, err)
}

func TestTarGz(t *testing.T) {
	MkdirAll(`testdata`, os.ModePerm)
	err := TarGz(`./encoding`, `testdata/test.tar.gz`)
	assert.NoError(t, err)

	_, err = UnTarGz(`testdata/test.tar.gz`, `testdata/unarchive`)
	assert.NoError(t, err)
}
