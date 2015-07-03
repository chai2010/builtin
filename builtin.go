// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// please copy this file and fix the package name.

package builtin

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"unsafe"
)

func callerFileLine() (file string, line int) {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		// Truncate file name at last file name separator.
		if index := strings.LastIndex(file, "/"); index >= 0 {
			file = file[index+1:]
		} else if index = strings.LastIndex(file, "\\"); index >= 0 {
			file = file[index+1:]
		}
	} else {
		file = "???"
		line = 1
	}
	return
}

func assert(condition bool, a ...interface{}) {
	if !condition {
		file, line := callerFileLine()
		if msg := fmt.Sprint(a...); msg != "" {
			fmt.Fprintf(os.Stderr, "%s:%d: Assert failed, %s", file, line, msg)
		} else {
			fmt.Fprintf(os.Stderr, "%s:%d: Assert failed", file, line)
		}
		os.Exit(1)
	}
}

func logf(format string, a ...interface{}) {
	file, line := callerFileLine()
	fmt.Fprintf(os.Stderr, "%s:%d: ", file, line)
	fmt.Fprintf(os.Stderr, format, a...)
}

func logln(a ...interface{}) {
	file, line := callerFileLine()
	fmt.Fprintf(os.Stderr, "%s:%d: ", file, line)
	fmt.Fprintln(os.Stderr, a...)
}

func fatalf(format string, a ...interface{}) {
	file, line := callerFileLine()
	fmt.Fprintf(os.Stderr, "%s:%d: ", file, line)
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}

func fatalln(a ...interface{}) {
	file, line := callerFileLine()
	fmt.Fprintf(os.Stderr, "%s:%d: ", file, line)
	fmt.Fprintln(os.Stderr, a...)
	os.Exit(1)
}

func panicf(format string, a ...interface{}) {
	file, line := callerFileLine()
	s := fmt.Sprintf("%s:%d: ", file, line)
	s += fmt.Sprint(a...)
	panic(s)
}

func panicln(a ...interface{}) {
	file, line := callerFileLine()
	s := fmt.Sprintf("%s:%d: ", file, line)
	s += fmt.Sprint(a...)
	panic(s)
}

func errorf(format string, a ...interface{}) error {
	file, line := callerFileLine()
	s := fmt.Sprintf("%s:%d: ", file, line)
	s += fmt.Sprintf(format, a...)
	return errors.New(s)
}

func printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}

func scan(a ...interface{}) (n int, err error) {
	return fmt.Scan(a...)
}

func scanf(format string, a ...interface{}) (n int, err error) {
	return fmt.Scanf(format, a...)
}

func scanln(a ...interface{}) (n int, err error) {
	return fmt.Scanln(a...)
}

func sprint(a ...interface{}) string {
	return fmt.Sprint(a...)
}

func sprintf(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

func sprintln(a ...interface{}) string {
	return fmt.Sprintln(a...)
}

func sscan(str string, a ...interface{}) (n int, err error) {
	return fmt.Sscan(str, a...)
}

func sscanf(str string, format string, a ...interface{}) (n int, err error) {
	return fmt.Sscanf(str, format, a...)
}

func sscanln(str string, a ...interface{}) (n int, err error) {
	return fmt.Sscanln(str, a...)
}

func byteSlice(d0 interface{}) (d1 []byte) {
	sv := reflect.ValueOf(d0)
	h := (*reflect.SliceHeader)((unsafe.Pointer(&d1)))
	h.Cap = sv.Cap() * int(sv.Type().Elem().Size())
	h.Len = sv.Len() * int(sv.Type().Elem().Size())
	h.Data = sv.Pointer()
	return
}

func uint16Slice(d0 []byte) (d1 []uint16) {
	h0 := (*reflect.SliceHeader)(unsafe.Pointer(&d0))
	h1 := (*reflect.SliceHeader)(unsafe.Pointer(&d1))

	h1.Cap = h0.Cap / 2
	h1.Len = h0.Len / 2
	h1.Data = h0.Data
	return
}

func uint32Slice(d0 []byte) (d1 []uint32) {
	h0 := (*reflect.SliceHeader)(unsafe.Pointer(&d0))
	h1 := (*reflect.SliceHeader)(unsafe.Pointer(&d1))

	h1.Cap = h0.Cap / 4
	h1.Len = h0.Len / 4
	h1.Data = h0.Data
	return
}

func float32Slice(d0 []byte) (d1 []float32) {
	h0 := (*reflect.SliceHeader)(unsafe.Pointer(&d0))
	h1 := (*reflect.SliceHeader)(unsafe.Pointer(&d1))

	h1.Cap = h0.Cap / 4
	h1.Len = h0.Len / 4
	h1.Data = h0.Data
	return
}

func float64Slice(d0 []byte) (d1 []float64) {
	h0 := (*reflect.SliceHeader)(unsafe.Pointer(&d0))
	h1 := (*reflect.SliceHeader)(unsafe.Pointer(&d1))

	h1.Cap = h0.Cap / 8
	h1.Len = h0.Len / 8
	h1.Data = h0.Data
	return
}

func unknownSlice(slice interface{}, newSliceType reflect.Type) interface{} {
	sv := reflect.ValueOf(slice)
	newSlice := reflect.New(newSliceType)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(newSlice.Pointer()))
	hdr.Cap = sv.Cap() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Len = sv.Len() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Data = uintptr(sv.Pointer())
	return newSlice.Elem().Interface()
}
