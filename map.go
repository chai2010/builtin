// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package builtin

import (
	"fmt"
	"reflect"
)

func MapSlice(slice interface{}, fn func(a interface{}) interface{}) interface{} {
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		panic(fmt.Sprintf("MapSlice called with non-slice value of type %T", slice))
	}
	out := reflect.MakeSlice(reflect.TypeOf(slice), val.Len(), val.Cap())
	for i := 0; i < val.Len(); i++ {
		out.Index(i).Set(
			reflect.ValueOf(fn(val.Index(i).Interface())),
		)
	}
	return out.Interface()
}

func MapMap(m interface{}, fn func(key, val interface{}) interface{}) interface{} {
	val := reflect.ValueOf(m)
	if val.Kind() != reflect.Map {
		panic(fmt.Sprintf("MapMap called with non-map value of type %T", m))
	}
	out := reflect.MakeMap(reflect.TypeOf(m))
	for _, key := range val.MapKeys() {
		out.SetMapIndex(
			key, reflect.ValueOf(fn(key.Interface(), val.MapIndex(key).Interface())),
		)
	}
	return out.Interface()
}
