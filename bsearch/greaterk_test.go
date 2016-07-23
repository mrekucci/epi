// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package bsearch

import (
	"math/rand"
	"sort"
	"testing"
)

func TestFirstGreaterK(t *testing.T) {
	for _, test := range []struct {
		xs   []int
		k    int
		want int
	}{
		{[]int(nil), 0, -1},
		{[]int{0}, 1, -1},
		{[]int{0}, 0, -1},
		{[]int{0, 1}, 0, 1},
		{[]int{0, 1}, 1, -1},
		{[]int{0, 1, 2}, 0, 1},
		{[]int{0, 1, 2}, 1, 2},
		{[]int{0, 1, 2}, 2, -1},
		{[]int{0, 1, 2, 3}, 0, 1},
		{[]int{0, 1, 2, 3}, 1, 2},
		{[]int{0, 1, 2, 3}, 2, 3},
		{[]int{0, 1, 2, 3}, 3, -1},
		{[]int{0, 0, 1, 2, 3}, 0, 2},
		{[]int{0, 1, 1, 2, 3}, 1, 3},
		{[]int{2, 2, 2, 2, 2}, 2, -1},
	} {
		if got := FirstGreaterK(test.xs, test.k); got != test.want {
			t.Errorf("FirstGreaterK(%d %v) = %d; want %d", test.k, test.xs, got, test.want)
		}
	}
}

func benchFirstGreaterK(b *testing.B, size int) {
	data := rand.Perm(size)
	sort.Ints(data)
	k := data[size-1]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FirstGreaterK(data, k)
	}
}

func BenchmarkFirstGreaterK1e2(b *testing.B) { benchFirstGreaterK(b, 1e2) }
func BenchmarkFirstGreaterK1e4(b *testing.B) { benchFirstGreaterK(b, 1e4) }
func BenchmarkFirstGreaterK1e6(b *testing.B) { benchFirstGreaterK(b, 1e6) }
