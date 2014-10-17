// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package builtin

import (
	"unsafe"
)

func Uint16(b []byte) uint16 {
	return *(*uint16)(unsafe.Pointer(&b[0]))
}

func PutUint16(b []byte, v uint16) {
	*(*uint16)(unsafe.Pointer(&b[0])) = v
}

func Uint32(b []byte) uint32 {
	return *(*uint32)(unsafe.Pointer(&b[0]))
}

func PutUint32(b []byte, v uint32) {
	*(*uint32)(unsafe.Pointer(&b[0])) = v
}

func Uint64(b []byte) uint64 {
	return *(*uint64)(unsafe.Pointer(&b[0]))
}

func PutUint64(b []byte, v uint64) {
	*(*uint64)(unsafe.Pointer(&b[0])) = v
}

func Float32(b []byte) float32 {
	return *(*float32)(unsafe.Pointer(&b[0]))
}

func PutFloat32(b []byte, v float32) {
	*(*float32)(unsafe.Pointer(&b[0])) = v
}

func Float64(b []byte) float64 {
	return *(*float64)(unsafe.Pointer(&b[0]))
}

func PutFloat64(b []byte, v float64) {
	*(*float64)(unsafe.Pointer(&b[0])) = v
}
