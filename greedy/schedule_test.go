// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package greedy

import (
	"testing"
	"math/rand"
)

func TestMinWaitingTime(t *testing.T) {
	for _, test := range []struct {
		in   []int
		want int
	}{
		{nil, 0},
		{[]int{1}, 0},
		{[]int{1, 2}, 1},
		{[]int{2, 3}, 2},
		{[]int{1, 2, 3}, 4},
		{[]int{1, 3, 2}, 4},
		{[]int{2, 1, 3}, 4},
		{[]int{2, 3, 1}, 4},
		{[]int{3, 1, 2}, 4},
		{[]int{3, 2, 1}, 4},
		{[]int{2, 3, 1, 5, 8, 4, 6, 9, 7, 0}, 120},
	} {
		if got := MinWaitingTime(test.in); got != test.want {
			t.Errorf("MinWaitingTime(%v) = %d; want %d", test.in, got, test.want)
		}
	}
}

func benchMinWaitingTime(b *testing.B, size int) {
	tasks :=  rand.New(rand.NewSource(int64(size))).Perm(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinWaitingTime(tasks)
	}
}

func BenchmarkMinWaitingTime1e2(b *testing.B) { benchMinWaitingTime(b, 1e2) }
func BenchmarkMinWaitingTime1e4(b *testing.B) { benchMinWaitingTime(b, 1e4) }
func BenchmarkMinWaitingTime1e6(b *testing.B) { benchMinWaitingTime(b, 1e6) }
