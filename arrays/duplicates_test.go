// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package arrays

import (
	"math/rand"
	"reflect"
	"testing"
	"sort"
)

func TestDeleteDuplicates(t *testing.T) {
	for _, test := range []struct {
		an     []int
		want   int
		update []int
	}{
		{nil, 0, nil},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 1}, 1, []int{1, 1}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 1, 2, 2, 3, 3}, 3, []int{1, 2, 3, 2, 3, 3}},
		{[]int{2, 3, 5, 5, 7, 11, 11, 11, 13}, 6, []int{2, 3, 5, 7, 11, 13, 11, 11, 13}},
	} {
		an := append([]int(nil), test.an...)
		if got := DeleteDuplicates(an); got != test.want {
			t.Errorf("DeleteDuplicates(%v) = %d; want %d", test.an, got, test.want)
		}
		if !reflect.DeepEqual(an, test.update) {
			t.Errorf("DeleteDuplicates(%v): updated slice %v; want %v", test.an, an, test.update)
		}
	}
}

func benchDeleteDuplicates(b *testing.B, size int) {
	b.StopTimer()
	ints := rand.New(rand.NewSource(int64(size))).Perm(size)
	sort.Ints(ints)
	an := make([]int, size)
	for i := 0; i < b.N; i++ {
		copy(an, ints)
		b.StartTimer()
		DeleteDuplicates(an)
		b.StopTimer()
	}
}

func BenchmarkDeleteDuplicates1e3(b *testing.B) { benchDeleteDuplicates(b, 1e3) }
func BenchmarkDeleteDuplicates1e5(b *testing.B) { benchDeleteDuplicates(b, 1e5) }
func BenchmarkDeleteDuplicates1e7(b *testing.B) { benchDeleteDuplicates(b, 1e7) }
