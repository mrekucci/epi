// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package invariants

import "sort"

// hasTwoSum returns true if the sum of two (not distinct) elements results into k.
// The time complexity is O(n), and O(1) additional space is needed.
func hasTwoSum(xs []int, k int) bool {
	i, j := 0, len(xs)-1
	for i <= j { // Operation <= takes not distinct values into a credit.
		switch sum := xs[i] + xs[j]; {
		case sum == k:
			return true
		case sum < k:
			i++
		case sum > k:
			j--
		}
	}
	return false
}

// HasThreeSum returns true if the sum of three (not distinct) elements results into k.
// The time complexity is O(n*n), and O(1) additional space is needed.
func HasThreeSum(xs []int, k int) bool {
	sort.Ints(xs)
	for _, v := range xs {
		if hasTwoSum(xs, k-v) { // Find if the sum of two numbers in xs results into k-v.
			return true
		}
	}
	return false
}
