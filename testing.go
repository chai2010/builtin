// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package builtin

import (
	"fmt"
	"math"
	"reflect"
	"regexp"
	"testing"
)

func Assert(t testing.TB, condition bool, args ...interface{}) {
	if !condition {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("Assert failed, %s", msg)
		} else {
			t.Fatal("Assert failed")
		}
	}
}

func AssertEQ(t testing.TB, expected, got interface{}, args ...interface{}) {
	if !reflect.DeepEqual(expected, got) {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertEQ failed, expected = %v, got = %v, %s", expected, got, msg)
		} else {
			t.Fatalf("AssertEQ failed, expected = %v, got = %v", expected, got)
		}
	}
}

func AssertNear(t testing.TB, expected, got, abs float64, args ...interface{}) {
	if math.Abs(expected-got) > abs {
		if msg := fmt.Sprint(args...); msg != "" {
			t.Fatalf("AssertNear failed, expected = %v, got = %v, abs = %v, %s", expected, got, abs, msg)
		} else {
			t.Fatalf("AssertNear failed, expected = %v, got = %v, abs = %v", expected, got, abs)
		}
	}
}

func AssertMatch(t *testing.T, expectedPattern, got string, args ...interface{}) {
	if matched, err := regexp.MatchString(expectedPattern, got); err != nil || !matched {
		if err != nil {
			if msg := fmt.Sprint(args...); msg != "" {
				t.Fatalf("AssertMatch failed, expected = %q, got = %v, err = %v, %s", expectedPattern, got, err, msg)
			} else {
				t.Fatalf("AssertMatch failed, expected = %q, got = %v, err = %v", expectedPattern, got, err)
			}
		} else {
			if msg := fmt.Sprint(args...); msg != "" {
				t.Fatalf("AssertMatch failed, expected = %q, got = %v, %s", expectedPattern, got, msg)
			} else {
				t.Fatalf("AssertMatch failed, expected = %q, got = %v", expectedPattern, got)
			}
		}
	}
}
