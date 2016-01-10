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

func TestMergeInPlace(t *testing.T) {
	for _, test := range []struct {
		x, y []int
		want []int
	}{
		{[]int(nil), []int(nil), []int(nil)},
		{[]int(nil), []int{1, 2}, []int{1, 2}},
		{[]int{1, 2}, []int(nil), []int{1, 2}},
		{[]int{1, 3}, []int{2}, []int{1, 2, 3}},
		{[]int{2}, []int{1, 3}, []int{1, 2, 3}},
		{[]int{1, 2}, []int{3}, []int{1, 2, 3}},
		{[]int{1}, []int{2, 3}, []int{1, 2, 3}},
		{[]int{5, 13, 17}, []int{3, 7, 11, 19}, []int{3, 5, 7, 11, 13, 17, 19}},
	} {
		got := MergeInPlace(append([]int(nil), test.x...), test.y)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("MergeInPlace(%v, %v) got %v; want %v", test.x, test.y, got, test.want)
		}
	}
}

func benchMergeInPlace(b *testing.B, size int) {
	s := rand.NewSource(int64(size))
	x := rand.New(s).Perm(2 * size / 3)
	y := rand.New(s).Perm(size / 3)
	sort.Ints(x)
	sort.Ints(y)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeInPlace(x, y)
	}
}

func BenchmarkMergeInPlace1e2(b *testing.B) { benchMergeInPlace(b, 1e2) }
func BenchmarkMergeInPlace1e4(b *testing.B) { benchMergeInPlace(b, 1e4) }
func BenchmarkMergeInPlace1e6(b *testing.B) { benchMergeInPlace(b, 1e6) }
