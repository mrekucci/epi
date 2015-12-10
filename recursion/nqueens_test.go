// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package recursion

import (
	"reflect"
	"testing"
)

func TestNQueens(t *testing.T) {
	for _, test := range []struct {
		in   int
		want [][]int
	}{
		{0, nil},
		{1, [][]int{{0}}},
		{2, nil},
		{3, nil},
		{4, [][]int{
			{1, 3, 0, 2},
			{2, 0, 3, 1}}},
		{5, [][]int{
			{0, 2, 4, 1, 3},
			{0, 3, 1, 4, 2},
			{1, 3, 0, 2, 4},
			{1, 4, 2, 0, 3},
			{2, 0, 3, 1, 4},
			{2, 4, 1, 3, 0},
			{3, 0, 2, 4, 1},
			{3, 1, 4, 2, 0},
			{4, 1, 3, 0, 2},
			{4, 2, 0, 3, 1}}},
	} {
		if got := NQueens(test.in); !reflect.DeepEqual(got, test.want) {
			t.Errorf("NQueens(%d) = %v; want %v", test.in, got, test.want)
		}
	}
}

func benchNQueens(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		NQueens(size)
	}
}

func BenchmarkNQueens5(b *testing.B) { benchNQueens(b, 5) }
func BenchmarkNQueens7(b *testing.B) { benchNQueens(b, 7) }
func BenchmarkNQueens9(b *testing.B) { benchNQueens(b, 9) }
