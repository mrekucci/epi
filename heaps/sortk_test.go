// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package heaps

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestSortIncDecK(t *testing.T) {
	// Integer limit values.
	const (
		maxInt = int(^uint(0) >> 1)
		minInt = -maxInt - 1
	)
	for _, test := range [][]int{
		[]int{0},
		[]int{1, 0},
		[]int{0, 1},
		[]int{minInt, 0, maxInt},
		[]int{maxInt, 0, minInt},
		[]int{-1, 1, -1, 1},
		[]int{1, -1, 1, -1},
		[]int{-1, 0, 1, 0, -1},
		[]int{1, 0, -1, 0, 1},
		[]int{57, 131, 493, 294, 221, 339, 418, 452, 442, 190},
		[]int{1, 2, 3, 2, 1, 4, 5, 10, 9, 4, 4, 1, -1},
	} {
		want := append([]int(nil), test...)
		sort.Ints(want)
		if got := SortK(append([]int(nil), test...)); !reflect.DeepEqual(got, want) {
			t.Errorf("SortIncDecK(%v)", test)
			t.Error(" got ", got)
			t.Error("want ", want)
		}
	}
}

func benchSortK(b *testing.B, size int) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := rand.New(rand.NewSource(int64(i))).Perm(size)
		o := -1
		for p, q := 0, size/10; q < size; p, q = q, p+size/10 {
			s := sort.IntSlice(data[p:q])
			if o < 0 {
				sort.Sort(sort.Reverse(s))
			} else {
				sort.Sort(s)
			}
		}
		b.StartTimer()
		SortK(data)
		b.StopTimer()
	}
}

func BenchmarkSortK1e2(b *testing.B) { benchSortK(b, 1e2) }
func BenchmarkSortK1e4(b *testing.B) { benchSortK(b, 1e4) }
func BenchmarkSortK1e6(b *testing.B) { benchSortK(b, 1e6) }
