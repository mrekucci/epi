// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package arrays

// Rearrange change order of given slice such that all elements
// less than xs[i] appear first, followed by elements equal to
// xs[i], followed by elements greater than xs[i].
// The time complexity is O(n) and O(1) additional space is needed.
func Rearrange(xs []int, i int) {
	x := xs[i]
	p, q, r := 0, 0, len(xs)-1
	// Loop invariant:
	// elements less then the pivot xs[0:p];
	// elements equal to the pivot xs[p:q];
	// elements that are unclassified xs[q:r+1];
	// elements greater then the pivot xs[r+1:len(xs)].
	// The time complexity is O(n) because on each iteration we decrease number of unclassified elements by 1.
	for q <= r {
		switch {
		case xs[q] < x:
			xs[q], xs[p] = xs[p], xs[q]
			p++
			q++
		case xs[q] > x:
			xs[q], xs[r] = xs[r], xs[q]
			r--
		case xs[q] == x:
			q++
		}
	}
}
