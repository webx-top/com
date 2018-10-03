/*

   Copyright 2016 Wenhui Shen <www.webx.top>

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

*/
package com

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Int64(i interface{}) int64 {
	if v, y := i.(int64); y {
		return v
	}
	if v, y := i.(int32); y {
		return int64(v)
	}
	if v, y := i.(uint32); y {
		return int64(v)
	}
	if v, y := i.(int); y {
		return int64(v)
	}
	if v, y := i.(uint); y {
		return int64(v)
	}
	if v, y := i.(string); y {
		v, _ := strconv.ParseInt(v, 10, 64)
		return v
	}
	in := Str(i)
	if len(in) == 0 {
		return 0
	}
	out, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		log.Printf("string[%s] covert int64 fail. %s", in, err)
		return 0
	}
	return out
}

func Int(i interface{}) int {
	if v, y := i.(int); y {
		return v
	}
	if v, y := i.(string); y {
		v, _ := strconv.Atoi(v)
		return v
	}
	in := Str(i)
	if len(in) == 0 {
		return 0
	}
	out, err := strconv.Atoi(in)
	if err != nil {
		log.Printf("string[%s] covert int fail. %s", in, err)
		return 0
	}
	return out
}

func Int32(i interface{}) int32 {
	if v, y := i.(int32); y {
		return v
	}
	if v, y := i.(string); y {
		v, _ := strconv.ParseInt(v, 10, 32)
		return int32(v)
	}
	in := Str(i)
	if len(in) == 0 {
		return 0
	}
	out, err := strconv.ParseInt(in, 10, 32)
	if err != nil {
		log.Printf("string[%s] covert int32 fail. %s", in, err)
		return 0
	}
	return int32(out)
}

func Uint64(i interface{}) uint64 {
	if v, y := i.(uint64); y {
		return v
	}
	if v, y := i.(string); y {
		v, _ := strconv.ParseUint(v, 10, 64)
		return v
	}
	in := Str(i)
	if len(in) == 0 {
		return 0
	}
	out, err := strconv.ParseUint(in, 10, 64)
	if err != nil {
		log.Printf("string[%s] covert uint64 fail. %s", in, err)
		return 0
	}
	return out
}

func Uint(i interface{}) uint {
	if v, y := i.(uint); y {
		return v
	}
	if v, y := i.(string); y {
		v, _ := strconv.ParseUint(v, 10, 32)
		return uint(v)
	}
	in := Str(i)
	if len(in) == 0 {
		return 0
	}
	out, err := strconv.ParseUint(in, 10, 32)
	if err != nil {
		log.Printf("string[%s] covert uint fail. %s", in, err)
		return 0
	}
	return uint(out)
}

func Uint32(i interface{}) uint32 {
	if v, y := i.(uint32); y {
		return v
	}
	if v, y := i.(string); y {
		v, _ := strconv.ParseUint(v, 10, 32)
		return uint32(v)
	}
	in := Str(i)
	if len(in) == 0 {
		return 0
	}
	out, err := strconv.ParseUint(in, 10, 32)
	if err != nil {
		log.Printf("string[%s] covert uint32 fail. %s", in, err)
		return 0
	}
	return uint32(out)
}

func Float32(i interface{}) float32 {
	if v, y := i.(float32); y {
		return v
	}
	if v, y := i.(int32); y {
		return float32(v)
	}
	if v, y := i.(uint32); y {
		return float32(v)
	}
	if v, y := i.(string); y {
		v, _ := strconv.ParseFloat(v, 32)
		return float32(v)
	}
	in := Str(i)
	if len(in) == 0 {
		return 0
	}
	out, err := strconv.ParseFloat(in, 32)
	if err != nil {
		log.Printf("string[%s] covert float32 fail. %s", in, err)
		return 0
	}
	return float32(out)
}

func Float64(i interface{}) float64 {
	if v, y := i.(float64); y {
		return v
	}
	if v, y := i.(int64); y {
		return float64(v)
	}
	if v, y := i.(uint64); y {
		return float64(v)
	}
	if v, y := i.(float32); y {
		return float64(v)
	}
	if v, y := i.(int32); y {
		return float64(v)
	}
	if v, y := i.(uint32); y {
		return float64(v)
	}
	if v, y := i.(int); y {
		return float64(v)
	}
	if v, y := i.(uint); y {
		return float64(v)
	}
	if v, y := i.(string); y {
		v, _ := strconv.ParseFloat(v, 64)
		return v
	}
	in := Str(i)
	if len(in) == 0 {
		return 0
	}
	out, err := strconv.ParseFloat(in, 64)
	if err != nil {
		log.Printf("string[%s] covert float64 fail. %s", in, err)
		return 0
	}
	return out
}

func Bool(i interface{}) bool {
	if v, y := i.(bool); y {
		return v
	}
	in := Str(i)
	if len(in) == 0 {
		return false
	}
	out, err := strconv.ParseBool(in)
	if err != nil {
		log.Printf("string[%s] covert bool fail. %s", in, err)
		return false
	}
	return out
}

func Str(v interface{}) string {
	if v, y := v.(string); y {
		return v
	}
	return fmt.Sprintf("%v", v)
}

func String(v interface{}) string {
	return Str(v)
}

// SeekRangeNumbers 遍历范围数值，支持设置步进值。格式例如：1-2,2-3:2
func SeekRangeNumbers(expr string, fn func(int) bool) {
	expa := strings.SplitN(expr, ":", 2)
	step := 1
	switch len(expa) {
	case 2:
		if i, _ := strconv.Atoi(strings.TrimSpace(expa[1])); i > 0 {
			step = i
		}
		fallthrough
	case 1:
		for _, exp := range strings.Split(strings.TrimSpace(expa[0]), `,`) {
			exp = strings.TrimSpace(exp)
			if len(exp) == 0 {
				continue
			}
			expb := strings.SplitN(exp, `-`, 2)
			var minN, maxN int
			switch len(expb) {
			case 2:
				maxN, _ = strconv.Atoi(strings.TrimSpace(expb[1]))
				fallthrough
			case 1:
				minN, _ = strconv.Atoi(strings.TrimSpace(expb[0]))
			}
			if maxN == 0 {
				if !fn(minN) {
					return
				}
				continue
			}
			for ; minN <= maxN; minN += step {
				if !fn(minN) {
					return
				}
			}
		}
	}
}
