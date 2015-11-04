// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package search

// Matrix returns true if m contains x, otherwise returns false.
// The time complexity is O(m+n) where m is number of row and n is
// number of columns in matrix. The O(1) additional space is needed.
func Matrix(m [][]int, x int) bool {
	r, c := 0, len(m)-1
	for r < len(m) && c >= 0 {
		switch v := m[r][c]; {
		case v == x:
			return true
		case v > x:
			c--
		case v < x:
			r++
		}
	}
	return false
}
