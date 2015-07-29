// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package epi

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestNext(t *testing.T) {
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

func benchNext(b *testing.B, size int) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		data := rand.New(rand.NewSource(int64(i))).Perm(size)
		b.StartTimer()
		NextPerm(data)
		b.StopTimer()
	}
}

func BenchmarkNext1e2(b *testing.B) { benchNext(b, 1e2) }
func BenchmarkNext1e4(b *testing.B) { benchNext(b, 1e4) }
func BenchmarkNext1e6(b *testing.B) { benchNext(b, 1e6) }
