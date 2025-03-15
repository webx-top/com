package com

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
	hashedPassword, err := BCryptMakePassword(`github.com/webx-top/com`)
	assert.NoError(t, err)
	fmt.Println(`BCrypt:`, string(hashedPassword), len(hashedPassword)) // $2a$10$7eI1PPyMbvY5E6g6IeOh.OunLNMrguV/tL.mK9HZIUf//iBZ49nW6
	err = BCryptCheckPassword(string(hashedPassword), `github.com/webx-top/com`)
	assert.NoError(t, err)

	hashed, err = SCryptMakePassword(`github.com/webx-top/com`)
	assert.NoError(t, err)
	fmt.Println(`SCrypt:`, hashed, len(hashed)) // 32768$8$1$JlhjCxGNkNubajnj_wQRSg$b8b2e5864a2724b635028a19a58f83efb6f876b4713f7631bf5aa663c6269f2b 97
	err = SCryptCheckPassword(hashed, `github.com/webx-top/com`)
	assert.NoError(t, err)

	hashed, err = Argon2MakePasswordShortly(`github.com/webx-top/com`, Salt())
	assert.NoError(t, err)
	fmt.Println(`Argon2Shortly:`, hashed, len(hashed))
	err = Argon2CheckPassword(hashed, `github.com/webx-top/com`)
	assert.NoError(t, err)
	hashed, err = Argon2MakePasswordShortly(`github.com/webx-top/com`)
	assert.NoError(t, err)
	fmt.Println(`Argon2Shortly:`, hashed, len(hashed))
	err = Argon2CheckPassword(hashed, `github.com/webx-top/com`)
	assert.NoError(t, err)

	hashed, err = Argon2MakePassword(`github.com/webx-top/com`, Salt())
	assert.NoError(t, err)
	fmt.Println(`Argon2:`, hashed, len(hashed))
	err = Argon2CheckPassword(hashed, `github.com/webx-top/com`)
	assert.NoError(t, err)
	hashed, err = Argon2MakePassword(`github.com/webx-top/com`)
	assert.NoError(t, err)
	fmt.Println(`Argon2:`, hashed, len(hashed))
	err = Argon2CheckPassword(hashed, `github.com/webx-top/com`)
	assert.NoError(t, err)
}
