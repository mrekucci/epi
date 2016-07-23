// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package recursion

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	for _, test := range []struct {
		in   []int
		want [][]int
	}{
		{nil, nil},
		{[]int{1}, [][]int{{1}}},
		{[]int{1, 2}, [][]int{{1, 2}, {2, 1}}},
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}}},
	} {
		if got := Permutations(test.in); !reflect.DeepEqual(got, test.want) {
			t.Errorf("Permutations(%v) = %v; want %v", test.in, got, test.want)
		}
	}
}

func benchPermutations(b *testing.B, size int) {
	xs := rand.New(rand.NewSource(int64(size))).Perm(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Permutations(xs)
	}
}

func BenchmarkPermutations5(b *testing.B) { benchPermutations(b, 5) }
func BenchmarkPermutations7(b *testing.B) { benchPermutations(b, 7) }
func BenchmarkPermutations9(b *testing.B) { benchPermutations(b, 9) }
