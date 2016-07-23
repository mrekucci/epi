// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package bsearch

// IndexEqualsEntry returns the index of the first entry in xs that equals to its index.
// The time complexity is O(log(n)). The O(1) additional space is needed.
func IndexEqualsEntry(xs []int) int {
	l, r := 0, len(xs)-1
	for l <= r {
		m := l + (r-l)/2
		switch d := xs[m] - m; {
		case d < 0:
			l = m + 1
		case d == 0:
			return m
		case d > 0:
			r = m - 1
		}
	}
	return -1
}
