// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package builtin

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestIf(t *testing.T) {
	a, b := 2, 3
	max := If(a > b, a, b).(int)
	if max != 3 {
		t.Fatalf("expect = %v, got = %v", 3, max)
	}
}

func TestByteSlice(t *testing.T) {
	type Point struct{ X, Y int }

	src := []Point{Point{1, 1}, Point{2, 2}, Point{3, 3}}
	dst := make([]Point, len(src))

	copy(ByteSlice(dst), ByteSlice(src))
	if !reflect.DeepEqual(dst, src) {
		t.Fatal("not equal")
	}
}

func TestSlice(t *testing.T) {
	type Point struct{ X, Y int32 }

	src := []Point{Point{1, 1}, Point{2, 2}, Point{3, 3}}
	dst := make([]Point, len(src))

	copy(
		Slice(dst, reflect.TypeOf([]byte(nil))).([]byte),
		Slice(src, reflect.TypeOf([]byte(nil))).([]byte),
	)

	if !reflect.DeepEqual(dst, src) {
		t.Fatal("not equal")
	}
}

func TestSlice_newType(t *testing.T) {
	type Point struct{ X, Y int32 }

	src := []Point{Point{1, 1}, Point{2, 2}, Point{3, 3}}
	dst := make([]Point, len(src))

	hDst := Slice(dst, reflect.TypeOf([]int64(nil))).([]int64)
	hSrc := Slice(src, reflect.TypeOf([]int64(nil))).([]int64)

	if len(hSrc) != len(src) {
		t.Fatal("bad size")
	}
	if len(hDst) != len(dst) {
		t.Fatal("bad size")
	}
	if len(hSrc) != len(hDst) {
		t.Fatal("bad size")
	}

	// copy as int64
	for i := 0; i < len(hDst) && i < len(hSrc); i++ {
		hDst[i] = hSrc[i]
	}
	if !reflect.DeepEqual(dst, src) {
		t.Fatal("not equal")
	}
}

func TestSlice_notAlign(t *testing.T) {
	src := []uint32{0xAAAAAAAA, 0xBBBBBBBB, 0xCCCCCCCC, 0xDDDDDDDD, 0xEEEEEEEE, 0, 0}
	tmp := make([]byte, len(src)*4+1) // Does &tmp[0] align with 8?
	dst := Slice(tmp[1:], reflect.TypeOf([]uint32(nil))).([]uint32)

	if h := (*reflect.SliceHeader)((unsafe.Pointer(&dst))); h.Data&7 == 0 {
		t.Fatal("dst address is align with 8")
	}

	copy(ByteSlice(dst), ByteSlice(src))
	t.Logf("tmp: %x\n", tmp)
	t.Logf("dst: %x\n", dst)

	for i := 0; i < len(src) && i < len(dst); i++ {
		if src[i] != dst[i] {
			t.Fatalf("not equal, src[%d] = %x, dst[%d] = %x", i, src[i], i, dst[i])
		}
	}
}

// BUG(chai2010): check gc moving/compacting
func TestSlice_gc_moving_compacting(t *testing.T) {
	// how to check ?
}
