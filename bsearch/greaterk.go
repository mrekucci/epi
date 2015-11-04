// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package bsearch

// FirstGreaterK returns the index of the first element in an greater then k.
// The time complexity is O(log(n)). The O(1) additional space is needed.
func FirstGreaterK(an []int, k int) int {
	g := -1
	l, r := 0, len(an)-1
	for l <= r {
		m := l + (r-l)/2
		if k < an[m] {
			g, r = m, m-1
		} else {
			l = m + 1
		}
	}
	return g
}
