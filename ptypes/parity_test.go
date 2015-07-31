// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package ptypes

import (
	"math"
	"testing"
)

const (
	odd  = 1
	even = 0
)

func TestParity(t *testing.T) {
	for _, test := range []struct {
		in   uint16
		want uint16
	}{
		{0, even},
		{1, odd},
		{2, odd},
		{3, even},
		{4, odd},
		{5, even},
		{6, even},
		{7, odd},
		{8, odd},
		{9, even},
		{math.MaxUint16, even},
	} {
		if got := Parity(test.in); got != test.want {
			t.Errorf("Parity(%d) = %d; want %d", test.in, got, test.want)
		}
	}
}

func BenchmarkParity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parity(uint16(i))
	}
}

func TestParityLookup(t *testing.T) {
	for _, test := range []struct {
		in   uint64
		want uint16
	}{
		{0, even},
		{1, odd},
		{2, odd},
		{3, even},
		{4, odd},
		{5, even},
		{6, even},
		{7, odd},
		{8, odd},
		{9, even},
		{math.MaxUint16, even},
	} {
		if got := ParityLookup(test.in); got != test.want {
			t.Errorf("ParityLookup(%d) = %d; want %d", test.in, got, test.want)
		}
	}
}

func BenchmarkParityLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParityLookup(uint64(i))
	}
}
