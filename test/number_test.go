// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Enumeration with an offset.
// Also includes a duplicate.

package main

import (
	"fmt"
	"testing"
)

// +test stringer
type Number int

const (
	_ Number = iota
	One
	Two
	Three
	AnotherOne = One // Duplicate; note that AnotherOne doesn't appear below.
)

func TestNumber(t *testing.T) {
	ckNumber(t, One, "One")
	ckNumber(t, Two, "Two")
	ckNumber(t, Three, "Three")
	ckNumber(t, AnotherOne, "One")
	ckNumber(t, 127, "Number(127)")
}

func ckNumber(t *testing.T, num Number, str string) {
	if fmt.Sprint(num) != str {
		t.Errorf("expected %s, got %s", str, fmt.Sprint(num))
	}
}
