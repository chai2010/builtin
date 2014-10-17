// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package builtin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMapSlice(t *testing.T) {
	a := MapSlice([]int{1, 2, 3, 4}, func(val interface{}) interface{} {
		return val.(int) * 2
	})
	if !reflect.DeepEqual(a, []int{2, 4, 6, 8}) {
		t.Fatal("not equal")
	}
}

func TestMapMap(t *testing.T) {
	a := map[int]string{1: "A", 2: "B", 3: "C", 4: "D"}
	b := MapMap(a, func(key, val interface{}) interface{} {
		return fmt.Sprintf("{key(%v),val(%v)}", key.(int), val.(string))
	})
	c := make(map[int]string)

	for k, v := range a {
		c[k] = fmt.Sprintf("{key(%v),val(%v)}", k, v)
	}
	if !reflect.DeepEqual(b, c) {
		t.Fatal("not equal")
	}
}
