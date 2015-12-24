// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package greedy

import (
	"reflect"
	"testing"
	"math/rand"
)

func TestPairTasks(t *testing.T) {
	for _, test := range []struct {
		in   []int
		want [][]int
	}{
		{nil, nil},
		{[]int{1}, [][]int{{1}}},
		{[]int{3, 1}, [][]int{{1, 3}}},
		{[]int{2, 3, 1}, [][]int{{1, 2}, {3}}},
		{[]int{5, 2, 1, 6, 4, 4}, [][]int{{1, 6}, {2, 5}, {4, 4}}},
	} {
		if got := PairTasks(test.in); !reflect.DeepEqual(got, test.want) {
			t.Errorf("PairTasks(%v) = %v; want %v", test.in, got, test.want)
		}
	}

}

func benchPairTasks(b *testing.B, size int) {
	tasks :=  rand.New(rand.NewSource(int64(size))).Perm(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PairTasks(tasks)
	}
}

func BenchmarkPairTasks1e2(b *testing.B) { benchPairTasks(b, 1e2) }
func BenchmarkPairTasks1e4(b *testing.B) { benchPairTasks(b, 1e4) }
func BenchmarkPairTasks1e6(b *testing.B) { benchPairTasks(b, 1e6) }