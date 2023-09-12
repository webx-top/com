package com

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTML2JS(t *testing.T) {
	r := []byte("a\r\n\r\nb\r\n\"quoted\"")
	r = HTML2JS(r)
	assert.Equal(t, `a\n\nb\n\"quoted\"`, string(r))
}

func TestNl2br(t *testing.T) {
	r := "a\r\n\r\nb\r\n"
	r = Nl2br(r)
	expected := `a<br /><br />b<br />`
	assert.Equal(t, expected, r)
	r = "a\n\nb\n"
	r = Nl2br(r)
	assert.Equal(t, expected, r)
}

func TestHTMLDecodeAll(t *testing.T) {
	r := `&amp;#039;`
	r = HTMLDecodeAll(r)
	assert.Equal(t, `'`, r)
}

func TestStripTags(t *testing.T) {
	sources := map[string]string{
		`<script>` + "\r\n" + `alert('js');` + "\r\n" + `</script>`:     ``,
		`<script>` + "\n" + `alert('js');` + "\n" + `</script no="no">`: ``,
		`<script type="text/javascript">alert('js');</script >`:         ``,
		`<style>` + "\r\n" + `.style{}` + "\r\n" + `</style>`:           ``,
		`<style>` + "\n" + `.style{}` + "\n" + `</style no="no">`:       ``,
		`<style type="text/css">.style{}</style >`:                      ``,
		`<a>ha</a>`:                                       `ha`,
		`<a href="#" >ha</a a="b">`:                       `ha`,
		" github  com  / webx-top   /com ":                `github com / webx-top /com`,
		"github\r\n\r\n\r\n[tab]		[/tab]\n\n\n\nwebx-top": "github\n[tab]\t[/tab]\nwebx-top",
	}
	for k, expected := range sources {
		k = StripTags(k)
		assert.Equal(t, expected, k)
	}
}

func TestRemoveXSS(t *testing.T) {
	sources := map[string]string{
		`<sCript>` + "\r\n" + `alert('js');` + "\r\n" + `</script>`:     "<_sCript>\r\nalert('js');\r\n</_script>",
		`<script>` + "\n" + `alert('js');` + "\n" + `</script no="no">`: "<_script>\nalert('js');\n</_script no=\"no\">",
		`<script type="text/javascript">alert('js');</script >`:         "<_script type=\"text/javascript\">alert('js');</_script >",
		`<style>` + "\r\n" + `.style{}` + "\r\n" + `</style>`:           "<_style>\r\n.style{}\r\n</_style>",
		`<style>` + "\n" + `.style{}` + "\n" + `</style no="no">`:       "<_style>\n.style{}\n</_style no=\"no\">",
		`<style type="text/css">.style{}</style >`:                      "<_style type=\"text/css\">.style{}</_style >",
		`<a onload="alert('js')">ha</a>`:                                "<a _onload=\"alert('js')\">ha</a>",
		`<a href="#" sTyle="express()">ha</a a="b">`:                    "<a href=\"#\" _sTyle=\"express()\">ha</a a=\"b\">",
		`<a href="javascript:alert('js')">ha</a a="b">`:                 "<a _href=\"_javascript:alert('js')\">ha</a a=\"b\">",
	}
	for k, expected := range sources {
		k = RemoveXSS(k)
		assert.Equal(t, expected, k)
	}
}
