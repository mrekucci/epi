// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package epi

// add adds elements to ac for offset off in spiral order and returns modified ac.
func add(a [][]int, ac []int, off int) []int {
	l := len(a) - 1 - off

	// For array with odd size append center element when layer is there.
	if off == l {
		return append(ac, a[off][off])
	}

	// From (0,0) to (O,n-2).
	for j := off; j < l; j++ {
		ac = append(ac, a[off][j])
	}

	// From (O,n-1) to (n-2,n-1).
	for i := off; i < l; i++ {
		ac = append(ac, a[i][l])
	}

	// From (n-1,n-1) to (n-1,1).
	for j := l; j > off; j-- {
		ac = append(ac, a[l][j])
	}

	// From (n-1,0) to (1,0).
	for i := l; i > off; i-- {
		ac = append(ac, a[i][off])
	}

	return ac
}

// Clockwise returns an array which elements are under
// clockwise spiral order of the original array a.
func Clockwise(a [][]int) []int {
	var ac []int

	// Center of a.
	ctr := len(a) >> 1
	if len(a)%2 != 0 {
		ctr++
	}
	for off := 0; off < ctr; off++ {
		ac = add(a, ac, off)
	}

	return ac
}
