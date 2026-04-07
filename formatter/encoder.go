package formatter

import (
	"github.com/admpub/pp/ppnocolor"
	"github.com/webx-top/com"
)

type Encoder interface {
	Encode(any) string
}

type EncoderFunc func(any) string

func (f EncoderFunc) Encode(v any) string {
	return f(v)
}

var JSONEncoder = jsonEncode{}
var PrettyEncoder = prettyEncode{}

type jsonEncode struct {
}

func (s jsonEncode) Encode(v any) string {
	return com.Dump(v, false)
}

type prettyEncode struct {
}

func (s prettyEncode) Encode(v any) string {
	return ppnocolor.Sprint(v)
}
