// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package ptypes

import "testing"

type reverseBitsFn func(uint64) uint64

func testReverseBitsFn(t *testing.T, fn reverseBitsFn, fnName string) {
	for _, test := range []struct {
		in   uint64
		want uint64
	}{
		{1, 1 << 63},
		{2, 1 << 62},
		{4, 1 << 61},
		{8, 1 << 60},
		{0x00000000ffffffff, 0xffffffff00000000},
		{0x00000000000000ff, 0xff00000000000000},
		{0x00000000ff000000, 0x000000ff00000000},
		{0xffffffffffffffff, 0xffffffffffffffff},
	} {
		if got := fn(test.in); got != test.want {
			t.Errorf("%s(%#.16x) = %#.16x; want %#.16x", fnName, test.in, got, test.want)
		}
	}
}

func TestReverseBits(t *testing.T)       { testReverseBitsFn(t, ReverseBits, "ReverseBits") }
func TestReverseBitsLookup(t *testing.T) { testReverseBitsFn(t, ReverseBitsLookup, "ReverseBitsLookup") }

func benchReverseBitsFn(b *testing.B, fn reverseBitsFn) {
	for i := 0; i < b.N; i++ {
		fn(0xffffffff00000000)
	}
}

func BenchmarkReverseBits(b *testing.B)       { benchReverseBitsFn(b, ReverseBits) }
func BenchmarkReverseBitsLookup(b *testing.B) { benchReverseBitsFn(b, ReverseBitsLookup) }
