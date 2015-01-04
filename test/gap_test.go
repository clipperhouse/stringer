// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Gaps and an offset.

package main

import (
	"fmt"
	"testing"
)

// +test stringer
type Gap int

const (
	gTwo    Gap = 2
	gThree  Gap = 3
	gFive   Gap = 5
	gSix    Gap = 6
	gSeven  Gap = 7
	gEight  Gap = 8
	gNine   Gap = 9
	gEleven Gap = 11
)

func TestGap(t *testing.T) {
	ckGap(t, 0, "Gap(0)")
	ckGap(t, 1, "Gap(1)")
	ckGap(t, gTwo, "gTwo")
	ckGap(t, gThree, "gThree")
	ckGap(t, 4, "Gap(4)")
	ckGap(t, gFive, "gFive")
	ckGap(t, gSix, "gSix")
	ckGap(t, gSeven, "gSeven")
	ckGap(t, gEight, "gEight")
	ckGap(t, gNine, "gNine")
	ckGap(t, 10, "Gap(10)")
	ckGap(t, gEleven, "gEleven")
	ckGap(t, 12, "Gap(12)")
}

func ckGap(t *testing.T, gap Gap, str string) {
	if fmt.Sprint(gap) != str {
		t.Errorf("expected %s, got %s", str, fmt.Sprint(gap))
	}
}
