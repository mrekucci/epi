// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package epi

import "testing"

var swapBitsTests = []struct {
	x    uint64
	i, j uint64
	want uint64
}{
	{1, 0, 1, 1 << 1},
	{5, 2, 4, 17},
	{20, 2, 4, 20},
	{1, 0, 63, 1 << 63},
}

func TestSwapBits(t *testing.T) {
	for _, tt := range swapBitsTests {
		got := SwapBits(tt.x, tt.i, tt.j)
		if got != tt.want {
			t.Errorf("SwapBits(%d, %d, %d) = %d; want %d", tt.x, tt.i, tt.j, got, tt.want)
		}
	}
}

func BenchmarkSwapBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SwapBits(uint64(i), 0, 63)
	}
}
