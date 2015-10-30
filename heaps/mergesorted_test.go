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

func TestMergeSorted(t *testing.T) {
	for _, test := range []struct {
		ss   [][]int
		want []int
	}{
		{[][]int{[]int(nil), []int(nil)}, []int(nil)},
		{[][]int{[]int{1, 2, 3}, []int(nil)}, []int{1, 2, 3}},
		{[][]int{[]int{3, 5, 7}, []int{0, 6}, []int{0, 6, 28}}, []int{0, 0, 3, 5, 6, 6, 7, 28}},
		{[][]int{[]int{2, 3, 5}, []int{3, 4, 6}, []int{1, 1, 1}}, []int{1, 1, 1, 2, 3, 3, 4, 5, 6}},
	} {
		if got := MergeSorted(test.ss); !reflect.DeepEqual(got, test.want) {
			t.Errorf("MergeSorted(%v) = %v; want %v", test.ss, got, test.want)
		}
	}
}

func benchMergeSorted(b *testing.B, size int) {
	b.StopTimer()
	j := size / 3
	for i := 0; i < b.N; i++ {
		data := rand.New(rand.NewSource(int64(i))).Perm(size)
		sort.Ints(data)
		x, y, z := data[:j], data[j+1:2*j], data[2*j+1:]
		b.StartTimer()
		MergeSorted([][]int{x, y, z})
		b.StopTimer()
	}
}

func BenchmarkMergeSorted1e2(b *testing.B) { benchMergeSorted(b, 1e2) }
func BenchmarkMergeSorted1e4(b *testing.B) { benchMergeSorted(b, 1e4) }
func BenchmarkMergeSorted1e6(b *testing.B) { benchMergeSorted(b, 1e6) }
