// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Unsigned integers spanning zero.

package main

import (
	"fmt"
	"testing"
)

// +test stringer
type Unum uint8

const (
	u_2 Unum = iota + 253
	u_1
)

const (
	u0 Unum = iota
	u1
	u2
)

func TestUnum(t *testing.T) {
	ckUnum(t, ^Unum(0)-3, "Unum(252)")
	ckUnum(t, u_2, "u_2")
	ckUnum(t, u_1, "u_1")
	ckUnum(t, u0, "u0")
	ckUnum(t, u1, "u1")
	ckUnum(t, u2, "u2")
	ckUnum(t, 3, "Unum(3)")
}

func ckUnum(t *testing.T, unum Unum, str string) {
	if fmt.Sprint(unum) != str {
		t.Errorf("expected %s, got %s", str, fmt.Sprint(unum))
	}
}
