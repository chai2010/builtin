// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package builtin

import (
	"reflect"
)

func Copy(to, from interface{}) (err error) {
	var (
		isSlice   bool
		fromType  reflect.Type
		toType    reflect.Type
		elemCount int
	)

	fromValue := reflect.ValueOf(from)
	toValue := reflect.ValueOf(to)
	fromElem := reflect.Indirect(fromValue)
	toElem := reflect.Indirect(toValue)

	if toElem.Kind() == reflect.Slice {
		isSlice = true
		if fromElem.Kind() == reflect.Slice {
			fromType = fromElem.Type().Elem()
			elemCount = fromElem.Len()
		} else {
			fromType = fromElem.Type()
			elemCount = 1
		}
		toType = toElem.Type().Elem()
	} else {
		fromType = fromElem.Type()
		toType = toElem.Type()
		elemCount = 1
	}

	for e := 0; e < elemCount; e++ {
		var dest, source reflect.Value
		if isSlice {
			if fromElem.Kind() == reflect.Slice {
				source = fromElem.Index(e)
			} else {
				source = fromElem
			}
		} else {
			source = fromElem
		}

		if isSlice {
			dest = reflect.New(toType).Elem()
		} else {
			dest = toElem
		}

		for i := 0; i < fromType.NumField(); i++ {
			field := fromType.Field(i)
			if !field.Anonymous {
				name := field.Name
				fromField := source.FieldByName(name)
				toField := dest.FieldByName(name)
				toMethod := dest.Addr().MethodByName(name)

				if fromField.IsValid() && toField.IsValid() {
					toField.Set(fromField)
				}
				if fromField.IsValid() && toMethod.IsValid() {
					toMethod.Call([]reflect.Value{fromField})
				}
			}
		}

		for i := 0; i < dest.NumField(); i++ {
			field := toType.Field(i)
			if !field.Anonymous {
				name := field.Name
				fromMethod := source.Addr().MethodByName(name)
				toField := dest.FieldByName(name)

				if fromMethod.IsValid() && toField.IsValid() {
					values := fromMethod.Call([]reflect.Value{})
					if len(values) >= 1 {
						toField.Set(values[0])
					}
				}
			}
		}

		if isSlice {
			toElem.Set(reflect.Append(toElem, dest))
		}
	}
	return
}

func SameStruct(a, b interface{}) bool {
	panic("TODO")
}
