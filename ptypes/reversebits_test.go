// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package ptypes

import "testing"

func TestReverseBits(t *testing.T) {
	for _, test := range []struct {
		in   uint16
		want uint16
	}{
		{1, 1 << 15},
		{2, 1 << 14},
		{4, 1 << 13},
		{8, 1 << 12},
		{0x000f, 0xf000},
		{0x00f0, 0x0f00},
		{0x00ff, 0xff00},
		{0xffff, 0xffff},
	} {
		if got := ReverseBits(test.in); got != test.want {
			t.Errorf("ReverseBits(%#x) = %#x; want %#x", test.in, got, test.want)
		}
	}
}

func BenchmarkReverseBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseBits(uint16(i))
	}
}

func TestReverseBitsLookup(t *testing.T) {
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
		if got := ReverseBitsLookup(test.in); got != test.want {
			t.Errorf("ReverseBitsLookup(%#x) = % #x; want %#x", test.in, got, test.want)
		}
	}
}

func BenchmarkReverseBitsLookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseBitsLookup(uint64(i))
	}
}
