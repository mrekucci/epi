// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package bsearch

import (
	"math"
	"testing"
)

func TestSquareRootReal(t *testing.T) {
	for _, test := range []float64{-2.0, -1.0, -0, 5, 0.0, 0.25, 0.5, 0.75, 1.0, 2.0, 3.0, 4.0, math.NaN(), math.Inf(1)} {
		got, want := SquareRootReal(test), math.Sqrt(test)
		if cmp(got, want) != 0 {
			t.Errorf("SquareRootReal(%v) = %v; want %v", test, got, want)
		}
	}
}

func benchSquareRootReal(b *testing.B, size float64) {
	for i := 0; i < b.N; i++ {
		SquareRootReal(size)
	}
}

func BenchmarkSquareRootReal1e10(b *testing.B) { benchSquareRootReal(b, 1e10) }
func BenchmarkSquareRootReal1e20(b *testing.B) { benchSquareRootReal(b, 1e20) }
func BenchmarkSquareRootReal1e30(b *testing.B) { benchSquareRootReal(b, 1e30) }
