// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package arrays

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestNextPerm(t *testing.T) {
	for _, test := range []struct {
		in   []int
		want []int
	}{
		{[]int{1}, []int{}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{2, 1}, []int{}},
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{1, 3, 2}, []int{2, 1, 3}},
		{[]int{2, 1, 3}, []int{2, 3, 1}},
		{[]int{2, 3, 1}, []int{3, 1, 2}},
		{[]int{3, 1, 2}, []int{3, 2, 1}},
		{[]int{3, 2, 1}, []int{}},
	} {
		if got := NextPerm(test.in); !reflect.DeepEqual(got, test.want) {
			t.Errorf("NextPerm(%d) = %d; want %d", test.in, got, test.want)
		}
	}
}

func benchNextPerm(b *testing.B, size int) {
	data := rand.New(rand.NewSource(int64(size))).Perm(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NextPerm(data)
	}
}

func BenchmarkNextPerm1e2(b *testing.B) { benchNextPerm(b, 1e2) }
func BenchmarkNextPerm1e4(b *testing.B) { benchNextPerm(b, 1e4) }
func BenchmarkNextPerm1e6(b *testing.B) { benchNextPerm(b, 1e6) }
