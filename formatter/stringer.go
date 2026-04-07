package formatter

import (
	"fmt"
)

type stringer struct {
	raw any
	enc Encoder
}

func (s *stringer) SetEncoder(enc Encoder) {
	s.enc = enc
}

func (s stringer) String() string {
	return s.enc.Encode(s.raw)
}

func (s stringer) Raw() any {
	return s.raw
}

func AsStringer(v any, encoder ...Encoder) fmt.Stringer {
	var enc Encoder
	if len(encoder) > 0 {
		enc = encoder[0]
	}
	if enc == nil {
		enc = JSONEncoder
	}
	return stringer{raw: v, enc: enc}
}
