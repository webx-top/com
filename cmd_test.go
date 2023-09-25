package com

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEnvVar(t *testing.T) {
	os.Setenv(`TESTENV`, `1`)
	v := ParseEnvVar(`ab{$TESTENV}c`)
	assert.Equal(t, `ab1c`, v)
}

func TestParseWindowsEnvVar(t *testing.T) {
	os.Setenv(`TESTENV`, `2`)
	v := ParseWindowsEnvVar(`ab{%TESTENV%}c`)
	assert.Equal(t, `ab2c`, v)
}
