// Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package builtin

import (
	"fmt"
	"image/color"
	"reflect"
	"sort"
)

func ExampleByteSlice() {
	src := []color.Gray{color.Gray{0xAA}, color.Gray{0xBB}, color.Gray{0xCC}, color.Gray{0xDD}}
	dst := make([]byte, len(src))
	copy(byteSlice(dst), byteSlice(src))
	fmt.Printf("%X", dst)
	// Output: AABBCCDD
}

func ExampleSlice() {
	src := []byte{0xAA, 0xBB, 0xCC, 0xDD}
	dst := unknownSlice(src, reflect.TypeOf([]color.Gray(nil))).([]color.Gray)
	fmt.Printf("%X", dst)
	// Output:
	// [{AA} {BB} {CC} {DD}]
}

func ExampleSort_int() {
	arr := []int32{88, 56, 100, 2, 25}

	Sort(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr)
	// Output:
	// [2 25 56 88 100]
}

func ExampleSort_string() {
	arr := []string{"coffee", "flour", "tea"}

	Sort(arr, func(i, j int) bool {
		return arr[i] > arr[j] // descending
	})
	fmt.Println(arr)
	// Output:
	// [tea flour coffee]
}

func ExampleSortInterface() {
	arr := []int32{88, 56, 100, 2, 25}
	sort.Sort(SortInterface(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	}))
	fmt.Println(arr)
	// Output:
	// [2 25 56 88 100]
}

func ExampleMapSlice() {
	a := MapSlice([]float32{1, 2, 3, 4}, func(val interface{}) interface{} {
		return val.(float32) * 2
	})
	fmt.Printf("%T: %v\n", a, a)
	// Output:
	// []float32: [2 4 6 8]
}

func ExampleMapMap() {
	a := MapMap(map[int]string{1: "A", 2: "B", 3: "C", 4: "D"}, func(key, val interface{}) interface{} {
		return fmt.Sprintf("{key(%v),val(%v)}", key.(int), val.(string))
	})
	fmt.Printf("%T\n", a)
	fmt.Printf("%v\n", a.(map[int]string)[1])
	fmt.Printf("%v\n", a.(map[int]string)[2])
	fmt.Printf("%v\n", a.(map[int]string)[3])
	fmt.Printf("%v\n", a.(map[int]string)[4])
	// Output:
	// map[int]string
	// {key(1),val(A)}
	// {key(2),val(B)}
	// {key(3),val(C)}
	// {key(4),val(D)}
}
