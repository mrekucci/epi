// Copyright (c) 2015, Peter Mrekaj. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.txt file.

package search

// MinMax returns the min and max values from an.
// The time complexity is O(n). The O(1) additional space is needed.
func MinMax(an []int) (min, max int) {
	if len(an) == 0 {
		return 0, 0
	}

	min, max = an[0], an[0]
	for _, v := range an { // We start iterate again from index 0 to avoid array bounds checking and thus improve performance.
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return min, max
}

// compare returns a, b if a < b, otherwise b, a.
func compare(a, b int) (min, max int) {
	if a < b {
		return a, b
	}
	return b, a
}

// MinMaxAlt is an alternative function to MinMax for searching min and max in an.
// The time complexity is O(n). The O(1) additional space is needed.
// The time complexity hidden constant is more significant here than in the MaxMin.
func MinMaxAlt(an []int) (min, max int) {
	switch len(an) {
	case 0:
		return 0, 0
	case 1:
		return an[0], an[0]
	}

	min, max = compare(an[0], an[1])
	for i := 3; i < len(an); i += 2 { // Compare 3/2n - 2 times.
		nMin, nMax := compare(an[i-1], an[i])
		min, _ = compare(nMin, min)
		_, max = compare(nMax, max)
	}

	if len(an)%2 != 0 { // We need to compare the last element if length of an is odd.
		v := an[len(an)-1]
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}

	return min, max
}
