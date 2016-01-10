// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package sorting

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestIntersectSorted(t *testing.T) {
	for _, test := range []struct {
		x, y []int
		want []int
	}{
		{[]int(nil), []int(nil), []int(nil)},
		{[]int{0}, []int(nil), []int(nil)},
		{[]int(nil), []int{0}, []int(nil)},
		{[]int{0}, []int{1}, []int(nil)},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{0, 1}, []int{2}, []int(nil)},
		{[]int{0}, []int{0, 1}, []int{0}},
		{[]int{0, 1}, []int{1}, []int{1}},
		{[]int{0, 1}, []int{1, 2}, []int{1}},
		{[]int{1, 1, 1}, []int{1, 1}, []int{1}},
		{[]int{1, 1}, []int{1, 1, 1}, []int{1}},
		{[]int{1, 1, 2, 3, 4}, []int{0, 0, 1}, []int{1}},
		{[]int{0, 0, 1}, []int{1, 1, 2, 3, 4}, []int{1}},
		{[]int{2, 3, 3, 5, 5, 6, 7, 7, 8, 12}, []int{5, 5, 6, 8, 8, 9, 10, 10}, []int{5, 6, 8}},
	} {
		if got := IntersectSorted(test.x, test.y); !reflect.DeepEqual(got, test.want) {
			t.Errorf("IntersectSorted(%v, %v) = %v; want %v", test.x, test.y, got, test.want)
		}
	}
}

func benchIntersectSorted(b *testing.B, size int) {
	x := rand.New(rand.NewSource(int64(size))).Perm(size)
	sort.Ints(x)
	y := x[size-(size/2):]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IntersectSorted(x, y)
	}
}

func BenchmarkIntersectSorted1e2(b *testing.B) { benchIntersectSorted(b, 1e2) }
func BenchmarkIntersectSorted1e4(b *testing.B) { benchIntersectSorted(b, 1e4) }
func BenchmarkIntersectSorted1e6(b *testing.B) { benchIntersectSorted(b, 1e6) }
