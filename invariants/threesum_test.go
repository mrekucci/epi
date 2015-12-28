// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package invariants

import (
	"math/rand"
	"testing"
)

func TestHasThreeSum(t *testing.T) {
	for _, test := range []struct {
		an   []int
		k    int
		want bool
	}{
		{nil, 0, false},
		{[]int{1}, 1, false},
		{[]int{1, 2}, 2, false},
		{[]int{1, 2}, 3, true},
		{[]int{1, 2}, 4, true},
		{[]int{1, 2}, 5, true},
		{[]int{-1, 2}, 0, true},
		{[]int{-1, -2}, -5, true},
		{[]int{1, 1, 1}, 1, false},
		{[]int{1, 1, 1}, 2, false},
		{[]int{1, 1, 1}, 3, true},
		{[]int{1, 1, 1}, 4, false},
		{[]int{5, 7, 2, 3, 11}, 21, true},
		{[]int{-3, 7, -5, 2, 9}, -1, true},
	} {
		if got := HasThreeSum(test.an, test.k); got != test.want {
			t.Errorf("HasThreeSum(%v, %d) = %t; want %t", test.an, test.k, got, test.want)
		}
	}
}

func benchHasThreeSum(b *testing.B, size int) {
	an := rand.New(rand.NewSource(int64(size))).Perm(size)
	k := an[0] + an[size/3] + an[size-1]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HasThreeSum(an, k)
	}
}

func BenchmarkHasThreeSum1e2(b *testing.B) { benchHasThreeSum(b, 1e2) }
func BenchmarkHasThreeSum1e4(b *testing.B) { benchHasThreeSum(b, 1e4) }
func BenchmarkHasThreeSum1e6(b *testing.B) { benchHasThreeSum(b, 1e6) }
