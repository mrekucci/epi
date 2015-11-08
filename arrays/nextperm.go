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
// Returns empty slice if p already contains last permutation.
// The time complexity is O(n) and O(1) additional space is needed.
func NextPerm(p []int) []int {
	k := len(p) - 2
	for k >= 0 && p[k] >= p[k+1] {
		k--
	}

	if k == -1 { // p already contains last permutation in lexicographical order.
		return []int{}
	}

	l := 0
	for i := k + 1; i < len(p); i++ { // Find the smallest entry in p[k+1:] decreasing suffix that is largest then p[k].
		if p[k] < p[i] {
			l = i
		} else {
			break
		}
	}

	pn := append([]int(nil), p...)
	pn[k], pn[l] = pn[l], pn[k]
	// Because swapping pn[k] and pn[l] leaves the sequence after position k in decreasing order, we
	// must reverse this sub-sequence to produce its lexicographically minimal permutation.
	reverseInts(pn, k+1, len(pn)-1)

	return pn
}
