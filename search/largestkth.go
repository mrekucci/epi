// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package search

import "math/rand"

// partition returns an index of the new pivot and partitions elements in an such
// as: an[l:pivot] contains elements that are less then the pivot and
// an[pivot+1:r+1] contains elements that are greater then the pivot.
func partition(an []int, l, r, p int) int {
	v := an[p]
	an[p], an[r] = an[r], an[p]
	for i := l; i < r; i++ {
		if an[i] > v {
			an[i], an[l] = an[l], an[i]
			l++
		}
	}
	an[r], an[l] = an[l], an[r]
	return l
}

// LargestKth returns the value of k-th largest element in an.
// Slice must contain only distinct values an k must be from range: 0 < k >= len(an).
// The an may be modified during the function execution.
// The time complexity is O(n). The O(1) additional space is needed.
func LargestKth(an []int, k int) (e int, ok bool) {
	if k <= 0 || k > len(an) {
		return 0, false
	}

	l, r := 0, len(an)-1
	for l <= r {
		p := rand.Intn(r-l+1) + l
		switch np := partition(an, l, r, p); {
		case np == k-1:
			return an[np], true
		case np > k-1:
			r = np - 1
		case np < k-1:
			l = np + 1
		}
	}

	panic("unreachable statement") // This case should never happened.
}
