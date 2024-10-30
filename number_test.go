package com

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertNumberChToAr(t *testing.T) {
	v := ConvertNumberChToAr(`一千二百三十二亿四千五百六十七万六千七百八十九`)
	assert.Equal(t, int64(123245676789), v)
}

func TestConvertNumberArToCh(t *testing.T) {
	v := ConvertNumberArToCh(123245676789)
	assert.Equal(t, `一千二百三十二亿四千五百六十七万六千七百八十九`, v)
	assert.Equal(t, `壹仟贰佰叁拾贰亿肆仟伍佰陆拾柒万陆仟柒佰捌拾玖`, UpperChNumber(v))
}
