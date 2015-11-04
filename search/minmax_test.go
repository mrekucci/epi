// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package search

import (
	"math/rand"
	"testing"
)

type minMaxFn func(an []int) (min, max int)

func testMinMaxFn(t *testing.T, fn minMaxFn, fnName string) {
	for _, test := range []struct {
		in       []int
		min, max int
	}{
		{[]int(nil), 0, 0},
		{[]int{1}, 1, 1},
		{[]int{-1, 1}, -1, 1},
		{[]int{1, -1}, -1, 1},
		{[]int{-1, 0, 1}, -1, 1},
		{[]int{1, 0, -1}, -1, 1},
		{[]int{3, 2, 5, 1, 2, 4}, 1, 5},
	} {
		if gotMin, gotMax := fn(test.in); gotMin != test.min || gotMax != test.max {
			t.Errorf("%s(%v) = %d, %d; want %d, %d", fnName, test.in, gotMin, gotMax, test.min, test.max)
		}
	}
}

func TestMinMax(t *testing.T)    { testMinMaxFn(t, MinMax, "MinMax") }
func TestMinMaxAlt(t *testing.T) { testMinMaxFn(t, MinMaxAlt, "MinMaxAlt") }

func benchMinMaxFn(b *testing.B, size int, fn minMaxFn) {
	b.StopTimer()
	data := rand.Perm(size)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		fn(data)
	}
}

func BenchmarkMinMax1e2(b *testing.B)    { benchMinMaxFn(b, 1e2, MinMax) }
func BenchmarkMinMaxAlt1e2(b *testing.B) { benchMinMaxFn(b, 1e2, MinMaxAlt) }
func BenchmarkMinMax1e4(b *testing.B)    { benchMinMaxFn(b, 1e4, MinMax) }
func BenchmarkMinMaxAlt1e4(b *testing.B) { benchMinMaxFn(b, 1e4, MinMaxAlt) }
func BenchmarkMinMax1e6(b *testing.B)    { benchMinMaxFn(b, 1e6, MinMax) }
func BenchmarkMinMaxAlt1e6(b *testing.B) { benchMinMaxFn(b, 1e6, MinMaxAlt) }
