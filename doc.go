// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package builtin provides some usefull functions.

Convert []X to []Y:
	x := make([]X, xLen)
	y := Slice(x, reflect.TypeOf([]X(nil))).([]Y)

or

	x := make([]X, xLen)
	y := ((*[1 << 30]Y)(unsafe.Pointer(&x[0])))[:yLen]

Convert []X to []byte:
	x := make([]X, xLen)
	y := ByteSlice(x, reflect.TypeOf([]X(nil)))

or

	x := make([]X, xLen)
	y := ((*[1 << 30]byte)(unsafe.Pointer(&x[0])))[:yLen]
*/
package builtin
