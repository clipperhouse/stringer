// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Simple test: enumeration of type int starting at 0.

package main

import (
	"fmt"
	"testing"
)

// +test stringer
type Day int

const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func TestDay(t *testing.T) {
	ckDay(t, Monday, "Monday")
	ckDay(t, Tuesday, "Tuesday")
	ckDay(t, Wednesday, "Wednesday")
	ckDay(t, Thursday, "Thursday")
	ckDay(t, Friday, "Friday")
	ckDay(t, Saturday, "Saturday")
	ckDay(t, Sunday, "Sunday")
	ckDay(t, -127, "Day(-127)")
	ckDay(t, 127, "Day(127)")
}

func ckDay(t *testing.T, day Day, str string) {
	if fmt.Sprint(day) != str {
		t.Errorf("expected %s, got %s", str, fmt.Sprint(day))
	}
}
