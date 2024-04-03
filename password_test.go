package com

import (
	"crypto/sha1"
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
	dk := PBKDF2Key([]byte("some password"), []byte(salt), 4096, 32, sha1.New)
	fmt.Println(`PBKDF2:`, string(dk), len(dk))

	hashedPassword, err := BCryptMakePassword(`github.com/webx-top/com`)
	assert.NoError(t, err)
	fmt.Println(`BCrypt:`, string(hashedPassword), len(hashedPassword)) // $2a$10$7eI1PPyMbvY5E6g6IeOh.OunLNMrguV/tL.mK9HZIUf//iBZ49nW6
	err = BCryptCheckPassword(string(hashedPassword), `github.com/webx-top/com`)
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

func TestAbsURL(t *testing.T) {
	pageURL := AbsURL(`https://www.coscms.com/system/download/index`, `../download2/index`)
	assert.Equal(t, `https://www.coscms.com/system/download2/index`, pageURL)

	pageURL = AbsURL(`https://www.coscms.com/system/download/index`, `../../system2/download2/index`)
	assert.Equal(t, `https://www.coscms.com/system2/download2/index`, pageURL)

	pageURL = AbsURL(`https://www.coscms.com/system/download/index`, `/payment/index/index`)
	assert.Equal(t, `https://www.coscms.com/payment/index/index`, pageURL)

	pageURL = AbsURL(`https://www.coscms.com/system/download/index`, `./payment/index/index`)
	assert.Equal(t, `https://www.coscms.com/system/download/payment/index/index`, pageURL)

	fmt.Println(`SelfDir:`, SelfDir())
	fmt.Println(`SelfPath:`, SelfPath())
}
