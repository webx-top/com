package com

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFloat(t *testing.T) {
	assert.True(t, IsFloat(`0.123`))
	assert.True(t, IsFloat(`1.0`))
	assert.False(t, IsFloat(`a.0`))
	assert.True(t, IsFloat(`0.0`))
	assert.True(t, IsFloat(`-0.1`))
	assert.Equal(t, -0.1, Float64(`-0.1`))
}

func TestReplaceByMatchedIndex(t *testing.T) {
	content := `<body>
    {{Include "header"}}
	<div id="wrapper">
		{{Include "sidebar"}}
		<div class="body">
		</div>
	</div>
</body>`
	matches := regexp.MustCompile(`\{\{Include "([^"]+)"\}\}`).FindAllStringSubmatchIndex(content, -1)
	var replaced string
	fn := ReplaceByMatchedIndex(content, matches, &replaced)
	for k, v := range matches {
		var tmplFile, passObject string
		GetMatchedByIndex(content, v, nil, &tmplFile, &passObject)
		if k == 0 {
			assert.Equal(t, `header`, tmplFile)
			assert.Equal(t, ``, passObject)
		} else {
			assert.Equal(t, `sidebar`, tmplFile)
			assert.Equal(t, ``, passObject)
		}
		fn(k, v, `{P}`)
	}
	expected := `<body>
    {P}
	<div id="wrapper">
		{P}
		<div class="body">
		</div>
	</div>
</body>`
	assert.Equal(t, expected, replaced)
	var replaced2 string
	fn2 := ReplaceByMatchedIndex(content, matches, &replaced2)
	for k, v := range matches {
		fn2(k, v)
	}
	assert.Equal(t, content, replaced2)
}

func TestReplaceByMatchedIndex2(t *testing.T) {
	content := `{{Include "sub"}}`
	matches := regexp.MustCompile(`\{\{Include "([^"]+)"\}\}`).FindAllStringSubmatchIndex(content, -1)
	Dump(matches)
	var replaced string
	fn := ReplaceByMatchedIndex(content, matches, &replaced)
	for k, v := range matches {
		var tmplFile, passObject string
		GetMatchedByIndex(content, v, nil, &tmplFile, &passObject)
		assert.Equal(t, `sub`, tmplFile)
		assert.Equal(t, ``, passObject)
		fn(k, v, `{P}`)
	}
	expected := `{P}`
	assert.Equal(t, expected, replaced)
	var replaced2 string
	fn2 := ReplaceByMatchedIndex(content, matches, &replaced2)
	for k, v := range matches {
		fn2(k, v)
	}
	assert.Equal(t, content, replaced2)
}
