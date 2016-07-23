// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package search

import "math/rand"

// partition returns an index of the new pivot and partitions elements in xs such
// as: xs[l:pivot] contains elements that are greater then the pivot and
// xs[pivot+1:r+1] contains elements that are less then the pivot.
func partition(xs []int, l, r int) int {
	p := l
	pivot := rand.Intn(r-l+1) + l
	xs[pivot], xs[r] = xs[r], xs[pivot]
	// Loop invariant:
	// each element in data[l:p] is greater than the pivot;
	// each element in data[p:q] is less than or equal to the pivot;
	// each element in data[q:r] is unsorted;
	// the element data[r] holds the pivot.
	for q := p; q < r; q++ {
		if xs[q] > xs[r] {
			xs[q], xs[p] = xs[p], xs[q]
			p++
		}
	}
	xs[r], xs[p] = xs[p], xs[r]
	return p
}

// LargestKth returns the value of k-th largest element in xs.
// Slice must contain only distinct values and k must be from range: 0 < k >= len(xs).
// The xs may be modified during the function execution.
// The time complexity is O(n) in average. The O(1) additional space is needed.
func LargestKth(xs []int, k int) (e int, ok bool) {
	if k <= 0 || k > len(xs) {
		return 0, false
	}

	l, r := 0, len(xs)-1
	for l <= r {
		switch p := partition(xs, l, r); {
		case p == k-1:
			return xs[p], true
		case p > k-1:
			r = p - 1
		case p < k-1:
			l = p + 1
		}
	}

	panic("unreachable statement") // This case should never happened.
}
