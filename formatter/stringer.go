package formatter

import (
	"fmt"

	"github.com/admpub/pp/ppnocolor"
	"github.com/webx-top/com"
)

type Encoder interface {
	Encode(interface{}) string
}

type stringer struct {
	raw interface{}
	enc Encoder
}

func (s *stringer) SetEncoder(enc Encoder) {
	s.enc = enc
}

func (s stringer) String() string {
	return s.enc.Encode(s.raw)
}

func (s stringer) Raw() interface{} {
	return s.raw
}

var JSONEncoder = jsonEncode{}
var PrettyEncoder = prettyEncode{}

type jsonEncode struct {
}

func (s jsonEncode) Encode(v interface{}) string {
	return com.Dump(v, false)
}

type prettyEncode struct {
}

func (s prettyEncode) Encode(v interface{}) string {
	return ppnocolor.Sprint(v)
}

func AsStringer(v interface{}, encoder ...Encoder) fmt.Stringer {
	var enc Encoder
	if len(encoder) > 0 {
		enc = encoder[0]
	}
	if enc == nil {
		enc = JSONEncoder
	}
	return stringer{raw: v, enc: enc}
}
