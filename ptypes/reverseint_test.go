// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package ptypes

import (
	"math"
	"testing"
)

func TestReverseInt(t *testing.T) {
	for _, test := range []struct {
		in   int64
		want int64
		ok   bool
	}{
		{0, 0, true},
		{1, 1, true},
		{-1, -1, true},
		{-10, -1, true},
		{1234, 4321, true},
		{-1234, -4321, true},
		{math.MaxInt64, 7085774586302733229, true},
		{7085774586302733229, math.MaxInt64, true},
		{math.MinInt64, -8085774586302733229, true},
		{-8085774586302733229, math.MinInt64, true},
		{8085774586302733229, 0, false},
		{-9085774586302733229, 0, false},
	} {
		if got, ok := ReverseInt(test.in); got != test.want || ok != test.ok {
			t.Errorf("ReverseInt(%d) = %d, %t; want %d, %t", test.in, got, ok, test.want, test.ok)
		}
	}
}

func benchReverseInt(b *testing.B, size int) {
	x := int64(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseInt(x)
	}
}

func BenchmarkReverseInt1e6(b *testing.B)  { benchReverseInt(b, 1e6) }
func BenchmarkReverseInt1e12(b *testing.B) { benchReverseInt(b, 1e12) }
func BenchmarkReverseInt1e18(b *testing.B) { benchReverseInt(b, 1e18) }
