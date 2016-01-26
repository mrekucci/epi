// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package arrays

// reverseInts reverse elements of a[s:e] in a.
func reverseInts(a []int, s, e int) {
	for e > s {
		a[s], a[e] = a[e], a[s]
		s++
		e--
	}
}

// NextPerm returns next permutation of p under lexicographical order.
// Returns zero value for slice if p already contains last permutation.
// The time complexity is O(n) and O(1) additional space is needed.
func NextPerm(p []int) (next []int) {
	k := len(p) - 2
	for k >= 0 && p[k] >= p[k+1] {
		k--
	}

	if k == -1 { // p already contains last permutation in lexicographical order.
		return nil
	}

	l := 0
	for i := k + 1; i < len(p); i++ { // Find the smallest entry in p[k+1:] decreasing suffix that is largest then p[k].
		if p[k] < p[i] {
			l = i
		} else {
			break
		}
	}

	next= append(next, p...)
	next[k], next[l] = next[l], next[k]
	// Because swapping next[k] and next[l] leaves the sequence after position k in decreasing order, we
	// must reverse this sub-sequence to produce its lexicographically minimal permutation.
	reverseInts(next, k+1, len(next)-1)

	return next
}
