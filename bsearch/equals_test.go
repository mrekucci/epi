// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package bsearch

import "testing"

func TestIndexEqualsEntry(t *testing.T) {
	for _, test := range []struct {
		in   []int
		want int
	}{
		{[]int(nil), -1},
		{[]int{0}, 0},
		{[]int{1}, -1},
		{[]int{0, 1}, 0},
		{[]int{1, 0}, -1},
		{[]int{0, 1, 2}, 1},
		{[]int{1, 2, 3}, -1},
		{[]int{0, 1, 2, 3}, 1},
		{[]int{-2, -1, 0, 1, 2}, -1},
		{[]int{2, 2, 2, 2, 2}, 2},
	} {
		if got := IndexEqualsEntry(test.in); got != test.want {
			t.Errorf("IndexEqualsEntry(%v) = %d; want %d", test.in, got, test.want)
		}
	}

}

func benchIndexEqualsEntry(b *testing.B, size int) {
	b.StopTimer()
	data := make([]int, size) // Will be searching for 0.
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IndexEqualsEntry(data)
	}
}

func BenchmarkIndexEqualsEntry1e2(b *testing.B) { benchIndexEqualsEntry(b, 1e2) }
func BenchmarkIndexEqualsEntry1e4(b *testing.B) { benchIndexEqualsEntry(b, 1e4) }
func BenchmarkIndexEqualsEntry1e6(b *testing.B) { benchIndexEqualsEntry(b, 1e6) }
