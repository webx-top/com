package json

import (
	"fmt"
	"testing"
)

func TestJSON(t *testing.T) {
	b, err := Marshal(map[string]interface{}{
		`first`: 1,
		`yes`:   true,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
}
