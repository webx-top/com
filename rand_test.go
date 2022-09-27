package com

import (
	"fmt"
	"testing"
)

func TestRandStr(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println(`RandomString:`, RandomString(32))
	}
	for i := 0; i < 5; i++ {
		fmt.Println(`RandStr:`, RandStr(32))
	}
	for i := 0; i < 5; i++ {
		fmt.Println(`RandomASCII:`, RandomASCII(32))
	}
	for i := 0; i < 5; i++ {
		fmt.Println(`RandomAlphabetic:`, RandomAlphabetic(32))
	}
	for i := 0; i < 5; i++ {
		fmt.Println(`RandomAlphanumeric:`, RandomAlphanumeric(32))
	}
	for i := 0; i < 5; i++ {
		fmt.Println(`RandomNumeric:`, RandomNumeric(32))
	}
}
