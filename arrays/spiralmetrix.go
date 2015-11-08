// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package arrays

// add adds elements to ac for offset off in spiral order and returns modified mxc.
func add(mx [][]int, mxc []int, off int) []int {
	l := len(mx) - 1 - off // n := len(mx)-1

	// For matrix with odd size append center element when layer is there.
	if off == l {
		return append(mxc, mx[off][off])
	}

	// From (0,0) to (O,n-1).
	for j := off; j < l; j++ {
		mxc = append(mxc, mx[off][j])
	}

	// From (O,n) to (n-1,n).
	for i := off; i < l; i++ {
		mxc = append(mxc, mx[i][l])
	}

	// From (n,n) to (n,1).
	for j := l; j > off; j-- {
		mxc = append(mxc, mx[l][j])
	}

	// From (n,0) to (1,0).
	for i := l; i > off; i-- {
		mxc = append(mxc, mx[i][off])
	}

	return mxc
}

// Clockwise returns mxc slice which elements are ordered under
// clockwise spiral order of the original two-dimensional matrix mx.
// The time complexity is O(n*n) and O(1) additional space is needed.
func Clockwise(mx [][]int) (mxc []int) {
	ctr := len(mx) >> 1 // Center of matrix.
	if len(mx)%2 != 0 {
		ctr++
	}
	for off := 0; off < ctr; off++ {
		mxc = add(mx, mxc, off)
	}
	return mxc
}
