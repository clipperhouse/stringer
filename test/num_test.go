// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Signed integers spanning zero.

package main

import (
	"fmt"
	"testing"
)

// +test stringer
type Num int

const (
	m_2 Num = -2 + iota
	m_1
	m0
	m1
	m2
)

func TestNum(t *testing.T) {
	ckNum(t, -3, "Num(-3)")
	ckNum(t, m_2, "m_2")
	ckNum(t, m_1, "m_1")
	ckNum(t, m0, "m0")
	ckNum(t, m1, "m1")
	ckNum(t, m2, "m2")
	ckNum(t, 3, "Num(3)")
}

func ckNum(t *testing.T, num Num, str string) {
	if fmt.Sprint(num) != str {
		t.Errorf("expected %s, got %s", str, fmt.Sprint(num))
	}
}
