// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package ptypes

import "testing"

func TestSwapBits(t *testing.T) {
	for _, test := range []struct {
		x, i, j uint64
		want    uint64
	}{
		{1, 0, 1, 1 << 1},
		{5, 2, 4, 17},
		{20, 2, 4, 20},
		{1, 0, 63, 1 << 63},
	} {
		if got := SwapBits(test.x, test.i, test.j); got != test.want {
			t.Errorf("SwapBits(%.64b, %d, %d) = %.64b; want %.64b", test.x, test.i, test.j, got, test.want)
		}
	}
}

func BenchmarkSwapBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SwapBits(uint64(i), 0, 63)
	}
}
