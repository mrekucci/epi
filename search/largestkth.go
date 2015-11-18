// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package search

import "math/rand"

// partition returns an index of the new pivot and partitions elements in an such
// as: an[l:pivot] contains elements that are less then the pivot and
// an[pivot+1:r+1] contains elements that are greater then the pivot.
func partition(an []int, l, r int) int {
	q := l
	pivot := rand.Intn(r-l+1) + l
	an[pivot], an[r] = an[r], an[pivot]
	// Loop invariant:
	// each element in data[l:q] is less than the pivot (elements are distinct);
	// each element in data[q:u] is greater than the pivot;
	// each element in data[u:r] is unsorted;
	// the element data[r] holds the pivot.
	for u := q; u < r; u++ {
		if an[u] > an[r] {
			an[u], an[q] = an[q], an[u]
			q++
		}
	}
	an[r], an[q] = an[q], an[r]
	return q
}

// LargestKth returns the value of k-th largest element in an.
// Slice must contain only distinct values an k must be from range: 0 < k >= len(an).
// The an may be modified during the function execution.
// The time complexity is O(n) in average. The O(1) additional space is needed.
func LargestKth(an []int, k int) (e int, ok bool) {
	if k <= 0 || k > len(an) {
		return 0, false
	}

	l, r := 0, len(an)-1
	for l <= r {
		switch np := partition(an, l, r); {
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
