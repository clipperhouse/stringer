// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Enough gaps to trigger a map implementation of the method.
// Also includes a duplicate to test that it doesn't cause problems

package main

import (
	"fmt"
	"testing"
)

// +test stringer
type Prime int

const (
	p2  Prime = 2
	p3  Prime = 3
	p5  Prime = 5
	p7  Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)

func TestPrime(t *testing.T) {
	ckPrime(t, 0, "Prime(0)")
	ckPrime(t, 1, "Prime(1)")
	ckPrime(t, p2, "p2")
	ckPrime(t, p3, "p3")
	ckPrime(t, 4, "Prime(4)")
	ckPrime(t, p5, "p5")
	ckPrime(t, p7, "p7")
	ckPrime(t, p77, "p7")
	ckPrime(t, p11, "p11")
	ckPrime(t, p13, "p13")
	ckPrime(t, p17, "p17")
	ckPrime(t, p19, "p19")
	ckPrime(t, p23, "p23")
	ckPrime(t, p29, "p29")
	ckPrime(t, p37, "p37")
	ckPrime(t, p41, "p41")
	ckPrime(t, p43, "p43")
	ckPrime(t, 44, "Prime(44)")
}

func ckPrime(t *testing.T, prime Prime, str string) {
	if fmt.Sprint(prime) != str {
		t.Errorf("expected %s, got %s", str, fmt.Sprint(prime))
	}
}
