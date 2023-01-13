package com

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitHostPort(t *testing.T) {
	host, port := SplitHostPort(`[1:2:2:3]:9999`)
	assert.Equal(t, `[1:2:2:3]`, host)
	assert.Equal(t, `9999`, port)

	host, port = SplitHostPort(`example.com:9999`)
	assert.Equal(t, `example.com`, host)
	assert.Equal(t, `9999`, port)

	host, port = SplitHostPort(`127.0.0.1:9999`)
	assert.Equal(t, `127.0.0.1`, host)
	assert.Equal(t, `9999`, port)
}

func TestSplitHostPort2(t *testing.T) {
	host, port := SplitHostPort(`[1:2:2:3]`)
	assert.Equal(t, `[1:2:2:3]`, host)
	assert.Equal(t, ``, port)

	host, port = SplitHostPort(`example.com`)
	assert.Equal(t, `example.com`, host)
	assert.Equal(t, ``, port)

	host, port = SplitHostPort(`127.0.0.1`)
	assert.Equal(t, `127.0.0.1`, host)
	assert.Equal(t, ``, port)
}

func TestRawURLEncode(t *testing.T) {
	rawText := ` +Gopher`
	encoded := RawURLEncode(rawText)
	expected := `%20%2BGopher`
	assert.Equal(t, expected, encoded)
	result, _ := URLDecode(expected)
	assert.Equal(t, rawText, result)
	result, _ = RawURLDecode(expected)
	assert.Equal(t, rawText, result)
}
