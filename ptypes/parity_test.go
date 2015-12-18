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

type parityFn func(x int64) uint16

func testParityFn(t *testing.T, fn parityFn, fnName string) {
	for _, test := range []struct {
		in   int64
		want uint16
	}{
		{-9, odd},
		{-8, odd},
		{-7, even},
		{-6, even},
		{-5, odd},
		{-4, even},
		{-3, odd},
		{-2, odd},
		{-1, even},
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
		{math.MinInt64, odd},
		{math.MaxInt64, odd},
	} {
		if got := fn(test.in); got != test.want {
			t.Errorf("%s(%.64b) = %d; want %d", fnName, uint64(test.in), got, test.want)
		}
	}
}

func TestParity(t *testing.T)       { testParityFn(t, Parity, "Parity") }
func TestParityAlt(t *testing.T)    { testParityFn(t, ParityAlt, "ParityAlt") }
func TestParityLookup(t *testing.T) { testParityFn(t, ParityLookup, "ParityLookup") }

func benchParityFn(b *testing.B, fn parityFn) {
	for i := 0; i < b.N; i++ {
		fn(int64(i))
	}
}

func BenchmarkParity(b *testing.B)       { benchParityFn(b, Parity) }
func BenchmarkParityAlt(b *testing.B)    { benchParityFn(b, ParityAlt) }
func BenchmarkParityLookup(b *testing.B) { benchParityFn(b, ParityLookup) }
