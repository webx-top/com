/*
  Copyright 2015 Adrian Stanescu. All rights reserved.
  Use of this source code is governed by the MIT License (MIT) that can be found in the LICENSE file.

  Go program that compares software versions in the x.y.z format
  Usage:
  x := "1"
  y := "1.0.1"
  z := "1.0"
  fmt.Println(VersionCompare(x, y)) // 1 = y
  fmt.Println(VersionCompare(x, z)) // 0 = equal
  fmt.Println(VersionCompare(x, a)) // -1 = x
*/

package com

import (
	"strconv"
	"strings"
)

const (
    VersionCompareGt =-1
    VersionCompareEq = 0
    VersionCompareLt = 1
)

// VersionCompare compare two versions in x.y.z form
// @param  {string} a     version string
// @param  {string} b     version string
// @return {int}          -1 = a is higher, 0 = equal, 1 = b is higher
func VersionCompare(a, b string) (ret int) {
	as := strings.Split(a, ".")
	bs := strings.Split(b, ".")
	al := len(as)
	bl := len(bs)
	loopMax := bl

	if al > bl {
		loopMax = al
	}

	for i := 0; i < loopMax; i++ {
		var x, y string

		if al > i {
			x = as[i]
		}

		if bl > i {
			y = bs[i]
		}

		xi, _ := strconv.Atoi(x)
		yi, _ := strconv.Atoi(y)

		if xi > yi {
			ret = -1
		} else if xi < yi {
			ret = 1
		}

		if ret != 0 {
			break
		}
	}
	return
}
