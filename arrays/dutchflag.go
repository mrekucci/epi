// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package arrays

// Rearrange change order of given slice such that all elements
// less than an[i] appear first, followed by elements equal to
// an[i], followed by elements greater than an[i],
func Rearrange(an []int, i int) {
	pv := an[i]
	p, q := 0, 0
	r := len(an) - 1
	// Loop invariant: elements less then pivot an[0:p-1], elements equal to pivot an[p:q-1],
	// unclassified an[q:r], elements greater then pivot an[r+1:len(an)-1].
	for q <= r {
		switch {
		case an[q] < pv:
			an[q], an[p] = an[p], an[q]
			p++
			q++
		case an[q] > pv:
			an[q], an[r] = an[r], an[q]
			r--
		default: // a[q] == p
			q++
		}
	}
}
