// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package builtin

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func MinMaxInt(a, b int) (min, max int) {
	if a < b {
		return a, b
	}
	return b, a
}

func MinInt32(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}
func MaxInt32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
func MinMaxInt32(a, b int32) (min, max int32) {
	if a < b {
		return a, b
	}
	return b, a
}

func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
func MinMaxInt64(a, b int64) (min, max int64) {
	if a < b {
		return a, b
	}
	return b, a
}

func MinFloat32(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}
func MaxFloat32(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}
func MinMaxFloat32(a, b float32) (min, max float32) {
	if a < b {
		return a, b
	}
	return b, a
}

func MinFloat64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func MaxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
func MinMaxFloat64(a, b float64) (min, max float64) {
	if a < b {
		return a, b
	}
	return b, a
}
