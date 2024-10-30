package com

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubstr(t *testing.T) {
	assert.Equal(t, `web`, Substr(`webx_top`, ``, 3))
	assert.Equal(t, `历史书`, Substr(`历史书籍`, ``, 3))
}

func TestTitleCase(t *testing.T) {
	assert.Equal(t, `Webx_Top`, TitleCase(`webx_top`))
	assert.Equal(t, `Webx Top`, TitleCase(`webx top`))
}

func TestSlashes(t *testing.T) {
	assert.Equal(t, `webx\'top\\`, AddSlashes(`webx'top\`))
	assert.Equal(t, `webx'top\`, StripSlashes(`webx\'top\\`))
	assert.Equal(t, `webx\\\'top\\\\`, AddSlashes(`webx\'top\\`))
	assert.Equal(t, `webx\'top\\`, StripSlashes(`webx\\\'top\\\\`))
	assert.Equal(t, `webx'top'`, StripSlashes(`webx\'top\'`))
	s := `webx
eee
	www	www2
`
	actual := AddRSlashes(s)
	assert.Equal(t, `webx\neee\n\twww\twww2\n`, actual)
	fmt.Println(actual)
}

func TestSafeBase64(t *testing.T) {
	s, e := SafeBase64Decode(SafeBase64Encode(`webx_top`))
	if e != nil {
		panic(e)
	}
	assert.Equal(t, `webx_top`, s)
}

func TestTitle(t *testing.T) {
	v := Title(`nickName`)
	assert.Equal(t, `NickName`, v)
	v = Title(`nick_name`)
	assert.Equal(t, `Nick_name`, v)
}

func TestContainsWord(t *testing.T) {
	v := ContainsWord(`application/x-apple-diskimage`, `image`)
	assert.False(t, v)
	v = ContainsWord(`image/jpg`, `image`)
	assert.True(t, v)
	v = ContainsWord(`imagejpg`, `image`)
	assert.False(t, v)
	v = ContainsWord(`abc/image`, `image`)
	assert.True(t, v)
	v = ContainsWord(`abc/image/jpg`, `image`)
	assert.True(t, v)
	v = ContainsWord(`abc/imagejpg`, `image`)
	assert.False(t, v)
}

func TestMultipleBytesText(t *testing.T) {
	s := `。，（）-1！@234567890abc１２３４５６７８９ａｂｃ`
	v := MBToSBText(s)
	assert.Equal(t, `｡,()-1!@234567890abc123456789abc`, v)
	assert.Equal(t, `。，（）－１！＠２３４５６７８９０ａｂｃ１２３４５６７８９ａｂｃ`, SBToMBText(v))
}
