package com

import (
	"crypto/sha1"
	"fmt"
	"testing"
)

func TestMakePassword(t *testing.T) {
	hashed := MakePassword(`github.com/webx-top/com`, ``)
	fmt.Println(`hashed`, hashed, len(hashed))
	if CheckPassword(`github.com/webx-top/com`, hashed, ``) == false {
		t.Errorf(`The passwords do not match`)
	}
	salt := Salt()
	fmt.Println(`salt:`, salt)
	hashed = MakePassword(`github.com/webx-top/com`, salt, 2, 4, 6, 8, 13, 19, 32)
	fmt.Println(`hashed`, hashed, len(hashed))
	if CheckPassword(`github.com/webx-top/com`, hashed, salt, 2, 4, 6, 8, 13, 19, 32) == false {
		t.Errorf(`The passwords do not match`)
	}
	salt = Salt()
	hashed = MakePassword(`github.com/webx-top/com`, salt, 13)
	fmt.Println(`hashed`, hashed, len(hashed))
	if CheckPassword(`github.com/webx-top/com`, hashed, salt, 13) == false {
		t.Errorf(`The passwords do not match`)
	}
	salt = `github.com/webx-top/com`
	dk := PBKDF2Key([]byte("some password"), []byte(salt), 4096, 32, sha1.New)
	fmt.Println(`PBKDF2:`, string(dk))
}
