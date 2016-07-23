// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package bsearch

// FirstK returns the index of the first leftmost occurrence of k in xs.
// The time complexity is O(log(n)). The O(1) additional space is needed.
func FirstK(xs []int, k int) int {
	f := -1
	l, r := 0, len(xs)-1
	for l <= r {
		m := l + (r-l)/2 // Not (r+l)/2 'cause we want to avoid overflow.
		switch v := xs[m]; {
		case v > k:
			r = m - 1
		case v == k:
			f, r = m, m-1
		case v < k:
			l = m + 1
		}
	}
	return f
}
