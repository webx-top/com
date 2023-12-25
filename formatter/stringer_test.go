package formatter

import "testing"

func TestStringer(t *testing.T) {
	st := AsStringer(map[string]interface{}{
		`A`: 100,
		`B`: true,
	})
	t.Logf(`JSONEncoder: %s`, st)

	st = AsStringer(map[string]interface{}{
		`A`: 100,
		`B`: true,
	}, PrettyEncoder)
	t.Logf(`PrettyEncoder: %s`, st)
}
