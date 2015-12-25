// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package invariants

import "sort"

// hasTwoSum returns true if the sum of two (not distinct) elements results in k.
// The time complexity is O(n), and O(1) additional space is needed.
func hasTwoSum(an []int, k int) bool {
	i, j := 0, len(an)-1
	for i <= j { // Operation <= takes not distinct values into a credit.
		switch sum := an[i] + an[j]; {
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

// HasThreeSum returns true if the sum of three (not distinct) elements results in k.
// The time complexity is O(n*n), and O(1) additional space is needed.
func HasThreeSum(an []int, k int) bool {
	sort.Ints(an)
	for _, v := range an {
		if hasTwoSum(an, k-v) { // Find if the sum of two numbers in an results in k-v.
			return true
		}
	}
	return false
}
