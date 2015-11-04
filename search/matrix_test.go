// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package search

import (
	"math/rand"
	"sort"
	"testing"
)

func TestMatrix(t *testing.T) {
	for _, test := range []struct {
		m    [][]int
		x    int
		want bool
	}{
		{[][]int{[]int{0, 1}, []int{1, 2}}, 2, true},
		{[][]int{[]int{0, 1}, []int{1, 2}}, -1, false},
		{[][]int{[]int{0, 1, 2}, []int{1, 2, 3}, []int{2, 3, 4}}, 2, true},
		{[][]int{[]int{0, 1, 2}, []int{1, 2, 3}, []int{2, 3, 4}}, -1, false},
	} {
		if got := Matrix(test.m, test.x); got != test.want {
			t.Errorf("Matrix(%v, %d) = %t; want %t", test.m, test.x, got, test.want)
		}
	}
}

func benchMatrix(b *testing.B, size int) {
	b.StopTimer()
	n := size * size
	data := rand.Perm(n)
	sort.Ints(data)
	x := data[n/2]
	var m [][]int
	for i := 1; i <= size; i++ {
		m = append(m, data[(i-1)*size:i*size])
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Matrix(m, x)
	}
}

func BenchmarkMatrix1e1(b *testing.B) { benchMatrix(b, 1e1) }
func BenchmarkMatrix1e2(b *testing.B) { benchMatrix(b, 1e2) }
func BenchmarkMatrix1e3(b *testing.B) { benchMatrix(b, 1e3) }
