// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package search

import (
	"math/rand"
	"testing"
)

func TestLargestKth(t *testing.T) {
	for _, test := range []struct {
		an []int
		k  int
		e  int
		ok bool
	}{
		{[]int(nil), 0, 0, false},
		{[]int{0}, -1, 0, false},
		{[]int{-1}, 2, 0, false},
		{[]int{-1}, 1, -1, true},
		{[]int{0, 1, 2, 3, 4}, 1, 4, true},
		{[]int{0, 1, 2, 3, 4}, 2, 3, true},
		{[]int{0, 1, 2, 3, 4}, 3, 2, true},
		{[]int{0, 1, 2, 3, 4}, 4, 1, true},
		{[]int{0, 1, 2, 3, 4}, 5, 0, true},
		{[]int{-4, -3, -2, -1, 0}, 1, 0, true},
		{[]int{-4, -3, -2, -1, 0}, 3, -2, true},
		{[]int{-4, -3, -2, -1, 0}, 5, -4, true},
	} {
		an := append([]int(nil), test.an...) // Copy the slice to avoid modification of the original slice.
		if got, ok := LargestKth(an, test.k); got != test.e || ok != test.ok {
			t.Errorf("LargestKth(%v, %d) = %d, %t; want %d, %t", test.an, test.k, got, ok, test.e, test.ok)
		}
	}
}

func benchLargestKth(b *testing.B, size int) {
	b.StopTimer()
	data := rand.Perm(size)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		LargestKth(data, size)
	}
}

func BenchmarkLargestKth1e1(b *testing.B) { benchLargestKth(b, 1e1) }
func BenchmarkLargestKth1e2(b *testing.B) { benchLargestKth(b, 1e2) }
func BenchmarkLargestKth1e3(b *testing.B) { benchLargestKth(b, 1e3) }
