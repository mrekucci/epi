// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package epi

import (
	"math"
	"testing"
)

func TestIntWeight(t *testing.T) {
	for _, test := range []struct {
		in   uint64
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 2},
		{6, 2},
		{7, 3},
		{8, 1},
		{9, 2},
		{math.MaxUint64, 64},
	} {
		if got := IntWeight(test.in); got != test.want {
			t.Errorf("IntWeight(%d) = %d; want %d", test.in, got, test.want)
		}
	}
}

func BenchmarkIntWeight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntWeight(uint64(i))
	}
}

func TestClosestInt(t *testing.T) {
	for _, test := range []struct {
		in   uint64
		want uint64
		ok   bool
	}{
		{0x0000000000000000, 0x0, false},
		{0xffffffffffffffff, 0x0, false},
		{0x01, 0x02, true},
		{0x03, 0x05, true},
		{0x04, 0x02, true},
		{0x05, 0x06, true},
		{0x06, 0x05, true},
		{0x07, 0x0b, true},
		{0x08, 0x04, true},
		{0x09, 0x0a, true},
		{0x0a, 0x09, true},
		{0x0b, 0x0d, true},
		{0x0c, 0x0a, true},
		{0x0d, 0x0e, true},
		{0x0e, 0x0d, true},
		{0x0f, 0x17, true},
	} {
		if got, ok := ClosestInt(test.in); got != test.want || ok != test.ok {
			t.Errorf("ClosestInt(%#x) = %#x, %t; want %#x, %t", test.in, got, ok, test.want, test.ok)
		}
	}
}

func BenchmarkClosestInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClosestInt(uint64(i))
	}
}
