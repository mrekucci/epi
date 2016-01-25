// Copyright (c) 2016, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package recursion

import (
	"reflect"
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	for _, test := range []struct {
		partial [][]int
		want    bool
		solved  [][]int
	}{
		{nil, false, nil},
		{
			[][]int{
				{0, 2, 6, 0, 0, 0, 8, 1, 0},
				{3, 0, 0, 7, 0, 8, 0, 0, 6},
				{4, 0, 0, 0, 5, 0, 0, 0, 7},
				{0, 5, 0, 1, 0, 7, 0, 9, 0},
				{0, 0, 3, 9, 0, 5, 1, 0, 0},
				{0, 4, 0, 3, 0, 2, 0, 5, 0},
				{1, 0, 0, 0, 3, 0, 0, 0, 2},
				{5, 0, 0, 2, 0, 4, 0, 0, 9},
				{0, 3, 8, 0, 0, 0, 4, 6, 0}},
			true,
			[][]int{
				{7, 2, 6, 4, 9, 3, 8, 1, 5},
				{3, 1, 5, 7, 2, 8, 9, 4, 6},
				{4, 8, 9, 6, 5, 1, 2, 3, 7},
				{8, 5, 2, 1, 4, 7, 6, 9, 3},
				{6, 7, 3, 9, 8, 5, 1, 2, 4},
				{9, 4, 1, 3, 6, 2, 7, 5, 8},
				{1, 9, 4, 8, 3, 6, 5, 7, 2},
				{5, 6, 7, 2, 1, 4, 3, 8, 9},
				{2, 3, 8, 5, 7, 9, 4, 6, 1}},
		},
		{
			[][]int{
				{5, 3, 0, 0, 7, 0, 0, 0, 0},
				{6, 0, 0, 1, 9, 5, 0, 0, 0},
				{0, 9, 8, 0, 0, 0, 0, 6, 0},
				{8, 0, 0, 0, 6, 0, 0, 0, 3},
				{4, 0, 0, 8, 0, 3, 0, 0, 1},
				{7, 0, 0, 0, 2, 0, 0, 0, 6},
				{0, 6, 0, 0, 0, 0, 2, 8, 0},
				{0, 0, 0, 4, 1, 9, 0, 0, 5},
				{0, 0, 0, 0, 8, 0, 0, 7, 9}},
			true,
			[][]int{
				{5, 3, 4, 6, 7, 8, 9, 1, 2},
				{6, 7, 2, 1, 9, 5, 3, 4, 8},
				{1, 9, 8, 3, 4, 2, 5, 6, 7},
				{8, 5, 9, 7, 6, 1, 4, 2, 3},
				{4, 2, 6, 8, 5, 3, 7, 9, 1},
				{7, 1, 3, 9, 2, 4, 8, 5, 6},
				{9, 6, 1, 5, 3, 7, 2, 8, 4},
				{2, 8, 7, 4, 1, 9, 6, 3, 5},
				{3, 4, 5, 2, 8, 6, 1, 7, 9}},
		},
	} {
		if got := SolveSudoku(test.partial); got != test.want || !reflect.DeepEqual(test.partial, test.solved) {
			t.Errorf("SolveSudoku = %t; want %t\ngot:  %v\nwant: %v\n", got, test.want, test.partial, test.solved)
		}
	}
}

func BenchmarkSolveSudoku(b *testing.B) {
	grid := [][]int{
		{0, 2, 6, 0, 0, 0, 8, 1, 0},
		{3, 0, 0, 7, 0, 8, 0, 0, 6},
		{4, 0, 0, 0, 5, 0, 0, 0, 7},
		{0, 5, 0, 1, 0, 7, 0, 9, 0},
		{0, 0, 3, 9, 0, 5, 1, 0, 0},
		{0, 4, 0, 3, 0, 2, 0, 5, 0},
		{1, 0, 0, 0, 3, 0, 0, 0, 2},
		{5, 0, 0, 2, 0, 4, 0, 0, 9},
		{0, 3, 8, 0, 0, 0, 4, 6, 0}}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SolveSudoku(grid)
	}
}
