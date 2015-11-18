// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package arrays

// Rearrange change order of given slice such that all elements
// less than an[i] appear first, followed by elements equal to
// an[i], followed by elements greater than an[i].
// The time complexity is O(n) and O(1) additional space is needed.
func Rearrange(an []int, i int) {
	v := an[i]
	p, q, r := 0, 0, len(an)-1
	// Loop invariant:
	// elements less then the pivot an[0:p];
	// elements equal to the pivot an[p:q];
	// elements that are unclassified an[q:r+1];
	// elements greater then the pivot an[r+1:len(an)-1].
	// The time complexity is O(n) because on each iteration we decrease number of unclassified elements by 1.
	for q <= r {
		switch {
		case an[q] < v:
			an[q], an[p] = an[p], an[q]
			p++
			q++
		case an[q] > v:
			an[q], an[r] = an[r], an[q]
			r--
		case an[q] == v:
			q++
		}
	}
}
