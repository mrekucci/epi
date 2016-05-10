// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package arrays

// add adds elements of matrix represented by a square layer distant by offset from the most outer layer.
func add(clockwise []int, matrix [][]int, off int) []int {
	l := len(matrix) - 1 - off

	// For matrix with odd size, append the center element when layer is there.
	if off == l {
		return append(clockwise, matrix[off][off])
	}

	// From (0,0) to (O,n-1).
	for j := off; j < l; j++ {
		clockwise = append(clockwise, matrix[off][j])
	}

	// From (O,n) to (n-1,n).
	for i := off; i < l; i++ {
		clockwise = append(clockwise, matrix[i][l])
	}

	// From (n,n) to (n,1).
	for j := l; j > off; j-- {
		clockwise = append(clockwise, matrix[l][j])
	}

	// From (n,0) to (1,0).
	for i := l; i > off; i-- {
		clockwise = append(clockwise, matrix[i][off])
	}

	return clockwise
}

// Clockwise returns a slice which elements are ordered under
// clockwise spiral order of the original two-dimensional matrix.
// The time complexity is O(n*n) and O(1) additional space is needed.
func Clockwise(matrix [][]int) (clockwise []int) {
	mid := len(matrix) / 2
	if len(matrix)%2 != 0 {
		mid++
	}
	for off := 0; off < mid; off++ {
		clockwise = add(clockwise, matrix, off)
	}
	return clockwise
}
